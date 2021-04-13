package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // o compilador que vai contar e gerar a quantidade total de elementos na inicialização do array

	for i, numero := range numeros { // asssim o for será percorrido de acordo com o range
		fmt.Printf("%d) %d\n", i, numero)
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, numero := range numeros { // ignora o indice
		fmt.Printf("%d ", numero)
	}

}
