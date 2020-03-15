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
func ParseFeedURL(url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	client := http.Client{Timeout: 30 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36")
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	feed, err := fp.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return feed, err
}

//ParseCSVForURLs reads a local csv for url sources
func ParseCSVForURLs(fileName string) ([]string, error) {
	urlSet := []string{}
	pwd, _ := os.Getwd()
	url := filepath.Join(pwd, fileName)
	csvData, err := ioutil.ReadFile(url)
	if err != nil {
		fmt.Println("Error reading file")
		return nil, err
	}
	dataString := string(csvData)
	reader := csv.NewReader(strings.NewReader(dataString))
	for {
		url, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error occured reading csv")
			return nil, err
		}
		urlSet = append(urlSet, url[0])
	}
	return urlSet, nil
}
