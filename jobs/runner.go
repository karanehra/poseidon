package jobs

import (
	"log"
	"poseidon/cache"
	"time"
)

//CacheClient is the central cache client
var CacheClient *cache.Client

//LaunchRunner instantiates the ticker and defines the jobs to be done
func LaunchRunner() {
	cronTicker := time.NewTicker(20 * time.Minute)
	processTicker := time.NewTicker(10 * time.Second)
	CacheClient = &cache.Client{
		Port:    3009,
		BaseURL: "http://localhost",
	}
	err := CacheClient.Create()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case <-cronTicker.C:
				AddFeedsJob()
			case <-processTicker.C:
				CheckForProcesses()
			}
		}
	}()
}
