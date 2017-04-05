package main

import (
  "net/http"
  "fmt"
  "net"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "how did we get here?")
}

func main() {
  // http.HandleFunc("/", handler)
  // http.ListenAndServe(":8080", nil)
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
    conn.Close()
  }


}
