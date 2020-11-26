package util

//SliceContains checks if a string value exists in a slice
func SliceContains(s []string, v string) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}
	return false
}
