package jobs

import (
	"fmt"
	"poseidon/model"
	"poseidon/util"
)

type feedParsePayload struct {
	URL string
}

//ParseFeedJob routinely checks feed urls for updates
var ParseFeedJob *model.Job = &model.Job{
	Name:     "Parse Feeds",
	Executer: parseFeed,
}

func parseFeed(payload interface{}) {
	queryPayload, ok := payload.(map[string]string)
	if !ok {
		fmt.Println("Problem in payload casting")
		return
	}
	feed, _ := util.ParseFeedURL(queryPayload["URL"])

	var updateJob *model.Job = FeedExtractorJob.AddPayloadAndReturn(feed)
	JobMaster.AddJob(updateJob)

	return
}
