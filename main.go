package main

import (
	"context"
	"fmt"

	"github.com/A3R0-01/head-hunter/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func main() {
	db := db.NewMongoClient()
	bucket, err := gridfs.NewBucket(db.Database("test_db"))
	if err != nil {
		panic(err)
	}
	// file, err := os.Open("CV adapt.pdf")
	// uploadOpts := options.GridFSUpload().SetMetadata(bson.M{"user": 893838})
	// objectID, err := bucket.UploadFromStream("CV adapt.pdf", io.Reader(file), uploadOpts)
	// fmt.Println("id is ", objectID)
	// if err != nil {
	// 	panic(err)
	// }
	filter := bson.M{"metadata": bson.M{"user": 893838}}
	type gridfsFile struct {
		Name     string `bson:"filename"`
		Length   int64  `bson:"length"`
		Metadata struct {
			User int64 `bson:"user"`
		} `bson:"metadata"`
	}
	cursor, err := bucket.Find(filter)
	if err != nil {
		panic(err)
	}
	var foundFiles []gridfsFile
	if err = cursor.All(context.TODO(), &foundFiles); err != nil {
		panic(err)
	}
	for _, file := range foundFiles {
		fmt.Printf("filename: %s, length: %d\n, Name:%d", file.Name, file.Length, file.Metadata.User)
	}

}
