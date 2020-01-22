package model

//Feed defines the schema of a feed object
type Feed struct {
	URL   string   `json:"url"`
	Title string   `json:"title"`
	Tags  []string `json:"tags"`
}
