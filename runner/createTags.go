package runner

import (
	"context"
	"fmt"
	"poseidon/db"
	"poseidon/models"

	"go.mongodb.org/mongo-driver/bson"
)

func createTags(job models.Job) {
	job.AddLog("Starting...")

	articlesCursor, err := db.Instance.Collection("articles").Find(context.TODO(), bson.D{})

	if err != nil {
		job.AddLog("Job failed due to error")
		job.UpdateStatus("FAILED")
	} else {
		var article bson.M
		for articlesCursor.Next(context.TODO()) {
			err := articlesCursor.Decode(&article)
			if err == nil {
				fmt.Println(article["_id"])
			}
		}
	}

	job.AddLog("Finished")
}
