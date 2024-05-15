//A quantidade de portas contidas no TCP é de 65536 portas que podem usadas para diversas funcionalidades e conexões.
package main

import (
  "fmt"

  "github.com/isaquefranklin/tcp-fun/port"
)

func main(){
  fmt.Println("Port Scanning")

  results := port.InitialScan("localhost")
  fmt.Println(results)
}
