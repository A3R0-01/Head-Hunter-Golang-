package db

import (
	"context"
	"fmt"

	"github.com/A3R0-01/head-hunter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const IndustryCollName = "Industries"

type IndustryStore interface {
	GetIndustryByID(context.Context, string) (*types.Industry, error)
	GetIndustries(context.Context, Map) ([]*types.Industry, error)
	CreateIndustry(context.Context, *types.Industry) (*types.Industry, error)
	DeleteIndustry(context.Context, string) error
	// UpdateIndustry(context.Context, string, types.UpdateIndustryParams) error
}

type MongoIndustryStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoIndustryStore(client *mongo.Client) *MongoIndustryStore {
	return &MongoIndustryStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(IndustryCollName),
	}
}
func NewMongoIndustryStoreTest(client *mongo.Client) *MongoIndustryStore {
	return &MongoIndustryStore{
		client: client,
		coll:   client.Database(TestDBNAME).Collection(IndustryCollName),
	}
}

func (store *MongoIndustryStore) GetIndustryByID(ctx context.Context, id string) (*types.Industry, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	var industry types.Industry
	if err := store.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&industry); err != nil {
		return nil, err
	}
	return &industry, nil
}
func (store *MongoIndustryStore) GetIndustries(ctx context.Context, filter Map) ([]*types.Industry, error) {

	cur, err := store.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var industries []*types.Industry
	if err := cur.All(ctx, &industries); err != nil {
		return nil, err
	}
	return industries, nil
}

func (store *MongoIndustryStore) CreateIndustry(ctx context.Context, industry *types.Industry) (*types.Industry, error) {
	result, err := store.coll.InsertOne(ctx, industry)
	if err != nil {
		return nil, err
	}
	industry.ID = result.InsertedID.(primitive.ObjectID)
	return industry, nil
}

func (store *MongoIndustryStore) DeleteIndustry(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("invalid industry")
	}
	res, err := store.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	} else if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil

}

// func (store *MongoIndustryStore) UpdateIndustry(ctx context.Context, id string, values types.UpdateIndustryParams) error {
// 	updateVal, err := values.ToUpdateMongo()
// 	if err != nil {
// 		return err
// 	}
// 	update := bson.D{
// 		{
// 			"$set", updateVal,
// 		},
// 	}
// 	oid, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = store.coll.UpdateOne(ctx, bson.M{"_id": oid}, update)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
