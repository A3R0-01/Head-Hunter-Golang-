package main

import (
	"log"
	"os"

	"github.com/A3R0-01/head-hunter/api"
	"github.com/A3R0-01/head-hunter/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	client := db.NewMongoClient()

	app := fiber.New()
	var (
		store               = db.NewStore(client)
		companyHandler      = api.NewCompanyHandler(store, "company handler")
		verificationHandler = api.NewVerificationHandler(store, "verification handler")
		industryHandler     = api.NewIndustryHandler(store, "industry handler")
		userHandler         = api.NewUserHandler(store, "user handler")
		jobPostHandler      = api.NewJobPostHandler(store, "job post handler")
		applicationHandler  = api.NewApplicationHandler(store, "application handler")
		sessionHandler      = api.NewSessionHandler(store, "session Handler")
		messageHandler      = api.NewMessageHandler(store, "message handler")
	)
	// UserHandlers
	app.Post("/user", userHandler.HandlePostUser)
	app.Get("/user/:id", userHandler.HandleGetUser)
	app.Get("/user", userHandler.HandleGetUsers)
	app.Put("/user/:id", userHandler.HandlePut)
	app.Delete("/user/:id", userHandler.HandleDelete)
	// Company Handlers
	app.Post("/company", companyHandler.HandlePostCompany)
	app.Get("/company/:id", companyHandler.HandleGetCompany)
	app.Get("/company", companyHandler.HandleGetCompanies)
	app.Put("/company/:id", companyHandler.HandlePut)
	app.Delete("/company/:id", companyHandler.HandleDelete)
	// Industry Handlers
	app.Post("/industry", industryHandler.HandlePostIndustry)
	app.Get("/industry/:id", industryHandler.HandleGetIndustry)
	app.Get("/industry", industryHandler.HandleGetIndustries)
	app.Delete("/industry/:id", industryHandler.HandleDelete)
	// Verification Handlers
	app.Get("/verification/:id", verificationHandler.HandleGetVerification)
	app.Get("/verification", verificationHandler.HandleGetVerifications)
	app.Put("/verification/:id", verificationHandler.HandlePut)
	app.Post("/verification", verificationHandler.HandlePostVerification)
	app.Delete("/verification/:id", verificationHandler.HandleDelete)
	// Job Post Handlers
	app.Get("/jobPost/:id", jobPostHandler.HandleGetJobPost)
	app.Get("/jobPost", jobPostHandler.HandleGetJobPosts)
	app.Put("/jobPost/:id", jobPostHandler.HandlePut)
	app.Post("/jobPost", jobPostHandler.HandlePostJobPost)
	app.Delete("/jobPost/:id", jobPostHandler.HandleDelete)

	// Application Handlers
	app.Get("/application/:id", applicationHandler.HandleGetApplication)
	app.Get("/application", applicationHandler.HandleGetApplications)
	app.Post("/application", applicationHandler.HandlePostApplication)
	app.Put("/recruiter/application/:id", applicationHandler.HandlePut)
	app.Delete("/application/:id", applicationHandler.HandleDelete)

	// Message Handlers

	app.Post("/message", messageHandler.HandlePostMessage)
	app.Get("/message/:id", messageHandler.HandleGetMessage)
	app.Get("/message", messageHandler.HandleGetMessages)
	app.Delete("/message/:id", messageHandler.HandleDelete)

	// Session Handlers

	app.Post("/session", sessionHandler.HandlePostSession)
	app.Get("/session/:id", sessionHandler.HandleGetSession)
	app.Get("/session", sessionHandler.HandleGetSessions)
	app.Put("/session/:id", sessionHandler.HandlePut)
	app.Delete("/session/:id", sessionHandler.HandleDelete)

	// Administrative handlers

	// sessions

	app.Listen(os.Getenv("HTTP_LISTEN_ADDRESS"))

	// db := db.NewMongoClient()
	// bucket, err := gridfs.NewBucket(db.Database("test_db"))
	// if err != nil {
	// 	panic(err)
	// }
	// file, err := os.Open("CV adapt.pdf")
	// uploadOpts := options.GridFSUpload().SetMetadata(bson.M{"user": 893838})
	// objectID, err := bucket.UploadFromStream("CV adapt.pdf", io.Reader(file), uploadOpts)
	// fmt.Println("id is ", objectID)
	// if err != nil {
	// 	panic(err)
	// }
	// filter := bson.M{"metadata": bson.M{"User": 893838}}
	// type gridfsFile struct {
	// 	Name     string `bson:"filename"`
	// 	Length   int64  `bson:"length"`
	// 	Metadata struct {
	// 		User int64 `bson:"User"`
	// 	} `bson:"metadata"`
	// }
	// cursor, err := bucket.Find(filter)
	// if err != nil {
	// 	panic(err)
	// }
	// var foundFiles []gridfsFile
	// if err = cursor.All(context.TODO(), &foundFiles); err != nil {
	// 	panic(err)
	// }
	// for _, file := range foundFiles {
	// 	fmt.Printf("filename: %s, length: %d\n, Name:%d", file.Name, file.Length, file.Metadata.User)
	// }

}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
