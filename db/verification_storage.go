package db

import (
	"context"
	"fmt"

	"github.com/A3R0-01/head-hunter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type VerificationStore interface {
	GetVerificationByID(context.Context, string) (*types.Verification, error)
	GetVerifications(context.Context, Map) ([]*types.Verification, error)
	CreateVerification(context.Context, *types.Verification) (*types.Verification, error)
	DeleteVerification(context.Context, string) error
	// UpdateVerification(context.Context, string, types.UpdateVerificationParams) error
}

var VerificationCollName = "Verifcation"

type MongoVerificationStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoVerificationStore(client *mongo.Client) VerificationStore {
	return &MongoVerificationStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(VerificationCollName),
	}
}
func NewMongoVerificationStoreTest(client *mongo.Client) VerificationStore {
	return &MongoVerificationStore{
		client: client,
		coll:   client.Database(TestDBNAME).Collection(VerificationCollName),
	}
}

func (store *MongoVerificationStore) GetVerificationByID(ctx context.Context, id string) (*types.Verification, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	var verification types.Verification
	if err := store.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&verification); err != nil {
		return nil, err
	}
	return &verification, nil
}
func (store *MongoVerificationStore) GetVerifications(ctx context.Context, filter Map) ([]*types.Verification, error) {

	cur, err := store.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var verifications []*types.Verification
	if err := cur.All(ctx, &verifications); err != nil {
		return nil, err
	}
	return verifications, nil
}

func (store *MongoVerificationStore) CreateVerification(ctx context.Context, verification *types.Verification) (*types.Verification, error) {
	result, err := store.coll.InsertOne(ctx, verification)
	if err != nil {
		return nil, err
	}
	verification.ID = result.InsertedID.(primitive.ObjectID)
	return verification, nil
}

func (store *MongoVerificationStore) DeleteVerification(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("invalid verification")
	}
	res, err := store.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	} else if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil

}

// func (store *MongoVerificationStore) UpdateVerification(ctx context.Context, id string, values types.UpdateVerificationParams) error {
// 	updateVal, err := values.ToUpdateMongo()
// 	if err != nil {
// 		return err
// 	}
// 	update := bson.M{"$set": updateVal}
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
