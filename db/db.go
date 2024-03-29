package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBNAME string
var DBURI string
var TestDBNAME string

type Store struct {
	CompanyStore
	VerificationStore
	IndustryStore
	UserStore
	JobPostStore
	ApplicationStore
	OfferStore
	SessionStore
	MessageStore
	// FileStore
}
type Map map[string]any
type Dropper interface {
	Drop(context.Context) error
}

func NewStore(client *mongo.Client) *Store {
	var (
		userStore         = NewMongoUserStore(client)
		companyStore      = NewMongoCompanyStore(client)
		jobPostStore      = NewMongoJobStore(client)
		applicationStore  = NewMongoApplicationStore(client)
		offerStore        = NewMongoOfferStore(client)
		sessionStore      = NewMongoSessionStore(client)
		messageStore      = NewMongoMessageStore(client)
		industryStore     = NewMongoIndustryStore(client)
		verificationStore = NewMongoVerificationStore(client)
	)

	return &Store{
		UserStore:         userStore,
		CompanyStore:      companyStore,
		JobPostStore:      jobPostStore,
		ApplicationStore:  applicationStore,
		OfferStore:        offerStore,
		SessionStore:      sessionStore,
		IndustryStore:     industryStore,
		VerificationStore: verificationStore,
		MessageStore:      messageStore,
	}
}
func NewMongoClient() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DBURI))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func init() {
	if err := godotenv.Load(); err != nil {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatal(err)
		}
	}
	DBNAME = os.Getenv("MONGO_DB_NAME")
	DBURI = os.Getenv("MONGO_DB_URL")
	TestDBNAME = os.Getenv("TEST_DB_NAME")
}
