package main

import "fmt"

func main() {
	s := make([]int, 10) // slice de inteiros, com 10 elementos
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)        //reatribuindo a variável, sendo um slice de 10 posições e um array interno de 20
	fmt.Println(s, len(s), cap(s)) // len pega o tamanho e cap pega a capacidade máx dessa slice

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // já era a capacidade máx e add mais um elemento, n vai dar erro.. o slice passa a ter tamanho 21 com a cap max de 40
	fmt.Println(s, len(s), cap(s))
}
