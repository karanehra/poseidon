package jobs

import "poseidon/model"

import "fmt"

//ParseFeedsJob routinely checks feed urls for updates
var ParseFeedsJob *model.Job = &model.Job{
	Name:     "Parse Feeds",
	Executer: parseFeeds,
}

func parseFeeds(payload interface{}) {
	fmt.Printf("Parsing.... %v \n", payload)
	return
}
