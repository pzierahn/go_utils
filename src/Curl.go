package main

import (
	"net/http"
	"fmt"
)

const LOGTAG_CURL = "Curl"

func curl() {
	resp, err := http.Get("http://localhost")
	if err != nil {
		// handle err
		fmt.Println(LOGTAG_CURL, "-->", "curl:", err)
	}

	buffer := make([]byte, resp.ContentLength)
	resp.Body.Read(buffer)

	content := string(buffer)

	fmt.Println(LOGTAG_CURL, "-->", "content =", content)

	defer resp.Body.Close()
}
