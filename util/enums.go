package util

import "fmt"

var JobTypes map[string]interface{}

func init() {
	JobTypes = make(map[string]interface{})
	JobTypes["ADD_FEED"] = addFeed
	JobTypes["UPDATE_FEEDS"] = updateFeed
}

func addFeed() {
	fmt.Println("Called add feed")
}

func updateFeed() {
	fmt.Println("Called update feed")
}
