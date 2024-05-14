package fibbonacci_top_down

func Fibonacci(n int) int {
	return fibonacci(n, make(map[int]int))
}

func fibonacci(n int, mem map[int]int) int {
	key := n
	if result, ok := mem[key]; ok {
		return result
	}

	if n < 2 {
		mem[key] = n
		return mem[key]
	}
	mem[key] = fibonacci(n-1, mem) + fibonacci(n-2, mem)
	return mem[key]
}
