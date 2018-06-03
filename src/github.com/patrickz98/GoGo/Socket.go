package GoGo

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

const LOGTAG_SOCKET = "Socket"

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
				fmt.Println(LOGTAG_SOCKET, "-->", "read error =", err)
			}

			break
		}

		msg := string(buf[:readBytes])
		fmt.Println(LOGTAG_SOCKET, "-->", "readBytes =", readBytes, "msg =", msg)
	}

	var delay time.Duration = 3
	time.Sleep(delay * time.Second)

	fmt.Println(LOGTAG_SOCKET, "-->", "delay =", delay)

	msg := "Hallo"

	writer.Write([]byte(msg))
	writer.Flush()
	fmt.Println(LOGTAG_SOCKET, "-->", "msg =", msg)

	conn.Close()
}

func SocketServer(port int) {

	listen, err := net.Listen("tcp4", ":" + strconv.Itoa(port))
	defer listen.Close()

	if err != nil {
		fmt.Println(LOGTAG_SOCKET, "-->", "err != nil", "port =", port, "err =", err)
		os.Exit(1)
	}

	fmt.Println(LOGTAG_SOCKET, "-->", "port =", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go handler(conn)
	}
}
