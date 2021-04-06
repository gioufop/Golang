package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("A soma é:", a+b)
	fmt.Println("A subtração é:", a-b)
	fmt.Println("A multiplicação é:", a*b)
	fmt.Println("A divisão é:", a/b)
	fmt.Println("O módulo é:", a%b)

	//bitwise
	fmt.Println("E =>", a&b)   //11 & 10 = 10  <=> faz Verdadeiro/Falso no AND
	fmt.Println("OU =>", a|b)  //11 | 10 = 11  <=> faz Verdadeiro/Falso no OR
	fmt.Println("Xor =>", a^b) //11 ^ 10 = 01

	c := 3.0
	d := 2.0

	// outras operações usando Math...
	fmt.Println("Maior valor:", math.Max(float64(a), float64(b)))
	fmt.Println("Menor valor:", math.Min(c, d))
	fmt.Println("Exponencial:", math.Pow(c, d))

}
