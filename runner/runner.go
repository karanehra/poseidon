package runner

import (
	"fmt"
	"poseidon/model"
	"time"
)

var cronTicker *time.Ticker
var jobMaster *model.Master

//LaunchRunner instantiates the ticker and defines the jobs to be done
func LaunchRunner() {
	cronTicker = time.NewTicker(1000 * time.Millisecond)
	jobMaster = SpawnNewMaster(256)
	go func() {
		for {
			select {
			case t := <-cronTicker.C:
				fmt.Printf("Ticked at: %v\n", t.UnixNano())
			}
		}
	}()
}
