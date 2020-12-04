package runner

import (
	"context"
	"fmt"
	"log"
	"poseidon/db"
	"poseidon/util"

	"go.mongodb.org/mongo-driver/bson"
)

func updateFeedsJob() {
	rssFeedsColl := db.Instance.Collection("rssFeeds")
	cur, err := rssFeedsColl.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	var feeds []bson.M
	err = cur.All(context.TODO(), &feeds)

	if len(feeds) > 0 {
		for _, feed := range feeds {
			saveItemsFromRssFeed(feed)
		}
	}
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
					{"feedID", feedID.(string)},
				}
				articlesPayload = append(articlesPayload, payload)
			}
		}
	}

	res, err := articleColl.InsertMany(context.TODO(), articlesPayload)
	if err != nil {
		fmt.Println("Error While Article insert", err)
	} else {
		fmt.Println(res)
	}
}
