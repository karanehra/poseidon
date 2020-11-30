package runner

import (
	"context"
	"fmt"
	"poseidon/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var JobRunner instance = instance{}

//instance is used to define
type instance struct {
	Jobs []primitive.M
}

func (runner *instance) queueJob(jobData primitive.M) {
	runner.Jobs = append(runner.Jobs, jobData)
	coll := db.Instance.Collection("jobs")
	coll.FindOneAndUpdate(context.TODO(), bson.D{{"_id", jobData["_id"]}}, bson.D{{"status": "RUNNING"}})
	fmt.Printf("Queued a job to runner. Total %d \n", len(runner.Jobs))
}

var JobNotificationChan chan int = make(chan int)

//InitializeJobMaster starts the process of watching/checking for jobs
func InitializeJobMaster() {
	checkingTicker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			select {
			case <-checkingTicker.C:
				fmt.Println("Checking for jobs")
				availableJobs, err := checkJobs()
				if err != nil {
					fmt.Println("Unable to check for jobs")
				} else {
					fmt.Printf("%d,\n", len(availableJobs))
					for _, job := range availableJobs {
						JobRunner.queueJob(job)
					}
				}
			}
		}
	}()
}
