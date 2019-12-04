package jobs

import (
	"poseidon/model"
	"time"

	"github.com/mmcdole/gofeed"
)

//FeedExtractorJob gets out all information out of the feed
var FeedExtractorJob *model.Job = &model.Job{
	Name:     "Exractor",
	Executer: extractFeed,
}

func extractFeed(payload interface{}) {
	feed := payload.(*gofeed.Feed)
	data := map[string]interface{}{}
	data["title"] = feed.Title
	data["description"] = feed.Description
	data["link"] = feed.Link
	if feed.UpdatedParsed != nil {
		data["updated"] = feed.UpdatedParsed
	} else {
		data["updated"] = time.Now()
	}
	for i := range feed.Items {
		CacheClient.Set(feed.Items[i].Title, feed.Items[i])
	}
}
