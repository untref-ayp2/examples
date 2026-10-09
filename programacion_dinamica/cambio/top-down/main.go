package main

import "fmt"

func main() {
	// para experimentar: subí el monto (con memoización es O(m·n)) o cambiá
	// las denominaciones. Ojo con una moneda de 0: la recursión nunca termina.
	coins := []int{1, 5, 11}
	for _, monto := range []int{15, 22, 97} {
		fmt.Printf("Cambio(%d, %v) = %d monedas\n", monto, coins, Cambio(monto, coins))
	}
}
