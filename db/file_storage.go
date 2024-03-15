package db

import (
	"context"
	"fmt"

	"github.com/A3R0-01/head-hunter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

const FileCollName = "Files"

type FileStore interface {
	GetFileByID(context.Context, string) (*types.File, error)
	GetFiles(context.Context, Map) ([]*types.File, error)
	CreateFile(context.Context, *types.File) (*types.File, error)
	DeleteFile(context.Context, string) error
	UpdateFile(context.Context, string, Map) error
}

type MongoFileStore struct {
	client *mongo.Client
	coll   *mongo.Collection
	bucket *gridfs.Bucket
}

func NewMongoFileStore(client *mongo.Client) *MongoFileStore {
	bucket, err := gridfs.NewBucket(client.Database(DBNAME))
	if err != nil {
		panic(err)
	}
	return &MongoFileStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(FileCollName),
		bucket: bucket,
	}
}
func NewMongoFileStoreTest(client *mongo.Client) *MongoFileStore {
	bucket, err := gridfs.NewBucket(client.Database(DBNAME))
	if err != nil {
		panic(err)
	}
	return &MongoFileStore{
		client: client,
		coll:   client.Database(TestDBNAME).Collection(FileCollName),
		bucket: bucket,
	}
}

func (store *MongoFileStore) GetFileByID(ctx context.Context, id string) (*types.GridfsFile, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	filter := bson.M{"metadata": bson.M{"User": oid}}
	cursor, err := store.bucket.Find(filter)
	if err != nil {
		panic(err)
	}
	var files []types.GridfsFile
	if err = cursor.All(context.TODO(), &files); err != nil {
		panic(err)
	}
	return &files[0], nil
}
func (store *MongoFileStore) GetFiles(ctx context.Context, filter Map) ([]*types.File, error) {

	cur, err := store.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var files []*types.File
	if err := cur.All(ctx, &files); err != nil {
		return nil, err
	}
	return files, nil
}

func (store *MongoFileStore) CreateFile(ctx context.Context, file *types.File) (*types.File, error) {
	result, err := store.coll.InsertOne(ctx, file)
	if err != nil {
		return nil, err
	}
	file.ID = result.InsertedID.(primitive.ObjectID)
	return file, nil
}

func (store *MongoFileStore) DeleteFile(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("invalid file")
	}
	res, err := store.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	} else if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil

}
