package main

import "fmt"

func main() {
	// para experimentar: subí el monto (la tabulación es O(m·n)) o cambiá las
	// denominaciones. Con {1, 5, 11}, Cambio(97) = 11: el ávido (todo de 11)
	// daría 13, la DP encuentra el óptimo.
	coins := []int{1, 5, 11}
	for _, monto := range []int{15, 22, 97} {
		fmt.Printf("Cambio(%d, %v) = %d monedas\n", monto, coins, Cambio(monto, coins))
	}
}
