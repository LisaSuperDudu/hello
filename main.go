package main

import (
	"fmt"
	"net"
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
	fmt.Fprint(conn, "HTTP/1.0 200 OK\r\n\r\n")
	fmt.Fprint(conn, "Hi there\r\n")
	conn.Close()
}
