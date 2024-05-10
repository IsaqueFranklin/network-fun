//A quantidade de portas contidas no TCP é de 65536 portas que podem usadas para diversas funcionalidades e conexões.
package main

import (
  "fmt"
  "net"
  "time"
)

func main(){
  limit := 444
  for i := 1; i < limit; i++ {
    address := fmt.Sprintf("scanme.nmap.org:%d", i)
    //Enviando comando com tempo de espera limitado.
    conn, err := net.DialTimeout("tcp", address, 1*time.Second)

    if err != nil {
      //passe para a próxima porta.
      continue
    }

    //fechando a conexão aberta.
    conn.Close()

    //Informando qual porta respondeu o comando.
    fmt.Printf("Porta %d aberta\n", i)
  }
}
