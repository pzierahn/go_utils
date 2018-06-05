package main

import (
	"encoding/json"
	"fmt"
	"crypto/rand"
)

func json2Str(v interface{}) (string, bool) {
	jsonString, err := jsonfiy(v)

	if err != nil {
		return string(jsonString), true
	} else {
		fmt.Println(err)
		return "", false
	}
}

func jsonfiy(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "    ")
}

func containsInt(array []int, search int) bool {
	for _, value := range array {
		if value == search {
			return true
		}
	}

	return false
}

func containsString(array []string, search string) bool {
	for _, value := range array {
		if value == search {
			return true
		}
	}

	return false
}

func uuid() (uuid string) {
	buffer := make([]byte, 16)
	rand.Read(buffer)

	return fmt.Sprintf("%X-%X-%X-%X-%X", buffer[0:4], buffer[4:6], buffer[6:8], buffer[8:10], buffer[10:])
}