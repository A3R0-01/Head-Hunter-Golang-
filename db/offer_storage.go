package db

import (
	"context"
	"fmt"

	"github.com/A3R0-01/head-hunter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const OfferCollName = "Offers"

type OfferStore interface {
	GetOfferByID(context.Context, string) (*types.Offer, error)
	GetOffers(context.Context, Map) ([]*types.Offer, error)
	CreateOffer(context.Context, *types.Offer) (*types.Offer, error)
	DeleteOffer(context.Context, string) error
	// UpdateOffer(context.Context, string, Map) error
}

type MongoOfferStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoOfferStore(client *mongo.Client) OfferStore {
	return &MongoOfferStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(OfferCollName),
	}
}
func NewMongoOfferStoreTest(client *mongo.Client) OfferStore {
	return &MongoOfferStore{
		client: client,
		coll:   client.Database(TestDBNAME).Collection(UserCollName),
	}
}

func (store *MongoOfferStore) GetOfferByID(ctx context.Context, id string) (*types.Offer, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	var offer types.Offer
	if err := store.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&offer); err != nil {
		return nil, err
	}
	return &offer, nil
}
func (store *MongoOfferStore) GetOffers(ctx context.Context, filter Map) ([]*types.Offer, error) {

	cur, err := store.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var offers []*types.Offer
	if err := cur.All(ctx, &offers); err != nil {
		return nil, err
	}
	return offers, nil
}

func (store *MongoOfferStore) CreateOffer(ctx context.Context, offer *types.Offer) (*types.Offer, error) {
	result, err := store.coll.InsertOne(ctx, offer)
	if err != nil {
		return nil, err
	}
	offer.ID = result.InsertedID.(primitive.ObjectID)
	return offer, nil
}

func (store *MongoOfferStore) DeleteOffer(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("invalid offer")
	}
	res, err := store.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	} else if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil

}

// func (store *MongoOfferStore) UpdateOffer(ctx context.Context, id string, values types.UpdateOfferParams) error {
// 	updateVal, err := values.ToMongoBson()
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
