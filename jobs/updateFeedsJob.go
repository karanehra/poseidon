package jobs

import (
	"fmt"
	"poseidon/util"

	"github.com/fatih/color"
)

//UpdateFeedsJob parses urls taken from a local csv and aggregates a data
func UpdateFeedsJob() {
	color.Yellow("[INFO]:Starting Feed Update...")
	urls, err := util.ParseCSVForURLs("test.csv")
	if err != nil {
		color.Red("== [ERROR]:No URLs found!")
		return
	}
	color.Green("== [INFO]:Found %v URLs", len(urls))
	var articleData []map[string]string = []map[string]string{}
	var articleCount int
	for i := range urls {
		fmt.Println("")
		color.Yellow("==== [INFO]: Begin parse %v", urls[i])
		data, err := util.ParseFeedURL(urls[i])
		if err != nil {
			color.Red("==== [ERROR]: Can't parse %v", urls[i])
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
	color.Yellow("\n== [INFO]: Created data payload for %v articles", articleCount)
}
