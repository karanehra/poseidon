package jobs

import (
	"encoding/csv"
	"fmt"
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
