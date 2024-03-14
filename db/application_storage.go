package db

import (
	"context"
	"fmt"

	"github.com/A3R0-01/head-hunter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const ApplicationCollName = "Applications"

type ApplicationStore interface {
	GetApplicationByID(context.Context, string) (*types.Application, error)
	GetApplications(context.Context, Map) ([]*types.Application, error)
	CreateApplication(context.Context, *types.Application) (*types.Application, error)
	DeleteApplication(context.Context, string) error
}
type MongoApplicationStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongApplicationrStore(client *mongo.Client) *MongoApplicationStore {
	return &MongoApplicationStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(ApplicationCollName),
	}
}
func NewMongoApplicationStoreTest(client *mongo.Client) *MongoApplicationStore {
	return &MongoApplicationStore{
		client: client,
		coll:   client.Database(TestDBNAME).Collection(ApplicationCollName),
	}
}
func (store *MongoApplicationStore) GetApplicationByID(ctx context.Context, id string) (*types.Application, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	var application types.Application
	if err := store.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&application); err != nil {
		return nil, err
	}
	return &application, nil
}
func (store *MongoApplicationStore) GetApplications(ctx context.Context, filter Map) ([]*types.Application, error) {

	cur, err := store.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var applications []*types.Application
	if err := cur.All(ctx, &applications); err != nil {
		return nil, err
	}
	return applications, nil
}

func (store *MongoApplicationStore) CreateApplication(ctx context.Context, application *types.Application) (*types.Application, error) {
	result, err := store.coll.InsertOne(ctx, application)
	if err != nil {
		return nil, err
	}
	application.ID = result.InsertedID.(primitive.ObjectID)
	return application, nil
}

func (store *MongoApplicationStore) DeleteApplication(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("invalid application")
	}
	res, err := store.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	} else if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
