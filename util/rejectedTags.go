package util

//RejectedTagsMap maintains tags that whould be rejected upon detection
var RejectedTagsMap map[string]bool = map[string]bool{
	"a":         true,
	"an":        true,
	"the":       true,
	"in":        true,
	"out":       true,
	"today":     true,
	"there":     true,
	"of":        true,
	"after":     true,
	"yet":       true,
	"or":        true,
	"and":       true,
	"i":         true,
	"we":        true,
	"achieving": true,
	"dancing":   true,
	"achieve":   true,
	"dance":     true,
	"dawn":      true,
	"rare":      true,
	"but":       true,
}
