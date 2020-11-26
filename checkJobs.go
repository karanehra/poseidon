package main

import (
	"context"
	"fmt"
	"log"
	"poseidon/db"
	"poseidon/util"

	"go.mongodb.org/mongo-driver/bson"
)

func checkJobs() {
	collections, err := db.Instance.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if !util.SliceContains(collections, "jobs") {
		log.Fatal("Cannot Find Jobs Collection in DB")
	}

	jobsColl := db.Instance.Collection("jobs")

	cur, err := jobsColl.Find(context.TODO(), bson.D{{"status", "QUEUED"}})
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M
	err = cur.All(context.TODO(), &results)
	fmt.Printf("%d jobs available", len(results))
}
