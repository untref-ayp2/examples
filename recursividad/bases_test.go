package recursividad

import (
	"testing"
)

// Estas pruebas, en realidad, no comprueban nada porque son salidas por pantalla.
// Deberían modificarse las funciones para retornar valores y poder comprobar salidas.

func TestImprimirEnDecimal(t *testing.T) {
	ImprimirEnDecimal(127)
}

func TestImprimirEnBase(t *testing.T) {
	ImprimirEnBase(42, 2)
}
