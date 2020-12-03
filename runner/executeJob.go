package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"poseidon/db"
	"poseidon/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jobMap map[string]interface{} = map[string]interface{}{
	"ADD_FEEDS":    addFeedsJob,
	"UPDATE_FEEDS": updateFeedsJob,
}

func addFeedsJob(jobInfo primitive.M) {
	params := []byte(jobInfo["parameters"].(string))

	type FeedData struct {
		Feeds []string `json:"feeds"`
	}

	data := FeedData{}

	err := json.Unmarshal(params, &data)

	if err == nil {
		if len(data.Feeds) > 0 {
			for _, v := range data.Feeds {
				go addRssFeedToSources(v)
			}
		}
	} else {
		fmt.Println("Incorrect parameters for ADD_FEEDS job")
	}
}

func doesRssFeedExist(url string) bool {
	rssFeedsColl := db.Instance.Collection("rssFeeds")

	result := rssFeedsColl.FindOne(context.TODO(), bson.D{})
	return result.Err() == nil
}

func addRssFeedToSources(url string) {
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
}

func updateFeedsJob() {
	rssFeedsColl := db.Instance.Collection("rssFeeds")
	cur, err := rssFeedsColl.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M
	err = cur.All(context.TODO(), &results)

	var urls []string
	for _, v := range results {
		url := v["url"]
		if url != nil {
			urls = append(urls, url.(string))
		}
	}

	fmt.Printf("Total URLs found, %d", len(urls))
}

func executeJob(job primitive.M) {
	coll := db.Instance.Collection("jobs")
	filter := bson.M{"_id": job["_id"]}

	update := bson.M{
		"$set": bson.M{"status": "QUEUED"},
	}
	data := &bson.D{}
	decodeError := coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(data)
	if decodeError != nil {
		fmt.Println("Error during update")
	} else {
		funcMappedToJob := jobMap[job["name"].(string)]
		if funcMappedToJob != nil {
			go funcMappedToJob.(func(primitive.M))(job)
		} else {
			update = bson.M{
				"$set": bson.M{"status": "FAILED"},
			}
			data = &bson.D{}
			decodeError = coll.FindOneAndUpdate(context.TODO(), filter, update).Decode(data)
			if decodeError == nil {
				fmt.Println("Invalid Job found. Failing Job.")
			}
		}
	}
}
