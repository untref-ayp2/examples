package main

import (
	"math"
)

func Rob(houses []int) int {
	n := len(houses)
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return max(houses[0], houses[n-1])
	}

	prevPrev := houses[0]
	prev := houses[1]
	cur := 0

	for x := 2; x < n; x++ {
		cur = max(houses[x]+prevPrev, prev)

		prevPrev = prev
		prev = cur
	}

	return cur
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
