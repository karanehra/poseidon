package services

import (
	"context"
	"poseidon/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SetJobStatusInDB marks the given job as done. `_id` is needed to perform this operation
func SetJobStatusInDB(job primitive.M, status string) error {
	coll := db.Instance.Collection("jobs")
	filter := bson.M{"_id": job["_id"]}

	update := bson.M{
		"$set": bson.M{"status": status},
	}
	data := &bson.D{}
	decodeError := coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(data)
	return decodeError
}

//AppendLogsToJob is used to add log info to the given job. Multiple must be seperated by a newline
func AppendLogsToJob(job primitive.M, log string) error {
	coll := db.Instance.Collection("jobs")
	filter := bson.M{"_id": job["_id"]}

	currentLog := job["log"]

	if currentLog == nil {
		currentLog = ""
	}

	update := bson.M{
		"$set": bson.M{"log": currentLog.(string) + "\n" + log},
	}
	data := &bson.D{}
	decodeError := coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(data)
	return decodeError
}
