package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Instance is our main mongo database instance
var Instance *mongo.Database

//InitializeDatabase connects to given mongoDB instance and makes it available for the application
func InitializeDatabase() {
	databaseClientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
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
	Instance = mongoClient.Database("brutus")
	data, err := Instance.Client().ListDatabases(context, bson.D{})
	// data, err := DB.ListCollectionNames(context, bson.D{})
	fmt.Println(data)
	fmt.Println("Database Connection Success")
}
