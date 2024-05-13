package db

import (
	"context"
	"fmt"

	"github.com/A3R0-01/head-hunter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const JobPostCollName = "JobPosts"

type JobPostStore interface {
	GetJobPostByID(context.Context, string) (*types.JobPost, error)
	GetJobPosts(context.Context, Map) ([]*types.JobPost, error)
	CreateJobPost(context.Context, *types.JobPost) (*types.JobPost, error)
	DeleteJobPost(context.Context, string) error
	UpdateJobPost(context.Context, string, types.UpdateJobPostParams) error
}

type MongoJobStorage struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoJobStore(client *mongo.Client) JobPostStore {
	return &MongoJobStorage{
		client: client,
		coll:   client.Database(DBNAME).Collection(JobPostCollName),
	}
}
func NewMongoJobStoreTest(client *mongo.Client) JobPostStore {
	return &MongoJobStorage{
		client: client,
		coll:   client.Database(TestDBNAME).Collection(JobPostCollName),
	}
}

func (store *MongoJobStorage) GetJobPostByID(ctx context.Context, id string) (*types.JobPost, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	var JobPost types.JobPost
	if err := store.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&JobPost); err != nil {
		return nil, err
	}
	return &JobPost, nil
}
func (store *MongoJobStorage) GetJobPosts(ctx context.Context, filter Map) ([]*types.JobPost, error) {

	cur, err := store.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var jobPosts []*types.JobPost
	if err := cur.All(ctx, &jobPosts); err != nil {
		return nil, err
	}
	return jobPosts, nil
}

func (store *MongoJobStorage) CreateJobPost(ctx context.Context, jobPost *types.JobPost) (*types.JobPost, error) {
	result, err := store.coll.InsertOne(ctx, jobPost)
	if err != nil {
		return nil, err
	}
	jobPost.ID = result.InsertedID.(primitive.ObjectID)
	return jobPost, nil
}

func (store *MongoJobStorage) DeleteJobPost(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("invalid JobPost")
	}
	res, err := store.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	} else if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil

}
func (store *MongoJobStorage) UpdateJobPost(ctx context.Context, id string, values types.UpdateJobPostParams) error {
	updateVal, err := values.ToMongoBson()
	if err != nil {
		fmt.Println(err)
		return err
	}
	update := bson.M{"$set": updateVal}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("reached")
	_, err = store.coll.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return err
	}
	return nil
}
