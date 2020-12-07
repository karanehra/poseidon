package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"poseidon/db"
	"poseidon/models"
	"poseidon/util"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

func addFeedsJob(job models.Job) {
	wg := sync.WaitGroup{}

	if job.Parameters != "" {
		params := []byte(job.Parameters)

		type FeedData struct {
			Feeds []string `json:"feeds"`
		}

		data := FeedData{}

		err := json.Unmarshal(params, &data)

		if err == nil {
			if len(data.Feeds) > 0 {
				for _, v := range data.Feeds {
					wg.Add(1)
					go addRssFeedToSources(v, &wg)
				}

				wg.Wait()

			}
		}
	}
}

func doesRssFeedExist(url string) bool {
	rssFeedsColl := db.Instance.Collection("rssFeeds")

	result := rssFeedsColl.FindOne(context.TODO(), bson.D{})
	return result.Err() == nil
}

func addRssFeedToSources(url string, wg *sync.WaitGroup) {
	if !doesRssFeedExist(url) {
		data, err := util.ParseFeedURL(url, "")
		if err != nil {
			fmt.Println(err)
		} else {
			rssFeedDocument := bson.D{{"title", data.Title}, {"description", data.Description}, {"url", url}}
			rssFeedsColl := db.Instance.Collection("rssFeeds")
			rssFeedsColl.InsertOne(context.TODO(), rssFeedDocument)
		}
	}
	wg.Done()
}
