package jobs

import (
	"encoding/csv"
	"fmt"
	"os"
	"poseidon/db"
	"poseidon/logger"
	"strconv"
	"strings"

	"github.com/karanehra/schemas"
	"go.mongodb.org/mongo-driver/bson"
)

//DumpFeedsJob gets feeds out of DB and dumps them in a csv
func DumpFeedsJob() {
	logger := logger.Logger{}
	logger.INFO("Starting dump feeds job")
	feeds, err := schemas.GetFeeds(db.DB, bson.D{})
	fmt.Println(feeds)
	if err != nil {
		logger.ERROR("Unable to fetch feed data")
	}
	csvFile, err := os.Create("feed-dump.csv")
	if err != nil {
		logger.ERROR("Unable to fetch feed data")
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()
	for i := range feeds {
		dataRow := []string{
			feeds[i].Title,
			feeds[i].URL,
			strings.Join(feeds[i].Tags, "|"),
		}
		err := writer.Write(dataRow)
		if err != nil {
			logger.ERROR("Unable to fetch feed data")
		}
	}
	logger.SUCCESS("Job Done")
}

//DumpArticlesJob dumps articles in a csv
func DumpArticlesJob() {
	logger := logger.Logger{}
	logger.INFO("Starting dump articles job")
	articles, err := schemas.GetArticles(db.DB, bson.D{})
	if err != nil {
		logger.ERROR("Unable to fetch article data")
	}
	csvFile, err := os.Create("article-dump.csv")
	if err != nil {
		logger.ERROR("Unable to fetch article data")
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()
	headerRow := []string{
		"title",
		"url",
		"description",
		"content",
		"feedDescription",
		"feedTitle",
		"feedUrl",
	}
	err = writer.Write(headerRow)
	if err != nil {
		logger.ERROR("Unable to fetch article data")
	}
	for i := range articles {
		dataRow := []string{
			articles[i].Title,
			articles[i].URL,
			articles[i].Description,
			articles[i].Content,
			articles[i].FeedDescription,
			articles[i].FeedTitle,
			articles[i].FeedURL,
			strconv.Itoa(int(articles[i].CreatedAt)),
		}
		err = writer.Write(dataRow)
		if err != nil {
			logger.ERROR("Unable to fetch article data")
		}
	}
	logger.SUCCESS("Job Done")
}
