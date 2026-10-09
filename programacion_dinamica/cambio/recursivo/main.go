package main

import "fmt"

func main() {
	// para experimentar: subí el monto (la variante recursiva es exponencial)
	// o cambiá las denominaciones. Ojo con una moneda de 0: la recursión
	// nunca termina.
	coins := []int{1, 5, 11}
	for _, monto := range []int{15, 22} {
		fmt.Printf("Cambio(%d, %v) = %d monedas\n", monto, coins, Cambio(monto, coins))
	}
}
