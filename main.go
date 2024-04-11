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
		userHandler         = api.NewUserHandler(store, "user handler")
		companyHandler      = api.NewCompanyHandler(store, "company handler")
		industryHandler     = api.NewIndustryHandler(store, "industry handler")
		verificationHandler = api.NewVerificationHandler(store, "verification handler")
	)
	app.Post("/user", userHandler.HandlePostUser)
	app.Get("/user/:id", userHandler.HandleGetUser)
	app.Get("/user", userHandler.HandleGetUsers)
	app.Put("/user/:id", userHandler.HandlePut)
	app.Delete("/user/:id", userHandler.HandleDelete)

	app.Post("/company", companyHandler.HandlePostCompany)
	app.Get("/company/:id", companyHandler.HandleGetCompany)
	app.Get("/company", companyHandler.HandleGetCompanies)
	app.Put("/company/:id", companyHandler.HandlePut)
	app.Delete("/company/:id", companyHandler.HandleDelete)

	app.Post("/industry", industryHandler.HandlePostIndustry)
	app.Get("/industry/:id", industryHandler.HandleGetIndustry)
	app.Get("/industry", industryHandler.HandleGetIndustries)
	app.Delete("/industry/:id", industryHandler.HandleDelete)

	app.Get("/verification/:id", verificationHandler.HandleGetVerification)
	app.Get("/verification", verificationHandler.HandleGetVerifications)

	// Administrative handlers
	app.Put("/verification/:id", verificationHandler.HandlePut)
	app.Post("/verification", verificationHandler.HandlePostVerification)
	app.Delete("/verification/:id", verificationHandler.HandleDelete)

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
