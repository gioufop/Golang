package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y)) // deixa a linguagem mais verbosa, mas é como ela funciona. É preciso converter o inteiro para float no caso.

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal) //desconsidera as casas decimais [Professor ruim em?!]

	// cuidado...
	fmt.Println("Teste" + string(123)) // o 123 na verdade é o código desse caracter na tabela unicode que é "{"

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123") // Essa func retorna 2 valores um inteiro e um erro, o _ serve para ignorarmos o erro por enquanto
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") //ou o valor 1 que tbm é considerado verdadeiro
	if b {
		fmt.Println("Verdadeiro")
	}
}
