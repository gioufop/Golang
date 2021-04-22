package main

import "fmt"

func fatorial(n uint) uint { //recebe um inteiro sem sinal, retorna um inteiro sem sinal
	switch {
	case n == 0: //condição de parada
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
}
