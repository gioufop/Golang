package main

import "fmt"

func soma(a, b int) int { // Cria uma variável soma que recebe o resultado de uma função anonima
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func main() {
	fmt.Println(soma(2, 3))
	fmt.Println(sub(3, 2))
}
