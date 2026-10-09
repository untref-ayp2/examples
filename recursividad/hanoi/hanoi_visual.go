package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// limpiarPantalla borra la terminal y lleva el cursor al inicio (ANSI).
// Se usa para dibujar cada estado en la misma posición, como el ejemplo
// de sudoku.
const limpiarPantalla = "\033[2J\033[H"

// estadoHanoi representa el estado de las tres varas. Cada vara guarda sus
// discos del más chico (arriba) al más grande (abajo). Además lleva el
// contador de movimientos y el último disco movido, para poder resaltarlo
// en el dibujo.
type estadoHanoi struct {
	discos      int
	varas       [3][]int
	movimientos int // contador de movimientos realizados
	ultimoDisco int // tamaño del último disco movido (0 = ninguno)
	ultimaVara  int // vara donde quedó el último disco movido
}

// nuevoEstadoHanoi crea un estado inicial con los n discos apilados en la
// primera vara.
func nuevoEstadoHanoi(n int) *estadoHanoi {
	estado := &estadoHanoi{discos: n}
	for disco := n; disco >= 1; disco-- {
		estado.varas[0] = append(estado.varas[0], disco)
	}
	return estado
}

// mover pasa el disco superior de la vara origen a la vara destino, avanza
// el contador y recuerda el disco movido para resaltarlo en el dibujo.
func (e *estadoHanoi) mover(disco, origen, destino int) {
	tope := e.varas[origen][len(e.varas[origen])-1]
	if tope != disco {
		panic("se intentó mover un disco que no está en el tope")
	}
	e.varas[origen] = e.varas[origen][:len(e.varas[origen])-1]
	e.varas[destino] = append(e.varas[destino], disco)
	e.movimientos++
	e.ultimoDisco = disco
	e.ultimaVara = destino
}

func (e *estadoHanoi) alturaMaxima() int {
	return e.discos
}

// centrar centra contenido en una celda de ancho dado, contando caracteres
// (no bytes, porque los bordes y los discos usan UTF-8 multibyte).
func centrar(contenido string, ancho int) string {
	runas := len([]rune(contenido))
	izquierda := strings.Repeat(" ", (ancho-runas)/2)
	return izquierda + contenido + strings.Repeat(" ", ancho-runas-len([]rune(izquierda)))
}

// celdaDelDisco devuelve la celda de una vara con el disco dado, resaltado
// en cian cuando acabó de moverse.
func celdaDelDisco(disco, ancho int, resaltado bool) string {
	anchoDisco := 2*disco + 1
	barra := strings.Repeat("█", anchoDisco)
	if resaltado {
		barra = "\033[36m" + barra + "\033[0m"
	}
	izquierda := strings.Repeat(" ", (ancho-anchoDisco)/2)
	return izquierda + barra + strings.Repeat(" ", ancho-anchoDisco-len([]rune(izquierda)))
}

// Tablero devuelve el estado actual de las tres varas en un marco fijo, con
// el disco que acaba de moverse resaltado y el contador de movimientos.
func (e *estadoHanoi) Tablero() string {
	ancho := 2*e.discos + 1
	var b strings.Builder

	b.WriteString("┏" + strings.Repeat("━", ancho) + "┳" + strings.Repeat("━", ancho) + "┳" + strings.Repeat("━", ancho) + "┓\n")

	for nivel := e.discos; nivel >= 1; nivel-- {
		b.WriteString("┃")
		for i, vara := range e.varas {
			if len(vara) < nivel {
				b.WriteString(centrar("│", ancho))
			} else {
				disco := vara[len(vara)-nivel]
				resaltado := disco == e.ultimoDisco && i == e.ultimaVara
				b.WriteString(celdaDelDisco(disco, ancho, resaltado))
			}
			b.WriteString("┃")
		}
		b.WriteString("\n")
	}

	b.WriteString("┗" + strings.Repeat("━", ancho) + "┻" + strings.Repeat("━", ancho) + "┻" + strings.Repeat("━", ancho) + "┛\n")

	b.WriteString(" ")
	for _, nombre := range []string{"A", "B", "C"} {
		b.WriteString(centrar(nombre, ancho+1))
	}
	b.WriteString("\n")

	b.WriteString(fmt.Sprintf("\nMovimientos: %d\n", e.movimientos))

	return b.String()
}

// dibujar escribe el estado actual en w y espera delay antes de devolver.
// Borra la pantalla primero, de modo que cada estado se dibuja en el mismo
// lugar (posición fija). No imprime por su cuenta: todo el dibujo sale por w.
func (e *estadoHanoi) dibujar(w io.Writer, delay time.Duration) {
	fmt.Fprint(w, limpiarPantalla)
	fmt.Fprint(w, e.Tablero())
	time.Sleep(delay)
}
