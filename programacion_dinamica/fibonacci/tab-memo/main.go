package main

import (
	"fmt"
	"time"
)

func main() {
	// para experimentar: subí n (ambas variantes son O(n))
	for _, n := range []int{10, 20, 30, 45} {
		iniTab := time.Now()
		tab := FibonacciTab(n)
		tiempoTab := time.Since(iniTab)

		iniMemo := time.Now()
		memo := FibonacciMemo(n)
		tiempoMemo := time.Since(iniMemo)

		fmt.Printf("FibonacciTab(%d) = %d (%s) · FibonacciMemo(%d) = %d (%s)\n",
			n, tab, tiempoTab, n, memo, tiempoMemo)
	}
}
