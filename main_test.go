package main

import (
	"bufio"
	"fmt"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	go main()
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		t.Errorf("Couldn't launch server: %s", err)
		return
	}

	//Test that server responds to a Get request
	request := "GET / HTTP/1.1\r\n\r\n"
	fmt.Fprintf(conn, request)

	response := bufio.NewReader(conn)

	statusLine, err := response.ReadString('\n')

	if err != nil {
		t.Errorf("Couldn't read statusLine: %s", err)
		return
	}


	expectedStatusLine := "HTTP/1.1 200 OK\r\n"
	if statusLine != expectedStatusLine {
		t.Errorf("We expected: %s but we got: %s", expectedStatusLine, statusLine)
		return
	}

	currentLine, err := response.ReadString('\n')

	if err != nil {
		t.Errorf("Couldn't read currentLine: %s", err)
		return
	}

	expectedBlankLine := "\r\n"
	if currentLine != expectedBlankLine {
		t.Errorf("We expected %s but we got: %s", expectedBlankLine, currentLine )
		return
	}

	firstBodyLine, err := response.ReadString('\n')

	if err != nil {
		t.Errorf("Couldn't read firstBodyLine %s", err)
		return
	}

	expectedFirstBodyLine := "Hi there\r\n"
	if firstBodyLine != expectedFirstBodyLine {
		t.Errorf("We expected %s but we got: %s", expectedFirstBodyLine, firstBodyLine )
		return
	}

	return
}
