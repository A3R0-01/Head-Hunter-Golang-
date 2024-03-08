package db

import (
	"context"
	"fmt"

	"github.com/A3R0-01/head-hunter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	UserCollName = "User"
)

type UserStore interface {
	Dropper
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context, Map) ([]*types.User, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(context.Context, string, types.UpdateUserParams) error
	GetUserByEmail(context.Context, string) (*types.User, error)
}

type MongoUserStorage struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore() *MongoUserStorage {
	client := NewMongoClient()
	return &MongoUserStorage{
		client: client,
		coll:   client.Database(DBNAME).Collection(UserCollName),
	}
}
func NewMongoUserStoreTest() *MongoUserStorage {
	client := NewMongoClient()
	return &MongoUserStorage{
		client: client,
		coll:   client.Database(TestDBNAME).Collection(UserCollName),
	}
}

func (store *MongoUserStorage) GetUserByID(ctx context.Context, id string) (*types.User, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	var user types.User
	if err := store.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
func (store *MongoUserStorage) GetUsers(ctx context.Context, filter Map) ([]*types.User, error) {

	cur, err := store.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var users []*types.User
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (store *MongoUserStorage) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	result, err := store.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (store *MongoUserStorage) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("invalid user")
	}
	res, err := store.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	} else if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil

}
func (store *MongoUserStorage) UpdateUser(ctx context.Context, id string, values types.UpdateUserParams) error {
	updateVal, err := values.ToUpdateMongo()
	if err != nil {
		return err
	}
	update := bson.D{
		{
			"$set", updateVal,
		},
	}
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

func (store *MongoUserStorage) Drop(ctx context.Context) error {
	fmt.Println("----- dropping user collection------")
	return store.coll.Drop(ctx)
}

func (store *MongoUserStorage) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var user types.User
	if err := store.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
