package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// para experimentar: cambiá los ítems o la capacidad (probá capacidad 0 o
	// sin ítems: el valor máximo es 0)
	items := []Item{
		{Peso: 3, Valor: 2},
		{Peso: 4, Valor: 3},
		{Peso: 1, Valor: 4},
		{Peso: 2, Valor: 2},
		{Peso: 5, Valor: 5},
	}
	capacidad := 5
	const delay = 250 * time.Millisecond

	valor, seleccion := MochilaTabConItems(items, capacidad)

	fmt.Printf("Mochila 0/1: %d ítems, capacidad %d\n", len(items), capacidad)
	fmt.Printf("MochilaTab: %d · MochilaMemo: %d\n",
		MochilaTab(items, capacidad), MochilaMemo(items, capacidad))
	fmt.Printf("Items elegidos: %v (valor %d)\n", seleccion, valor)
	fmt.Println()
	dibujarTablaMochila(os.Stdout, items, capacidad, delay)
}
