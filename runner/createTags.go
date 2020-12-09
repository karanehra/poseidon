package runner

import (
	"context"
	"fmt"
	"poseidon/db"
	"poseidon/models"

	"github.com/jdkato/prose"
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
				title := article["title"]
				description := article["description"]

				var tags []string

				if title != nil {
					doc, _ := prose.NewDocument(title.(string))
					for _, entity := range doc.Entities() {
						tags = append(tags, entity.Text)
					}
				}

				if description != nil {
					doc, _ := prose.NewDocument(description.(string))
					for _, entity := range doc.Entities() {
						tags = append(tags, entity.Text)
					}
				}

				fmt.Println(tags)
			}
		}
	}

	job.AddLog("Finished")
}
