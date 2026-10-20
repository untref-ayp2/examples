package main

import (
	"math"
)

func Rob(houses []int) int {
	n := len(houses)

	// casos base: sin casas no se roba nada; con una o dos, se elige la que
	// rinde más, porque no se pueden robar dos casas vecinas
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return max(houses[0], houses[n-1])
	}

	table := make([]int, n)
	table[0] = houses[0]
	table[1] = max(houses[0], houses[1])

	for x := 2; x < n; x++ {
		table[x] = max(houses[x]+table[x-2], table[x-1])
	}

	return table[n-1]
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
