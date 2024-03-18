package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// client := db.NewMongoClient()

	app := fiber.New()

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
