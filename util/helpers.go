package util

import "regexp"

//StripHTMLTags removes all html tags from the give text
func StripHTMLTags(text string) string {
	regex := regexp.MustCompile("<[^>]*>")
	return regex.ReplaceAllLiteralString(text, "")
}
