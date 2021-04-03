package main

import "fmt"

func main() {
	fmt.Print("Mesma") //não quebra a linha
	fmt.Print(" linha!")

	fmt.Println(" Nova")
	//quebra a linha depois da execução
	fmt.Println("linha.")

	x := 3.141516

	// fmt.Println("O valor de x é " + x) não funciona em GO tem que fazer virar string:
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f \n", x) // %.2f = número de casas decimais arredondando!

	fmt.Printf("O valor de x é %f \n", x) // %f quer dizer que vai imprimir um fload

	/*
		%d inteiro
		%f float
		%t boolean
		%s string
		%v serve para vários tipos diferentes
	*/

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)

	fmt.Printf("\n %v %v %v %v", a, b, c, d)

}
