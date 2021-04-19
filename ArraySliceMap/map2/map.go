package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.00,
	}

	funcsESalarios["Rafael Filho"] = 1350.00 // ele vai adicionar
	delete(funcsESalarios, "Inexistente")    // não dá erro, mas não vai fazer nada, pq não tem essa chave

	for nome, salario := range funcsESalarios {
		fmt.Printf("%s: R$%.2f\n", nome, salario)
	}

}
