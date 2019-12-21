package jobs

import (
	"context"
	"fmt"
	"poseidon/db"
	"poseidon/util"
	"time"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
)

//UpdateFeedsJob parses urls taken from a local csv and aggregates a data
func UpdateFeedsJob() {
	color.Yellow("[INFO]:Starting Feed Update...")
	urls, err := util.ParseCSVForURLs("test.csv")
	if err != nil {
		color.Red("== [ERROR]:No URLs found!")
		return
	}
	color.Green("== [INFO]:Found %v URLs...", len(urls))
	var articleData []map[string]string = []map[string]string{}
	var articleCount int
	for i := range urls {
		fmt.Println("")
		color.Yellow("==== [INFO]: Begin parse %v...", urls[i])
		data, err := util.ParseFeedURL(urls[i])
		if err != nil {
			color.Red("==== [ERROR]: Can't parse %v!", urls[i])
		}
		color.Green("==== [INFO]: Parsed %v articles", len(data.Items))
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
			}
			feedArticles = append(feedArticles, payload)
			articleCount++
		}

		articleData = append(articleData, feedArticles...)
	}
	color.Yellow("== [INFO]: Created data payload for %v articles...", articleCount)
	color.Yellow("== [INFO]: Finding Mongo Collection...")
	coll := db.DB.Collection("x")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	var documents []interface{} = []interface{}{}
	color.Yellow("\n== [INFO]: Creating Bulk Insert Payload...")
	for i := range articleData {
		documents = append(documents, bson.M{
			"feedTitle":       articleData[i]["feedTitle"],
			"feedDescription": articleData[i]["feedDescription"],
			"feedURL":         articleData[i]["feedURL"],
			"title":           articleData[i]["title"],
			"content":         articleData[i]["content"],
			"description":     articleData[i]["description"],
			"updated":         articleData[i]["updated"],
		})
	}
	color.Yellow("== [INFO]: Created %v Articles Payload. Writing to DB...", len(documents))
	_, error := coll.InsertMany(ctx, documents)
	if error != nil {
		color.Red("== [ERROR]: Error occured during database write...")
	}
	color.Green("== [SUCCESS]: Wrote %v articles to DB", len(documents))
	color.Yellow("[INFO]: Task Finished!")
}
