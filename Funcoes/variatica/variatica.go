package main

import "fmt"

func media(numeros ...float64) float64 { // ... temos uma função variática que pode receber parametros variáveis
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros)) // retorna a média: "soma dos números" / "quantidade de números"
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 9.9))
}
