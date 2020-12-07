package runner

import (
	"context"
	"fmt"
	"log"
	"poseidon/db"
	"poseidon/models"
	"poseidon/util"

	"go.mongodb.org/mongo-driver/bson"
)

func updateFeedsJob(job models.Job) {
	job.AddLog("Starting...")

	rssFeedsColl := db.Instance.Collection("rssFeeds")
	cur, err := rssFeedsColl.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	var feeds []bson.M
	err = cur.All(context.TODO(), &feeds)

	if len(feeds) > 0 {
		job.AddLog(fmt.Sprintf("Saving items from %d feeds", len(feeds)))

		for _, feed := range feeds {
			saveItemsFromRssFeed(feed)
		}
	}
	job.UpdateStatus("FINISHED")
}

func saveItemsFromRssFeed(feedData bson.M) {
	feedURL := feedData["url"]
	feedID := feedData["_id"]

	articleColl := db.Instance.Collection("articles")
	articlesPayload := []interface{}{}

	if feedURL != nil {
		data, err := util.ParseFeedURL(feedURL.(string), "")
		if err != nil {
			fmt.Println("Error while getting rss data")
		} else {
			items := data.Items

			for _, v := range items {
				payload := bson.D{
					{"title", v.Title},
					{"description", v.Description},
					{"url", v.Link},
					{"feedID", util.ObjectIDToHexString(feedID)},
				}
				articlesPayload = append(articlesPayload, payload)
			}
		}
	}

	if len(articlesPayload) > 0 {
		res, err := articleColl.InsertMany(context.TODO(), articlesPayload)
		if err != nil {
			fmt.Println("Error While Article insert", err)
		} else {
			fmt.Println(res)
		}
	} else {
		fmt.Println("No articles to insert found")
	}
}
