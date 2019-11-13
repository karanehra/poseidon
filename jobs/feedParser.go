package jobs

import "poseidon/model"

import "fmt"

//ParseFeedsJob routinely checks feed urls for updates
var ParseFeedsJob *model.Job = &model.Job{
	Name:    "Parse Feeds",
	Execute: parseFeeds,
}

func parseFeeds() {
	fmt.Println("Parsing....")
}
