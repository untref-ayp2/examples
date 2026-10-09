package main

import "fmt"

func main() {
	lista := &Lista{}
	lista.AgregarAlFinal(1)
	lista.AgregarAlFinal(2)
	lista.AgregarAlFinal(3)

	it := lista.Iterador()
	for it.Siguiente() {
		fmt.Println(it.Valor())
	}
	// Salida:
	// 1
	// 2
	// 3
}
