package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestDibujarArbolLlamadas(t *testing.T) {
	var buf bytes.Buffer
	dibujarArbolLlamadas(&buf, 5, 0)

	salida := buf.String()
	if !strings.Contains(salida, "fib(5) = 5") {
		t.Errorf("el árbol debe arrancar en fib(5); salida:\n%s", salida)
	}
	if !strings.Contains(salida, "fib(0) = 0") {
		t.Errorf("el árbol debe llegar hasta fib(0); salida:\n%s", salida)
	}
}

func TestDibujarTablaFibonacci(t *testing.T) {
	var buf bytes.Buffer
	dibujarTablaFibonacci(&buf, 8, 0)

	salida := buf.String()
	if !strings.Contains(salida, "dp[0] = 0") {
		t.Errorf("la tabla debe arrancar en dp[0]; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "dp[8] = 21") {
		t.Errorf("la tabla debe terminar en dp[8] = 21; salida:\n%s", salida)
	}
}
