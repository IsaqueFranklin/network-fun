//implementing a simple port scan function in Go using the net package.

package main

import (
  "fmt"
  "net"
  "strconv"
  "time"
)

func main() {
  fmt.Println("Port Scanning")
  open := scanPort("tcp", "localhost", 1313)
  //The ports within the 1-1024 range are typically locked down by default.
  fmt.Printf("Port Open: %t\n", open)
}
