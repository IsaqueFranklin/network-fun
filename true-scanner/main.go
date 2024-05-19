//implementing a simple port scan function in Go using the net package.

package main

import (
  "fmt"
  "github.com/isaquefranklin/network-fun/port"
)

func main() {
  fmt.Println("Port Scanning")
  results := port.InitialScan("localhost") 
  //The ports within the 1-1024 range are typically locked down by default.
  fmt.Println(results)
}
