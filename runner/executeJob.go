package runner

import (
	"context"
	"fmt"
	"poseidon/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jobMap map[string]interface{} = map[string]interface{}{
	"ADD_FEEDS":    addFeedsJob,
	"UPDATE_FEEDS": updateFeedsJob,
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
