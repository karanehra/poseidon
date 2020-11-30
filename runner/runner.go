package runner

import (
	"fmt"
	"time"
)

type JobRunner struct {
	Jobs []string
}

var JobNotificationChan chan int = make(chan int)

func init() {
	checkingTicker := time.NewTicker(1 * time.Minute)

	go func() {
		for {
			select {
			case <-JobNotificationChan:
				fmt.Println("Received a new Job")
			case <-checkingTicker.C:
				fmt.Println("Checking for job based on timer")
			}
		}
	}()
}
