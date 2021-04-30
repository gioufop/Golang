package main

import "fmt"

func main() {
	ch := make(chan int, 1) // dentro desse canal de comunicação serão enviados valores inteiros

	ch <- 1 // enviando dados para o canal (escrita) no caso o valor 1
	<-ch    // recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)

}
