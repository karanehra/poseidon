package main

import (
	"context"
	"fmt"
	"log"
	"poseidon/db"

	"go.mongodb.org/mongo-driver/bson"
)

func checkJobs() {
	data, err := db.DB.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for _, val := range data {
		if val == "jobs" {
			fmt.Println("Found collection")
		}
	}
}
