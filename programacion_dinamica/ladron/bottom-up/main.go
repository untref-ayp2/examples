package main

import "fmt"

func main() {
	// para experimentar: cambiá los valores de las casas o agregá más (la
	// tabulación es O(n) y asume al menos dos casas)
	calle := []int{2, 10, 3, 6, 8, 1, 7}
	fmt.Printf("Rob(%v) = %d\n", calle, Rob(calle))
}
