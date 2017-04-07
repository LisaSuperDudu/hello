package main

import (
	"bufio"
	"fmt"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	go main()

	//Test Launch Server
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		t.Errorf("Couldn't launch server: %s", err)
		return
	}

	//Test that server response to a Get request
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	response := bufio.NewReader(conn)

	statusLine, err := response.ReadString('\n')

	if err != nil {
		t.Errorf("Couldn't read statusLine: %s", err)
		return
	}


	expectedStatusLine := "HTTP/1.0 200 OK\r\n"
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
	//_, err := response.ReadString('\n')
	//assertEquals("how did we get here?", body, t)

}
