package main

import "runtime/debug"

func f3() {
	debug.PrintStack() // Imprime a pilha de execução de um programa assim que essa função é chamada
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()
}
