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
