package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95135745682] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados { // navega por chave e valor
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[95135745682]) // ler o valor
	delete(aprovados, 95135745682)      //deletar um valor do map: qual map e qual a chave
	fmt.Println(aprovados[95135745682]) // resultado é vazio

	alunos := make(map[string]string)

	alunos["Ana"] = "Reprovada"
	alunos["Pedro"] = "Aprovado"
	alunos["Maria"] = "Aprovada"

	for nome, resultado := range alunos {
		fmt.Printf("Resultado final => %s: %s\n", nome, resultado)
	}

	fmt.Println("A aluna Ana foi:", alunos["Ana"])

}
