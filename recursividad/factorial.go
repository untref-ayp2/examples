package recursividad

func Factorial(n int) int { //n debe ser mayor o igual a 0
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
