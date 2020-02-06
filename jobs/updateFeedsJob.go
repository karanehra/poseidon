package jobs

import (
	"context"
	"fmt"
	"poseidon/db"
	"poseidon/logger"
	"poseidon/util"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/karanehra/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//UpdateFeedsJob parses urls taken from a local csv and aggregates a data
func UpdateFeedsJob() {

	logger := &logger.Logger{}

	logger.INFO("Starting Update Job...")
	logger.DepthIn()

	feeds, err := schemas.GetFeeds(db.DB, bson.D{})
	fmt.Println(feeds)

	logger.INFO("Deduping feed URLS")

	logger.SUCCESS(fmt.Sprintf("Found %v URLs...", len(feeds)))
	logger.DepthIn()

	var articleData = []map[string]string{}
	var articleCount int

	for i := range feeds {
		feed := feeds[i]
		logger.INFO(fmt.Sprintf("Begin parse %v...", feed.URL))

		data, err := util.ParseFeedURL(feed.URL)
		if err != nil {
			logger.ERROR(fmt.Sprintf("Can't parse %v!", feed.URL))
			continue
		}

		logger.SUCCESS(fmt.Sprintf("Parsed %v articles", len(data.Items)))

		var feedArticles []map[string]string = []map[string]string{}

		for j := range data.Items {
			payload := map[string]string{
				"feedTitle":       data.Title,
				"feedDescription": data.Description,
				"feedURL":         feed.URL,
				"title":           util.StripHTMLTags(data.Items[j].Title),
				"content":         util.StripHTMLTags(data.Items[j].Content),
				"description":     util.StripHTMLTags(data.Items[j].Description),
				"updated":         strconv.Itoa(int(time.Now().Unix() * 1000)),
				"created":         strconv.Itoa(int(time.Now().Unix() * 1000)),
				"URL":             util.StripHTMLTags(data.Items[j].Link),
			}
			feedArticles = append(feedArticles, payload)
			articleCount++
		}

		articleData = append(articleData, feedArticles...)
	}

	logger.DepthOut()
	logger.INFO(fmt.Sprintf("Created data payload for %v articles...", articleCount))
	logger.INFO("Finding Mongo Collection...")

	coll := db.DB.Collection("articles")

	var documents []interface{} = []interface{}{}

	logger.INFO("Deduping and Creating Bulk Insert Payload..")

	for i := range articleData {
		if !doesArticleExist(util.CreateHashSHA(articleData[i]["URL"]), coll) {
			documents = append(documents, bson.M{
				"feedTitle":       articleData[i]["feedTitle"],
				"feedDescription": articleData[i]["feedDescription"],
				"feedURL":         articleData[i]["feedURL"],
				"title":           articleData[i]["title"],
				"content":         articleData[i]["content"],
				"description":     articleData[i]["description"],
				"updated":         articleData[i]["updated"],
				"created":         articleData[i]["created"],
				"URL":             articleData[i]["URL"],
				"urlHash":         util.CreateHashSHA(articleData[i]["URL"]),
			})
		}
	}

	var tagData map[string]int32 = map[string]int32{}

	for i := range articleData {
		article := articleData[i]
		title := strings.ReplaceAll(article["title"], ".", " ")
		description := strings.ReplaceAll(article["description"], ".", " ")
		updateTagDataFromString(title+" "+description, tagData)
	}

	// err = CacheClient.Set("POSEIDON_ARTICLE_TAGS", tagData)

	if err != nil {
		logger.ERROR("Cannot set values in cache")
	}

	logger.SUCCESS("Stored tagset to cache")
	if len(documents) == 0 {
		logger.INFO("No new articles...")
		logger.DepthOut()
		logger.INFO("Task Finished!")
		return
	}
	logger.INFO(fmt.Sprintf("Created %v Articles Payload. Writing to DB...", len(documents)))
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	_, error := coll.InsertMany(ctx, documents)
	if error != nil {
		logger.ERROR("Error occured during database write...")
	}
	logger.SUCCESS(fmt.Sprintf("Wrote %v articles to DB", len(documents)))
	logger.DepthOut()
	logger.INFO("Task Finished!")
}

func updateTagDataFromString(data string, tagData map[string]int32) map[string]int32 {
	tags := getStringTags(data)
	for i := range tags {
		tagData[tags[i]]++
	}
	return tagData
}

func doesArticleExist(hash string, coll *mongo.Collection) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	result := coll.FindOne(ctx, bson.M{"urlHash": hash})
	if result.Err() != nil {
		return false
	}
	return true
}

func getStringTags(data string) []string {
	words := strings.Split(data, " ")
	tags := []string{}
	for i := range words {
		word := []rune(words[i])
		if len(word) > 0 && unicode.IsUpper(word[0]) {
			tags = append(tags, words[i])
		}
	}
	return cleanTags(tags)
}

func cleanTags(tags []string) []string {
	var clean = []string{}
	for i := range tags {
		if len(tags[i]) == 1 {
			continue
		}
		tag := strings.ToLower(tags[i])
		regex := regexp.MustCompile("([[:punct:]])")
		tag = regex.ReplaceAllLiteralString(tag, "")
		tag = strings.ReplaceAll(tag, "’s", "")
		if !isTagRejected(tag) {
			clean = append(clean, tag)
		}
	}
	return clean
}

func isTagRejected(tag string) bool {
	return util.RejectedTagsMap[tag]
}
