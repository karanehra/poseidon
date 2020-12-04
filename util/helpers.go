package util

import "go.mongodb.org/mongo-driver/bson/primitive"

//SliceContains checks if a string value exists in a slice
func SliceContains(s []string, v string) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}
	return false
}

//ObjectIDToHexString Converts an object ID primitive to a reusable hex string
func ObjectIDToHexString(objectID interface{}) string {
	return objectID.(primitive.ObjectID).Hex()
}
