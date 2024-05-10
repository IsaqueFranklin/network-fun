package main

import (
  "fmt"
  "net"
  "time"
)

//Aqui geramos as goroutines para o processo de validação das portas.
func generateGoroutines(host string, ports, resultScanner chan int) {
  //percorrendo todas as portas que chegaram no canal
  for p := range ports {
    //padronizando o endereço para o scanner de host + porta.
    address := fmt.Sprintf("%v:%d", host, p)
    //tentando conectar na porta com o protocolo TCP e aguardando um tempo.
    conn, err := net.DialTimeout("tcp", address, 1*time.Second)
    //caso tenha um erro
    if err != nil {
      //passe para a próxima porta e informe que a porta está fechada.
      resultScanner <- 0
      continue
    }

    //fechando a conexão aberta
    conn.Close()

    //retornando a porta que está aberta.
    resultScanner <- p
  }
}

func SenderPortsProcess(limitPort int, ports chan int) {
  //percorrendo a quantidade de portas para análise.
  for i := 1; i <= limitPort; i++ {
    //enviando cada porta da lista
    ports <- i
  }
}

//Recebendo todas as portas que estão abertas.
func getResultPortsOpen(limitPort int, results chan int) []int {
  var listOpenPorts []int

  //percorrendo todas as portas informadas.
  for i := 0; i < limitPort; i++ {
    //pegando todas as portas que foram validadas.
    port := <- results
    //verificando qual porta está aberta
    if port != 0 {
      //adicionando lista de portas abertas
      listOpenPorts = append(listOpenPorts, port)
    }
  }

  //retornando a lista de portas abertas.
  return listOpenPorts
}

func main() {
  //adicionando o host para verificação.
  var host string = "scanme.nmap.org"

  //a quantidade de portas a serem verificadas.
  var limitPort int = 1024
  //canal com quantidade de portas.
  ports := make(chan int, limitPort)
  //canal que recebe todos os resultados de portas abertas ou não.
  resultsScan := make(chan int)

  //percorrendo os canais e criando as goroutines para serem executadas.
  for i := 0; i < cap(ports); i++ {
    //criando goroutines
    go generateGoroutines(host, ports, resultsScan)
  }

  //função responsável para começar a validação das portas.
  go SenderPortsProcess(limitPort, ports)

  //recebendo todos os resultados das portas abertas ou não.
  portsOpens := getResultPortsOpen(limitPort, resultsScan)

  //fechando todos os canais de portas.
  close(ports)

  //fechando todos os canais de resuktado do scanner.
  close(resultsScan)

  //informando qual host foi validado.
  fmt.Printf("Endereço verificado - %v\n", host)

  //percorrendo todas as portas que estão abertas.
  for _, port := range portsOpens {
    //informando as portas abertas.
    fmt.Printf("%d aberta\n", port)
  }
}
