package util

import (
	"crypto/sha1"
	"encoding/hex"
	"regexp"
)

//StripHTMLTags removes all html tags from the give text
func StripHTMLTags(text string) string {
	regex := regexp.MustCompile("<[^>]*>")
	return regex.ReplaceAllLiteralString(text, "")
}

//CreateHashSHA hashes a given string input using the SHA-1 algorithm
// and returns a hex representation of it
func CreateHashSHA(value string) string {
	hasher := sha1.New()
	hasher.Write([]byte(value))
	res := hex.EncodeToString(hasher.Sum(nil))
	return res
}
