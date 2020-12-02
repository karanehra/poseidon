package runner

import (
	"fmt"
	"time"
)

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
