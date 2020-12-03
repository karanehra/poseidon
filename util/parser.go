package util

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

//ParseFeedURL uses gofeed to fetch the rss feed contents
func ParseFeedURL(url string, ua string) (*gofeed.Feed, error) {
	if ua == "" {
		ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36"
	}
	fp := gofeed.NewParser()
	client := http.Client{Timeout: 30 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", ua)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	feed, err := fp.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return feed, err
}

//ParseCSVForUAs reads a local csv for User agents
func ParseCSVForUAs(fileName string) ([]string, error) {
	uaSet := []string{}
	pwd, _ := os.Getwd()
	url := filepath.Join(pwd, fileName)
	csvData, err := ioutil.ReadFile(url)
	if err != nil {
		fmt.Println("Error reading file")
		return nil, err
	}
	dataString := string(csvData)
	reader := csv.NewReader(strings.NewReader(dataString))
	reader.Comma = rune('|')
	for {
		url, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error occured reading csv")
			return nil, err
		}
		uaSet = append(uaSet, url[0])
	}
	return uaSet, nil
}
