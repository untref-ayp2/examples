package main

import (
	"fmt"
	"time"
)

func main() {
	// para experimentar: subí n (la versión optimizada es O(n) de tiempo y
	// O(1) de espacio)
	for _, n := range []int{10, 20, 30, 45} {
		start := time.Now()
		resultado := Fibonacci(n)
		fmt.Printf("Fibonacci(%d) = %d (%s)\n", n, resultado, time.Since(start))
	}
}
