package main

import "fmt"

func main() {
	i := 1

	var p *int = nil // cria um ponteiro p, do tipo int que não está apontando para lugar nenhum no momento
	p = &i           // pegando o endereço da variável
	*p++             // *ponteiro vc "desreferencia", vc pega o valor associado a esse ponteiro de fato
	i++

	// Go não tem aritmética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i) // retorna: 0xc0000b8010 3 3 0xc00001a108
}
