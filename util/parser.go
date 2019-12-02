package util

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

//const a string = "https://timesofindia.indiatimes.com/rssfeedstopstories.cms"

//ParseFeedURL uses gofeed to fetch the rss feed contents
func ParseFeedURL(url string) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)
	fmt.Printf("Parsed new feed %v\n", feed)
}
