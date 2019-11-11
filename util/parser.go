package util

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

//ParseFeedURL uses gofeed to fetch the rss feed contents
func ParseFeedURL() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://timesofindia.indiatimes.com/rssfeedstopstories.cms")
	items := feed.Items
	for i := range items {
		fmt.Println(items[i].Title)
		fmt.Println(items[i].Description)
		fmt.Println(StripHTMLTags(items[i].Content))
		fmt.Println(items[i].Updated)
	}
}
