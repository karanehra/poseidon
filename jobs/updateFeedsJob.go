package jobs

import (
	"context"
	"fmt"
	"log"
	"poseidon/db"
	"poseidon/logger"
	"poseidon/util"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/karanehra/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var articleData = make([]map[string]string, 100000)
var articleCount int

//UpdateFeedsJob parses urls taken from a local csv and aggregates a data
func UpdateFeedsJob() {

	logger := &logger.Logger{}

	logger.INFO("Starting Update Job...")
	logger.DepthIn()

	feeds, err := schemas.GetFeeds(db.DB, bson.D{})

	var wg sync.WaitGroup

	for i := range feeds {
		feed := feeds[i]
		wg.Add(1)
		go parseFeedWorker(feed.URL, &wg)
	}

	wg.Wait()
	logger.INFO("WG: finish")

	logger.DepthOut()
	logger.INFO(fmt.Sprintf("Created data payload for %v articles...", articleCount))
	logger.INFO("Finding Mongo Collection...")

	coll := db.DB.Collection("articles")

	var documents []interface{} = []interface{}{}

	logger.INFO("Deduping and Creating Bulk Insert Payload..")

	for i := 0; i < articleCount; i++ {
		hash := util.CreateHashSHA(articleData[i]["URL"])
		if !doesArticleExist(hash, coll) {
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
				"urlHash":         hash,
			})
			err = CacheClient.Set(hash, 1)
			if err != nil {
				log.Fatal("Cache crash")
			}
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

func parseFeedWorker(url string, wg *sync.WaitGroup) {
	data, err := util.ParseFeedURL(url)
	lock := sync.Mutex{}
	if err != nil {
		wg.Done()
		return
	}

	var feedArticles []map[string]string = []map[string]string{}

	for j := range data.Items {
		payload := map[string]string{
			"feedTitle":       data.Title,
			"feedDescription": data.Description,
			"feedURL":         url,
			"title":           util.StripHTMLTags(data.Items[j].Title),
			"content":         util.StripHTMLTags(data.Items[j].Content),
			"description":     util.StripHTMLTags(data.Items[j].Description),
			"updated":         strconv.Itoa(int(time.Now().Unix() * 1000)),
			"created":         strconv.Itoa(int(time.Now().Unix() * 1000)),
			"URL":             util.StripHTMLTags(data.Items[j].Link),
		}
		feedArticles = append(feedArticles, payload)
	}
	lock.Lock()
	defer lock.Unlock()
	articleData = append(articleData, feedArticles...)
	articleCount += len(feedArticles)
	wg.Done()
}

func doesArticleExist(hash string, coll *mongo.Collection) bool {
	val, err := CacheClient.Get(hash)
	if val == nil || err != nil {
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
