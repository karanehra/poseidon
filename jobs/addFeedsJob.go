package jobs

import (
	"context"
	"fmt"
	"juno/database"
	"poseidon/logger"
	"poseidon/util"

	"github.com/karanehra/schemas"
	"go.mongodb.org/mongo-driver/bson"
)

//AddFeedsJob parses various sources for urls and parses them
func AddFeedsJob() {
	logger := &logger.Logger{}
	logger.INFO("Starting Add Feeds Job")
	urls, err := util.ParseCSVForURLs("new.csv")
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
	logger.INFO(fmt.Sprintf("Dedupe yeilded %v new URLS in sources", len(newFeeds)))

	if len(newFeeds) == 0 {
		logger.INFO(fmt.Sprintf("No new feeds found"))
		return
	}

	var feedDocuments = []interface{}{}
	for i := range newFeeds {
		data, err := util.ParseFeedURL(newFeeds[i])
		if err != nil {
			logger.ERROR("Cant parse URL")
			continue
		}
		feedData := schemas.Feed{}
		if data.Title != "" {
			feedData.Title = data.Title
		}
		feedData.URL = newFeeds[i]
		feedDocuments = append(feedDocuments, feedData)
	}
	coll := database.DB.Collection("feeds")
	_, insertError := coll.InsertMany(context.Background(), feedDocuments)
	if insertError != nil {
		fmt.Println(insertError)
		logger.ERROR("Cant add data to DB")
		return
	}
	logger.INFO("Task completed")
}

func doesFeedExist(URL string) bool {
	coll := database.DB.Collection("feeds")
	result := coll.FindOne(context.TODO(), bson.M{"URL": URL})
	if result.Err() != nil {
		return false
	}
	return true
}
