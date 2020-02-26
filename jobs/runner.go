package jobs

import (
	"log"
	"poseidon/cache"
)

//CacheClient is the central cache client
var CacheClient *cache.Client

//LaunchRunner instantiates the ticker and defines the jobs to be done
func LaunchRunner() {
	CacheClient = &cache.Client{
		Port:    3009,
		BaseURL: "http://localhost",
	}
	err := CacheClient.Create()
	if err != nil {
		log.Fatal(err)
	}
	UpdateFeedsJob()

	// updateTicker := time.NewTicker(30 * time.Minute)
	// processTicker := time.NewTicker(40 * time.Minute)

	// go func() {
	// 	for {
	// 		select {
	// 		case <-updateTicker.C:
	// 			go UpdateFeedsJob()
	// 		case <-processTicker.C:
	// 			go CheckForProcesses()
	// 		}
	// 	}
	// }()
}
