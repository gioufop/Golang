package main

import "fmt"

func main() {
	// Array: estrutura homogênea (mesmo tipo) e estática (fixo/não cresce)
	var notas [3]float64 // definindo o tamanho (array de 3 posições) e o tipo do array (float64)
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1

	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Media %.2f\n", media)
}
