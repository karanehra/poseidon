package runner

import (
	"fmt"
	"poseidon/models"
)

var jobMap map[string]interface{} = map[string]interface{}{
	"ADD_FEEDS":    addFeedsJob,
	"UPDATE_FEEDS": updateFeedsJob,
	"CREATE_TAGS":  createTags,
}

func executeJob(job models.Job) {
	err := job.UpdateStatus("RUNNING")
	if err != nil {
		fmt.Println("Error during update")
	} else {
		funcMappedToJob := jobMap[job.Name]
		if funcMappedToJob != nil {
			go funcMappedToJob.(func(models.Job))(job)
		} else {
			err = job.UpdateStatus("FAILED")
			if err != nil {
				fmt.Println("Invalid Job found. Failing Job.")
			}
		}
	}
}
