package jobs

import (
	"log"
	"poseidon/model"
	"poseidon/util"
)

type feedParsePayload struct {
	URL string
}

//ParseFeedsJob routinely checks feed urls for updates
var ParseFeedsJob *model.Job = &model.Job{
	Name:     "Parse Feeds",
	Executer: parseFeeds,
}

func parseFeeds(payload interface{}) {
	queryPayload, ok := payload.(map[string]string)
	if !ok {
		log.Fatal(ok)
	} else {
		util.ParseFeedURL(queryPayload["URL"])
	}
	return
}
