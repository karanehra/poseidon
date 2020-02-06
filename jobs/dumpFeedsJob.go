package jobs

import (
	"os"
	"poseidon/db"
	"poseidon/logger"
	"strings"

	"github.com/karanehra/schemas"
	"go.mongodb.org/mongo-driver/bson"
)

//DumpFeedsJob gets feeds out of DB and dumps them in a csv
func DumpFeedsJob() {
	logger := logger.Logger{}
	logger.INFO("Starting dump feeds job")
	feeds, err := schemas.GetFeeds(db.DB, bson.D{})
	if err != nil {
		logger.ERROR("Unable to fetch feed data")
	}
	csvFile, err := os.Create("feed-dump.csv")
	for i := range feeds {
		dataRow := []string{
			feeds[i].Title,
			feeds[i].URL,
			strings.Join(feeds[i].Tags, "|"),
		}
	}
	logger.SUCCESS("Job Done")
}
