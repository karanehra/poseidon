package jobs

import (
	"time"
)

var cronTicker *time.Ticker

//CacheClient is the central cache client

//LaunchRunner instantiates the ticker and defines the jobs to be done
func LaunchRunner() {
	cronTicker = time.NewTicker(20 * time.Second)
	go func() {
		for {
			select {
			case <-cronTicker.C:
				UpdateFeedsJob()
			}
		}
	}()
}
