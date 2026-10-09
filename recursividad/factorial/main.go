package main

import "fmt"

func main() {
	for n := 0; n <= 20; n++ {
		fmt.Printf("%d! = %d\n", n, Factorial(n))
	}
}
