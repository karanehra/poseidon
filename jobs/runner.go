package jobs

import (
	"log"
	"poseidon/cache"
	"poseidon/util"
	"time"
)

//CacheClient is the central cache client
var CacheClient *cache.Client

//UaMasterList contains lots of user agents for scraping rotation
var UaMasterList []string

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
	go UpdateFeedsJob()

	UaMasterList, _ = util.ParseCSVForUAs("ca.csv")

	updateTicker := time.NewTicker(30 * time.Minute)
	processTicker := time.NewTicker(5 * time.Minute)

	go func() {
		for {
			select {
			case <-updateTicker.C:
				go UpdateFeedsJob()
			case <-processTicker.C:
				go CheckForProcesses()
			}
		}
	}()
}
