package runner

import (
	"poseidon/jobs"
	"poseidon/model"
	"time"
)

var cronTicker *time.Ticker

//JobMaster is the main job distributor
var JobMaster *model.Master

//LaunchRunner instantiates the ticker and defines the jobs to be done
func LaunchRunner() {
	cronTicker = time.NewTicker(10000 * time.Millisecond)
	JobMaster = SpawnNewMaster(256)
	go func() {
		for {
			select {
			case <-cronTicker.C:
				JobMaster.AddJob(jobs.ParseFeedsJob.AddPayloadAndReturn(map[string]string{"URL": "hello"}))
			}
		}
	}()
}
