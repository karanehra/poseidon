package jobs

import (
	"context"
	"fmt"
	"poseidon/db"
	"poseidon/logger"
	"poseidon/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//UpdateFeedsJob parses urls taken from a local csv and aggregates a data
func UpdateFeedsJob() {
	logger := &logger.Logger{}
	logger.INFO("Starting Update Job...")
	logger.DepthIn()
	urls, err := util.ParseCSVForURLs("test.csv")
	if err != nil {
		logger.ERROR("No URLs found!")
		return
	}
	logger.SUCCESS(fmt.Sprintf("Found %v URLs...", len(urls)))
	logger.DepthIn()
	var articleData []map[string]string = []map[string]string{}
	var articleCount int
	for i := range urls {
		logger.INFO(fmt.Sprintf("Begin parse %v...", urls[i]))
		data, err := util.ParseFeedURL(urls[i])
		if err != nil {
			logger.ERROR(fmt.Sprintf("Can't parse %v!", urls[i]))
		}
		logger.SUCCESS(fmt.Sprintf("Parsed %v articles", len(data.Items)))
		var feedArticles []map[string]string = []map[string]string{}
		for j := range data.Items {
			payload := map[string]string{
				"feedTitle":       data.Title,
				"feedDescription": data.Description,
				"feedURL":         urls[i],
				"title":           util.StripHTMLTags(data.Items[j].Title),
				"content":         util.StripHTMLTags(data.Items[j].Content),
				"description":     util.StripHTMLTags(data.Items[j].Description),
				"updated":         data.Items[j].Updated,
				"URL":             util.StripHTMLTags(data.Items[j].Link),
			}
			feedArticles = append(feedArticles, payload)
			articleCount++
		}

		articleData = append(articleData, feedArticles...)
	}
	logger.DepthOut()
	logger.INFO(fmt.Sprintf("Created data payload for %v articles...", articleCount))
	logger.INFO("Finding Mongo Collection...")
	coll := db.DB.Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	var documents []interface{} = []interface{}{}
	logger.INFO("Deduping and Creating Bulk Insert Payload..")
	for i := range articleData {
		if !doesArticleExist(util.CreateHashSHA(articleData[i]["URL"]), coll) {
			documents = append(documents, bson.M{
				"feedTitle":       articleData[i]["feedTitle"],
				"feedDescription": articleData[i]["feedDescription"],
				"feedURL":         articleData[i]["feedURL"],
				"title":           articleData[i]["title"],
				"content":         articleData[i]["content"],
				"description":     articleData[i]["description"],
				"updated":         articleData[i]["updated"],
				"URL":             articleData[i]["URL"],
				"urlHash":         util.CreateHashSHA(articleData[i]["URL"]),
			})
		}
	}
	if len(documents) > 0 {
		logger.INFO(fmt.Sprintf("Created %v Articles Payload. Writing to DB...", len(documents)))
		_, error := coll.InsertMany(ctx, documents)
		if error != nil {
			logger.ERROR("Error occured during database write...")
		}
		logger.SUCCESS(fmt.Sprintf("Wrote %v articles to DB", len(documents)))
	} else {
		logger.INFO("No new articles...")
	}
	logger.DepthOut()
	logger.INFO("[Task Finished!")
}

func doesArticleExist(hash string, coll *mongo.Collection) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	result := coll.FindOne(ctx, bson.M{"urlHash": hash})
	return result != nil
}