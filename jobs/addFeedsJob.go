package jobs

import (
	"context"
	"fmt"
	"juno/database"
	"poseidon/logger"
	"poseidon/util"

	"go.mongodb.org/mongo-driver/bson"
)

//AddFeedsJob parses various sources for urls and parses them
func AddFeedsJob() {
	logger := &logger.Logger{}
	logger.INFO("Starting Add Feeds Job")
	urls, err := util.ParseCSVForURLs("test.csv")
	if err != nil {
		logger.ERROR("Cant parse csv file")
		return
	}
	logger.INFO(fmt.Sprintf("Found %v URLS in sources", len(urls)))

	logger.INFO("Deduping feed URLS")
	var newFeeds = []string{}
	for i := range urls {
		if !doesFeedExist(urls[i]) {
			newFeeds = append(newFeeds, urls[i])
		}
	}
	logger.INFO(fmt.Sprintf("Dedupe yeilded %v new URLS in sources", len(urls)))

	for i := range newFeeds {
		data, err := util.ParseFeedURL(newFeeds[i])
		if err != nil {
			logger.ERROR("Cant parse URL")
			return
		}
		feedData := map[string]string{}
	}
}

func doesFeedExist(URL string) bool {
	coll := database.DB.Collection("feeds")
	result := coll.FindOne(context.TODO(), bson.M{"URL": URL})
	if result.Err() != nil {
		return false
	}
	return true
}
