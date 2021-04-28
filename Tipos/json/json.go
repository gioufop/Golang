package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct para json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção, Eletrônico"}}
	p1Json, _ := json.Marshal(p1) //Marshal retorna 2 valores, o Json e o erro, estamos ignorando o erro
	fmt.Println(string(p1Json))

	//json para struct
	var p2 produto
	jsonString := `{"id":2, "nome":"caneta","preco":8.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2) // &p2: passando a referencia para o produto p2
	fmt.Println(p2.Tags[1])
	fmt.Println(p2)

}
