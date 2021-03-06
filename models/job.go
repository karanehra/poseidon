package models

import (
	"context"
	"poseidon/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Job defines the schema of a job and used to hold common util methods
type Job struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	Status     string             `json:"status" bson:"status"`
	Log        string             `json:"log" bson:"log"`
	Parameters string             `json:"parameters" bson:"parameters"`
	CreatedAt  primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt  primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
}

func (job *Job) UpdateStatus(status string) error {
	coll := db.Instance.Collection("jobs")
	filter := bson.M{"_id": job.ID}

	update := bson.M{
		"$set": bson.M{"status": status},
	}
	data := &bson.D{}
	decodeError := coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(data)
	return decodeError
}

func (job *Job) AddLog(logLine string) error {
	coll := db.Instance.Collection("jobs")
	filter := bson.M{"_id": job.ID}

	currentLog := job.Log

	update := bson.M{
		"$set": bson.M{"log": currentLog + "\n" + logLine},
	}

	data := &bson.D{}
	decodeError := coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(data)
	return decodeError
}
