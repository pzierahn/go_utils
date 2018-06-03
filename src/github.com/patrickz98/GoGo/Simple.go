package GoGo

import (
	"encoding/json"
	"fmt"
)

func json2Str(v interface{}) (string, bool) {
	jsonString, err := json.Marshal(v)

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
