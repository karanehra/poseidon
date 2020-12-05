package runner

import (
	"fmt"
	"poseidon/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jobMap map[string]interface{} = map[string]interface{}{
	"ADD_FEEDS":    addFeedsJob,
	"UPDATE_FEEDS": updateFeedsJob,
}

func executeJob(job primitive.M) {
	err := services.SetJobStatusInDB(job, "RUNNING")
	if err != nil {
		fmt.Println("Error during update")
	} else {
		funcMappedToJob := jobMap[job["name"].(string)]
		if funcMappedToJob != nil {
			go funcMappedToJob.(func(primitive.M))(job)
		} else {
			err = services.SetJobStatusInDB(job, "FAILED")
			if err == nil {
				fmt.Println("Invalid Job found. Failing Job.")
			}
		}
	}
}
