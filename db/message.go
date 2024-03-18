package db

import (
	"context"
	"fmt"

	"github.com/A3R0-01/head-hunter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const MessageCollName = "Messages"

type MessageStore interface {
	GetMessageByID(context.Context, string) (*types.Message, error)
	GetMessages(context.Context, Map) ([]*types.Message, error)
	CreateMessage(context.Context, *types.Message) (*types.Message, error)
	DeleteMessages(context.Context, string) error
}

type MongoMessageStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoMessageStore(client *mongo.Client) MessageStore {
	return &MongoMessageStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(MessageCollName),
	}
}
func NewMongoMessageStoreTest(client *mongo.Client) MessageStore {
	return &MongoMessageStore{
		client: client,
		coll:   client.Database(TestDBNAME).Collection(MessageCollName),
	}
}

func (store *MongoMessageStore) GetMessageByID(ctx context.Context, id string) (*types.Message, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	var message types.Message
	if err := store.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&message); err != nil {
		return nil, err
	}
	return &message, nil
}
func (store *MongoMessageStore) GetMessages(ctx context.Context, filter Map) ([]*types.Message, error) {

	cur, err := store.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var messages []*types.Message
	if err := cur.All(ctx, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

func (store *MongoMessageStore) CreateMessage(ctx context.Context, message *types.Message) (*types.Message, error) {
	result, err := store.coll.InsertOne(ctx, message)
	if err != nil {
		return nil, err
	}
	message.ID = result.InsertedID.(primitive.ObjectID)
	return message, nil
}

func (store *MongoMessageStore) DeleteMessages(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("invalid message")
	}
	res, err := store.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	} else if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil

}
