package main

import (
  "testing"
  "net"
  "fmt"
  "bufio"
  "strings"
)

func assertEquals(expected string, actual string, t *testing.T) {
  if strings.TrimSpace(actual) != strings.TrimSpace(expected) {
    t.Errorf("We expected status line : %s but we got: %s", expected, actual)
  }
}

func TestServer(t *testing.T) {
  go main()

  //Test Launch Server
  conn, err := net.Dial("tcp", "localhost:8080")

  if err != nil {
    t.Errorf("Couldn't launch server: %s", err)
    return
  }

  fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
  statusLine, err := bufio.NewReader(conn).ReadString('\n')
  assertEquals("HTTP/1.0 200 OK", statusLine, t);
}
