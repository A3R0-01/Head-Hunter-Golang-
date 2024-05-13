package db

import (
	"context"
	"fmt"

	"github.com/A3R0-01/head-hunter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const SessionCollName = "Sessions"

type SessionStore interface {
	GetSessionByID(context.Context, string) (*types.Session, error)
	GetSessions(context.Context, Map) ([]*types.Session, error)
	CreateSession(context.Context, *types.Session) (*types.Session, error)
	UpdateSession(context.Context, string, *types.UpdateSessionParams) error
	DeleteSession(context.Context, string) error
}

type MongoSessionStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoSessionStore(client *mongo.Client) SessionStore {
	return &MongoSessionStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(SessionCollName),
	}
}
func NewMongoSessionStoreTest(client *mongo.Client) SessionStore {
	return &MongoSessionStore{
		client: client,
		coll:   client.Database(TestDBNAME).Collection(SessionCollName),
	}
}

func (store *MongoSessionStore) GetSessionByID(ctx context.Context, id string) (*types.Session, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	var session types.Session
	if err := store.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&session); err != nil {
		return nil, err
	}
	return &session, nil
}
func (store *MongoSessionStore) GetSessions(ctx context.Context, filter Map) ([]*types.Session, error) {

	cur, err := store.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var sessions []*types.Session
	if err := cur.All(ctx, &sessions); err != nil {
		return nil, err
	}
	return sessions, nil
}

func (store *MongoSessionStore) CreateSession(ctx context.Context, session *types.Session) (*types.Session, error) {
	result, err := store.coll.InsertOne(ctx, session)
	if err != nil {
		return nil, err
	}
	session.ID = result.InsertedID.(primitive.ObjectID)
	return session, nil
}

func (store *MongoSessionStore) DeleteSession(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("invalid session")
	}
	res, err := store.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	} else if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil

}

func (store *MongoSessionStore) UpdateSession(ctx context.Context, id string, updateParams *types.UpdateSessionParams) error {
	update := bson.M{"$set": updateParams.ToUpdateMongo()}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = store.coll.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return err
	}
	return nil
}
