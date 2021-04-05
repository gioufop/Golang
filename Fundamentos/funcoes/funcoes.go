package main

import "fmt"

func somar(a int, b int) int { // Os dois parametros de entrada e o tipo de retorno
	return a + b // como tem o int no tipo de retorno eu preciso chamar o return
}

func imprimir(valor int) {
	fmt.Println(valor) // aqui não precisa retornada nd pq n tenho o "int" no tipo de retorno
}

/* Usando na função main.go
func main() {
	resultado := somar(3, 4) //declarando e inicializando a variável :=
	imprimir(resultado)
}
*/
