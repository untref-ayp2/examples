package fibonacci_recursivo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	assert.Equal(t, 55, Fibonacci(10))
}
