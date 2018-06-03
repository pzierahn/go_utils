package main

import (
	"net/http"
	"fmt"
)

func sayHello(responseWriter http.ResponseWriter, request *http.Request) {
	// message := request.URL.Path
	// message = strings.TrimPrefix(message, "/")
	// message := "<h1>Hello</h1>"

	jsonTest := JsonStruc{
		Test1: "asdfasdf",
		Test2: 12334,
		Test3: []int{3, 6},
		Test4: []string{"asdfas", "asgasd", "asdgeqcv"},
	}

	fmt.Println(jsonTest)

	jsonBytes, _ := jsonfiy(jsonTest)
	responseWriter.Write([]byte(jsonBytes))
}

func StartWebServer() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}