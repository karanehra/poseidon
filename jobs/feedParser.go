package jobs

import (
	"fmt"
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
	err := CacheClient.Set("key", "asfsf")
	if err != nil {
		fmt.Printf("Error setting cache val %v\n", err.Error())
	}
	return
}
