package runner

import (
	"poseidon/jobs"
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
			case <-cronTicker.C:
				jobMaster.AddJob(jobs.ParseFeedsJob)
			}
		}
	}()
}
