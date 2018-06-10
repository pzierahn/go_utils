package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(writer http.ResponseWriter, request *http.Request) {
	connection, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer connection.Close()

	for {
		mt, message, err := connection.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("message=%s", message)

		err = connection.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func home(writer http.ResponseWriter, request *http.Request) {
	// homeTemplate.Execute(writer, "ws://" + request.Host + "/echo")
	homeTemplate.Execute(writer, nil)
}

func startWebsocket() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var html = `
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
` + readFileStr("/Users/patrick/github/GoGo/main.js") + `
</script>
</head>
<body>
</body>
</html>
`
var homeTemplate = template.Must(template.New("").Parse(html))
