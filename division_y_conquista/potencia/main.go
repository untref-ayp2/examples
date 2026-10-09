package main

import "fmt"

func main() {
	// para experimentar: cambiá la base n y el exponente x
	casos := [][2]int{
		{2, 2},  // exponente par
		{2, 3},  // exponente impar
		{3, 0},  // exponente cero
		{2, 16}, // potencias grandes
		{16, 6}, // otra base
	}
	for _, c := range casos {
		n, x := c[0], c[1]
		fmt.Printf("%d^%d = %d\n", n, x, Potencia(n, x))
	}
}
