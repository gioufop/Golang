package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) { //retorno nomeado
	primeiro = p1
	segundo = p2
	// return primeiro, segundo
	return //retorno limpo, não precisa passar segundo e primeiro
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
