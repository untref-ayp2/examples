package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestNuevoEstadoHanoi(t *testing.T) {
	estado := nuevoEstadoHanoi(3)

	esperado := []int{3, 2, 1}
	if len(estado.varas[0]) != len(esperado) {
		t.Fatalf("vara A tiene %d discos; se esperaban %d", len(estado.varas[0]), len(esperado))
	}
	for i, disco := range estado.varas[0] {
		if disco != esperado[i] {
			t.Errorf("vara A, posición %d = %d; se esperaba %d", i, disco, esperado[i])
		}
	}
	if len(estado.varas[1]) != 0 || len(estado.varas[2]) != 0 {
		t.Errorf("las varas B y C deben arrancar vacías")
	}
	if estado.movimientos != 0 {
		t.Errorf("el contador debe arrancar en 0; es %d", estado.movimientos)
	}
}

func TestMoverDisco(t *testing.T) {
	estado := nuevoEstadoHanoi(2)

	estado.mover(1, 0, 1)
	estado.mover(2, 0, 2)
	estado.mover(1, 1, 2)

	if len(estado.varas[2]) != 2 {
		t.Fatalf("vara C tiene %d discos; se esperaban 2", len(estado.varas[2]))
	}
	if estado.varas[2][0] != 2 || estado.varas[2][1] != 1 {
		t.Errorf("vara C = %v; se esperaba [2 1]", estado.varas[2])
	}
	if estado.movimientos != 3 {
		t.Errorf("el contador debe marcar 3 movimientos; marca %d", estado.movimientos)
	}
	if estado.ultimoDisco != 1 || estado.ultimaVara != 2 {
		t.Errorf("debe resaltarse el disco %d en la vara %d; got disco=%d vara=%d",
			1, 2, estado.ultimoDisco, estado.ultimaVara)
	}
}

func TestDibujar(t *testing.T) {
	estado := nuevoEstadoHanoi(1)

	var buf bytes.Buffer
	estado.dibujar(&buf, 0)

	salida := buf.String()
	for _, fragmento := range []string{"\033[2J\033[H", "█", "│", "┏", "┗", "A", "B", "C", "Movimientos: 0"} {
		if !strings.Contains(salida, fragmento) {
			t.Errorf("el dibujo debe contener %q; salida:\n%s", fragmento, salida)
		}
	}
}

// sinAnsi quita las secuencias de color ANSI (\033[...m), para poder
// comparar el dibujo sin depender del resaltado.
func sinAnsi(s string) string {
	var b strings.Builder
	enSecuencia := false
	for _, r := range s {
		if r == '\033' {
			enSecuencia = true
		}
		if !enSecuencia {
			b.WriteRune(r)
		}
		if enSecuencia && r == 'm' {
			enSecuencia = false
		}
	}
	return b.String()
}

func TestTableroResaltaUltimoMovimiento(t *testing.T) {
	estado := nuevoEstadoHanoi(2)

	if strings.Contains(estado.Tablero(), "\033[36m") {
		t.Errorf("sin movimientos no debe haber disco resaltado")
	}

	estado.mover(1, 0, 1)
	tablero := estado.Tablero()
	if !strings.Contains(tablero, "\033[36m") {
		t.Errorf("el disco recién movido debe resaltarse en cian; tablero:\n%s", tablero)
	}
	if !strings.Contains(tablero, "Movimientos: 1") {
		t.Errorf("el contador debe marcar 1 movimiento; tablero:\n%s", tablero)
	}
}

func TestTableroApilaDiscosDeAbajoParaArriba(t *testing.T) {
	estado := nuevoEstadoHanoi(3)
	// Secuencia clásica de 7 movimientos: al final todos los discos quedan en C.
	for _, paso := range [][3]int{{1, 0, 2}, {2, 0, 1}, {1, 2, 1}, {3, 0, 2}, {1, 1, 0}, {2, 1, 2}, {1, 0, 2}} {
		estado.mover(paso[0], paso[1], paso[2])
	}

	var filas []string
	for _, linea := range strings.Split(estado.Tablero(), "\n") {
		if strings.Contains(linea, "┃") {
			filas = append(filas, sinAnsi(linea))
		}
	}
	if len(filas) != 3 {
		t.Fatalf("se esperaban 3 filas de discos; hay %d:\n%s", len(filas), strings.Join(filas, "\n"))
	}

	// El disco más chico (1) es el de arriba y el más grande (3) el de abajo:
	// un disco chico jamás queda debajo de uno grande.
	if !strings.Contains(filas[0], "  ███  ") || strings.Contains(filas[0], "███████") {
		t.Errorf("la fila superior debe mostrar el disco 1 (chico) y no el 3 (grande):\n%s", filas[0])
	}
	if !strings.Contains(filas[2], "███████") || strings.Contains(filas[2], "  ███  ") {
		t.Errorf("la fila inferior debe mostrar el disco 3 (grande) y no el 1 (chico):\n%s", filas[2])
	}
}
