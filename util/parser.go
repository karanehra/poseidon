package util

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mmcdole/gofeed"
)

//ParseFeedURL uses gofeed to fetch the rss feed contents
func ParseFeedURL(url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		fmt.Println("Error while parsing")
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
