package jobs

import (
	"log"
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
		log.Fatal(ok)
	} else {
		util.ParseFeedURL(queryPayload["URL"])
	}
	return
}
