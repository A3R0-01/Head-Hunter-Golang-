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
	CompanyCollName = "Companies"
)

type CompanyStore interface {
	GetCompanyByID(context.Context, string) (*types.Company, error)
	GetCompanies(context.Context, Map) ([]*types.Company, error)
	CreateCompany(context.Context, *types.Company) (*types.Company, error)
	CreateRecruiterToken(context.Context, *types.Company) (string, error)
	GetRecruiters(context.Context, string) ([]*types.User, error)
	DeleteCompany(context.Context, string) error
	UpdateCompany(context.Context, string, types.UpdateCompanyParams) error
}

type MongoCompanyStorage struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoCompanyStore(client *mongo.Client) *MongoCompanyStorage {
	return &MongoCompanyStorage{
		client: client,
		coll:   client.Database(DBNAME).Collection(CompanyCollName),
	}
}
func NewMongoCompanyStoreTest(client *mongo.Client) *MongoCompanyStorage {
	return &MongoCompanyStorage{
		client: client,
		coll:   client.Database(TestDBNAME).Collection(CompanyCollName),
	}
}

func (store *MongoCompanyStorage) GetCompanyByID(ctx context.Context, id string) (*types.Company, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	var company types.Company
	if err := store.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&company); err != nil {
		return nil, err
	}
	return &company, nil
}
func (store *MongoCompanyStorage) GetCompanies(ctx context.Context, filter Map) ([]*types.Company, error) {

	cur, err := store.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var companies []*types.Company
	if err := cur.All(ctx, &companies); err != nil {
		return nil, err
	}
	return companies, nil
}

func (store *MongoCompanyStorage) CreateCompany(ctx context.Context, company *types.Company) (*types.Company, error) {
	result, err := store.coll.InsertOne(ctx, company)
	if err != nil {
		return nil, err
	}
	company.ID = result.InsertedID.(primitive.ObjectID)
	return company, nil
}

func (store *MongoCompanyStorage) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("invalid company")
	}
	res, err := store.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	} else if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil

}
func (store *MongoCompanyStorage) UpdateCompany(ctx context.Context, id string, values types.UpdateCompanyParams) error {
	updateVal, err := values.ToMongoBson()
	if err != nil {
		return err
	}
	update := bson.M{"$set": updateVal}
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
