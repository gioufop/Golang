package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... unit8 unit16 unit32 unit64
	var b byte = 255 //bite é um alias para unit8, nem todos tem alias
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2) //retorna o valor da letra a na tabela unicode que é igual a 97

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // por padrão vai ser do tipo float64

	var y = float32(49.99) // tbm funciona
	fmt.Println("O tipo de y é", reflect.TypeOf(y))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // por padrão vai ser do tipo float64

	// boolean
	bo := true
	fmt.Println("O valor de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é Giovanni"
	fmt.Println(s1 + "!")
	fmt.Println(len(s1))

	// string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Giovanni`
	fmt.Println(s2)
	fmt.Println(len(s2))
}
