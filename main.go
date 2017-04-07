package main

import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		fmt.Printf("Couldn't listen for incoming connection : %s", err)
	}

	fmt.Println("Listening for incoming connection")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Couldn't accept incoming connection: %s", err)
		}
		fmt.Println("Incoming connection accepted")
		go HandleConnection(conn)
	}

}

func HandleConnection(conn net.Conn) {
	request := bufio.NewReader(conn)
	requestStatusLine, err := request.ReadString('\n')
	if err != nil {
		fmt.Printf("Couldn't read requestStatusLine: %s", err)
		return
	}
	fmt.Printf("requestStatusLine: %s", requestStatusLine)

	fmt.Fprint(conn, expectedStatusLine(requestStatusLine))
	conn.Close()
}

func expectedStatusLine(request string) string {
	if request != "GET / HTTP/1.1\r\n" {
		return "HTTP/1.1 404 NOT FOUND\r\n\r\nSorry!\r\n"
	}
	return "HTTP/1.1 200 OK\r\n\r\nHi there\r\n"
}