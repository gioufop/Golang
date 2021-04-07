package main

import "fmt"

// Não tem operador ternário
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado" // não precisa do else nesse caso!

	/* No caso de um operador ternário ser Go suportasse, seria algo assim:
	return nota >= 6 ? "Aprovado" : "Reprovado"
	*/
}

func main() {
	fmt.Println(obterResultado(6.2))
}
