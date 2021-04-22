package main

import "fmt"

func imprimirAprovados(aprovados ...string) { // quando recebo parametros variaveis, dentro da função eu os trato como array
	fmt.Println("Lista de aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	imprimirAprovados(aprovados...) //assim que passo um slice para uma funcão variatica

}
