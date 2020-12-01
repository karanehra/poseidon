package runner

import (
	"context"
	"fmt"
	"poseidon/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jobMap map[string]interface{} = map[string]interface{}{
	"ADD_FEEDS":    addFeedsJob,
	"UPDATE_FEEDS": updateFeedsJob,
}

func addFeedsJob() {
	fmt.Println("Add feed job executed")
}

func updateFeedsJob() {
	fmt.Println("Update feed job executed")
}

func executeJob(job primitive.M) {
	coll := db.Instance.Collection("jobs")
	filter := bson.M{"_id": job["_id"]}

	update := bson.M{
		"$set": bson.M{"status": "RUNNING"},
	}
	data := &bson.D{}
	decodeError := coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(data)
	if decodeError != nil {
		fmt.Println("Error during update")
	} else {
		funcMappedToJob := jobMap[job["name"].(string)]
		if funcMappedToJob != nil {
			go funcMappedToJob.(func())()
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
					fmt.Printf("Jobs found: %d,\n", len(availableJobs))
					for _, job := range availableJobs {
						go executeJob(job)
					}
				}
			}
		}
	}()
}
