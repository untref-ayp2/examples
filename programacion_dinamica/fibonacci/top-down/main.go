package main

import (
	"fmt"
	"time"
)

func main() {
	// para experimentar: subí n (con memoización es O(n))
	for _, n := range []int{10, 20, 30, 45} {
		start := time.Now()
		resultado := Fibonacci(n)
		fmt.Printf("Fibonacci(%d) = %d (%s)\n", n, resultado, time.Since(start))
	}
}
