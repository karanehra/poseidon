package jobs

import (
	"fmt"
	"poseidon/cache"
	"poseidon/model"
	"time"
)

var cronTicker *time.Ticker

//JobMaster is the main job distributor
var JobMaster *model.Master

//CacheClient is the central cache client
var CacheClient *cache.Client

//LaunchRunner instantiates the ticker and defines the jobs to be done
func LaunchRunner() {
	cronTicker = time.NewTicker(10000 * time.Millisecond)
	JobMaster = SpawnNewMaster(256)
	CacheClient = &cache.Client{
		BaseURL: "http://localhost",
		Port:    3009,
	}
	err := CacheClient.Create()
	if err != nil {
		fmt.Println("Cant connect to cache.")
	}
	go func() {
		for {
			select {
			case <-cronTicker.C:
				JobMaster.AddJob(UpdateFeedsJob.AddPayloadAndReturn(map[string]string{"URL": "hello"}))
			}
		}
	}()
}
