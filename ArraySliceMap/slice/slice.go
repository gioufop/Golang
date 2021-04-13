package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array! Slice define um pedaço de um array
	s2 := a2[1:3] // pega o elemento de índice 1 e NÃO inclui o elemento de índice 3 ou seja, tá pegando os elementos 2 e 3
	fmt.Println(a2, s2)

	s3 := a2[:2] // novo slice, mas apontando para o mesmo array que vai do inicio do array até a posição de indice 2, sem incluir o indice 2!
	fmt.Println(a2, s3)

	// Você pode imaginar que um slice: uma estrutura que tem um tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
