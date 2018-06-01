package main

import "encoding/json"

func JsonStr(v interface{}) (string, error) {
	jsonString, err := json.Marshal(v)

	if err != nil {
		return string(jsonString), err
	}

	return "", err
}