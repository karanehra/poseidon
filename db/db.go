package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DB is a mongo database instance
var DB *mongo.Database

func init() {
	mongoDBUri := os.Getenv("MONGO_DB_URL")
	if mongoDBUri == "" {
		log.Fatal("Env Variable MONGO_DB_URL not specified")
	}

	databaseClientOptions := options.Client().ApplyURI(mongoDBUri)
	mongoClient, err := mongo.NewClient(databaseClientOptions)
	if err != nil {
		log.Fatal(err)
	}
	context, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = mongoClient.Connect(context)
	if err != nil {
		log.Fatal(err)
	}
	mongoDBName := os.Getenv("MONGO_DB_NAME")
	if mongoDBName == "" {
		log.Fatal("Env Variable MONGO_DB_NAME not specified")
	}
	DB = mongoClient.Database(mongoDBName)
	fmt.Println("Database Connection Success")
}
