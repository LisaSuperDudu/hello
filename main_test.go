package main

import (
  "testing"
  "net"
)

func TestSum(t *testing.T) {
  total := Sum(5, 5)
  if total != 10 {
    t.Errorf("Sum was incorrect, got %d, want %d", total, 10)
  }
}

func TestLaunchServer(t *testing.T) {
  go main()

  _, err := net.Dial("tcp", "localhost:8080")

  if err != nil {
    t.Errorf("Couldn't launch server: %s", err)
  }
}
