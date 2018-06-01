package main

import (
	"fmt"
	"net"
	"bufio"
	"log"
	"strconv"
	"os"
	"time"
	"io"
)

func handler(conn net.Conn) {

	var (
		buf = make([]byte, 1024)
		reader = bufio.NewReader(conn)
		writer = bufio.NewWriter(conn)
	)

	for {
		readBytes, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}

			break
		}

		msg := string(buf[:readBytes])
		fmt.Println("readBytes =", readBytes, "msg =", msg)
	}

	var delay time.Duration = 3
	time.Sleep(delay * time.Second)

	fmt.Println("Waiting for", delay, "s")

	msg := "Hallo"

	writer.Write([]byte(msg))
	writer.Flush()
	fmt.Println("Send: ", msg)

	conn.Close()
}

func SocketServer(port int) {

	listen, err := net.Listen("tcp4", ":" + strconv.Itoa(port))
	defer listen.Close()

	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", port, err)
		os.Exit(1)
	}

	log.Printf("Begin listen port: %d", port)

	for true {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go handler(conn)
	}
}

func main() {
	SocketServer(3333)
}
