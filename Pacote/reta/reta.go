package main

import "math"

// Padrão utilizado no GO
// Iniciando com letra maiúscula é PÚBLICO (visivel dentro e fora do pacote)!
// Iniciando com letra minúscula é PRIVADO (visivel apenas dentro do pacote)!

// Exmeplo...
// Ponto - gerará algo público
// pontou ou _Ponto - gerará algo privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) { // Privado: catetos é visivel apenas dentro do pacote
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsável pro calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 { // Pública: Distancia é visivel fora do pacote
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2)) // calcula a raiz quadrada da soma dos quadrados dos catetos
}
