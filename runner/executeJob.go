package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"poseidon/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jobMap map[string]interface{} = map[string]interface{}{
	"ADD_FEEDS":    addFeedsJob,
	"UPDATE_FEEDS": updateFeedsJob,
}

func addFeedsJob(jobInfo primitive.M) {
	params := []byte(jobInfo["parameters"].(string))

	type FeedData struct {
		Feeds []string `json:"feeds"`
	}

	data := FeedData{}

	err := json.Unmarshal(params, &data)

	if err == nil {
		if len(data.Feeds) > 0 {
			for _, v := range data.Feeds {
				go addRssFeedToSources(v)
			}
		}
	} else {
		fmt.Println("Incorrect parameters for ADD_FEEDS job")
	}
}

func addRssFeedToSources(url string) {
	fmt.Println(url)
}

func updateFeedsJob() {
	fmt.Println("Update feed job executed")
}

func executeJob(job primitive.M) {
	coll := db.Instance.Collection("jobs")
	filter := bson.M{"_id": job["_id"]}

	update := bson.M{
		"$set": bson.M{"status": "QUEUED"},
	}
	data := &bson.D{}
	decodeError := coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(data)
	if decodeError != nil {
		fmt.Println("Error during update")
	} else {
		funcMappedToJob := jobMap[job["name"].(string)]
		if funcMappedToJob != nil {
			go funcMappedToJob.(func(primitive.M))(job)
		} else {
			update = bson.M{
				"$set": bson.M{"status": "FAILED"},
			}
			data = &bson.D{}
			decodeError = coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(data)
			if decodeError == nil {
				fmt.Println("Invalid Job found. Failing Job.")
			}
		}
	}
}
