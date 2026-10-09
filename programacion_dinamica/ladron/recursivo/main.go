package main

import "fmt"

func main() {
	// para experimentar: probá calles vacías, de una sola casa o más largas
	// (la variante recursiva es exponencial)
	calles := [][]int{
		{},
		{10},
		{2, 10, 3, 6, 8, 1, 7},
	}
	for _, calle := range calles {
		fmt.Printf("Rob(%v) = %d\n", calle, Rob(calle))
	}
}
