package jobs

import (
	"fmt"
	"poseidon/model"
	"poseidon/util"
)

//UpdateFeedsJob describes the feed updater job
var UpdateFeedsJob *model.Job = &model.Job{
	Name:     "Update Feeds",
	Executer: updateFeeds,
}

//UpdateFeeds triggers a feed recheck
func updateFeeds(payload interface{}) {
	//getting all feeds logic
	feeds, err := util.ParseCSVForURLs("test.csv")

	if err != nil {
		fmt.Println("An error occured while fetching feeds")
		return
	}

	for i := range feeds {
		payload := map[string]string{"URL": feeds[i]}
		var updateJob *model.Job = ParseFeedJob.AddPayloadAndReturn(payload)
		JobMaster.AddJob(updateJob)
	}
}
