package main

import (
  "testing"
  "net"
  "fmt"
  "bufio"
)

func assertEquals(expected string, actual string, t *testing.T) {
  if actual != expected {
    t.Errorf("We expected: %s but we got: %s", expected, actual)
  }
}

func TestServer(t *testing.T) {
  go main()
  fmt.Printf("coucou 4")


  //Test Launch Server
  conn, err := net.Dial("tcp", "localhost:8080")

  if err != nil {
    t.Errorf("Couldn't launch server: %s", err)
    return
  }

  //Test that server response to a Get request
  fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

  fmt.Printf("coucou 7")
  response := bufio.NewReader(conn)

  fmt.Printf("coucou 8")

  // if len(response.buf) == 0 {
  //   t.Errorf("Couldn't read response")
  //   return
  // }

  statusLine, err := response.ReadString('\n')

  if statusLine == "" {
    t.Errorf("Couldn't read statusLine")
    return
  }

  fmt.Printf("coucou 9")
  assertEquals("HTTP/1.0 kjkhjk00 OK\n", statusLine, t)

  fmt.Printf("coucou 10")
  currentLine, err := response.ReadString('\n')

  headerBodySeparator := "\r\n"
  for currentLine != headerBodySeparator {
    currentLine, err = response.ReadString('\n')
  }
  fmt.Printf("coucou 11")

  body, err := response.ReadString('\n')
  assertEquals("how did we get here?", body, t)

}
