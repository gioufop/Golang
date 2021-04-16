package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int
	fmt.Println(array1)

	// array1 = append(array1, 4, 5, 6)  <= não é possível
	slice1 = append(slice1, 4, 5, 6) // append add elementos dentro de um slice
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) // o copy não faz com o que o slice cresça! vai copiar os 2 primeiros elementos do slice1 para o slice 2 que só tem 2 posições
	fmt.Println(slice2)
}
