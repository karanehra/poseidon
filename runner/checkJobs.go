package runner

import (
	"context"
	"log"
	"poseidon/db"
	"poseidon/models"
	"poseidon/util"

	"go.mongodb.org/mongo-driver/bson"
)

func checkJobs() ([]models.Job, error) {
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
	var results []models.Job
	err = cur.All(context.TODO(), &results)

	return results, err
}
