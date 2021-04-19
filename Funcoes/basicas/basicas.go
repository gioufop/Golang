package main

import "fmt"

func f1() { //Não recebe nada como entrada de dados e não retorna nada
	fmt.Println("F1")
}

func f2(p1 string, p2 string) { // recebe parametros: p1 e p2 do tipo de paramentro: string  e não retorna nada
	fmt.Printf("F2: %s %s \n", p1, p2)
}

func f3() string { //Não recebe nenhum parametro e retorna uma string
	return "F3"
}

func f4(p1, p2 string) string { // recebe 2 parametros e retorna um valor string
	return fmt.Sprintf("F4: %s %s", p1, p2) //Sprintf não imprime no console e sim retorna uma string formatada
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Param1", "Param2")

	r3, r4 := f3(), f4("Param1", "Param2")

	fmt.Println(r3)
	fmt.Println(r4)
	fmt.Println(r3, r4)

	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)
}
