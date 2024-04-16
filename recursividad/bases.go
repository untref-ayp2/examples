package recursividad

import "fmt"

func ImprimirEnDecimal(n int) {
	if n >= 10 {
		ImprimirEnDecimal(n / 10)
	}

	fmt.Println(n % 10)
}

const digits string = "0123456789abcdef"

func ImprimirEnBase(n int, base int) {
	if n >= base {
		ImprimirEnBase(n/base, base)
	}

	fmt.Println(string(digits[n%base]))
}
