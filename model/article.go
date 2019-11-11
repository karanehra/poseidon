package model

//Article defines the schema of an article object
type Article struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	URL         string `json:"url"`
	FeedID      string `json:"feedID"`
}

//Validate ensures that the article is in a correct format
func (article *Article) Validate() []string {
	var errorData []string = []string{}
	if article.Title == "" {
		errorData = append(errorData, "article.title: field is required")
	}
	if article.Content == "" {
		errorData = append(errorData, "article.content: field is required")
	}
	if article.Description == "" {
		errorData = append(errorData, "article.description: field is required")
	}
	if article.URL == "" {
		errorData = append(errorData, "article.url: field is required")
	}
	return errorData
}

//Articles is a collection of articles
type Articles []Article
