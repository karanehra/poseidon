package jobs

import (
	"poseidon/model"
)

//UpdateFeedsJob describes the feed updater job
var UpdateFeedsJob *model.Job = &model.Job{
	Name:     "Update Feeds",
	Executer: updateFeeds,
}

//UpdateFeeds triggers a feed recheck
func updateFeeds(payload interface{}) {
	//getting all feeds logic
	feeds := []string{"https://timesofindia.indiatimes.com/rssfeedstopstories.cms"}

	for i := range feeds {
		payload := map[string]string{"URL": feeds[i]}
		var updateJob *model.Job = ParseFeedJob.AddPayloadAndReturn(payload)
		JobMaster.AddJob(updateJob)
	}
}
