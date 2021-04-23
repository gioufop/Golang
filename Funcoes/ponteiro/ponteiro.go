package main

import "fmt"

func inc1(n int) {
	n++ // n = n + 1
}

// revisão: um ponteiro é representado por um *
func inc2(n *int) {
	// revisão: nesse caso o "*" é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++ // aqui ele vai retornar o valor do ponteiro
}

func main() {
	n := 1

	inc1(n) //por valor
	fmt.Println(n)

	// revisão: "&" é usado para obter o endereço da variável

	inc2(&n) // por referencia, nesse caso ela vai ser incrementada
	fmt.Println(n)
}
