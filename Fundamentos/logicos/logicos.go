package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2    // E
	comprarTv32 := trab1 != trab2    // ou exclusivo
	comprarSorvete := trab1 || trab2 // OU

	return comprarTv50, comprarTv32, comprarSorvete // vou retornar os 3 booleans da função compras!
}

func main() {
	tv50, tv32, sorvete := compras(true, true)                                                 // os true de compras são os 2 operandos da função, ou seja as duas entradas que vão gerar o resultado!
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete) // %t: para imprimir o boolean
}
