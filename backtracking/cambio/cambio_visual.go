package main

import (
	"fmt"
	"io"
	"time"
)

// Este archivo es la capa de visualización del ejemplo. Muestra, paso a paso,
// las combinaciones que explora el backtracking y, al final, la que usa la
// menor cantidad de piezas. No contiene lógica de backtracking: sólo consume
// los eventos que expone CambioDeMonedaConPaso.

// limpiarPantalla borra la terminal y lleva el cursor al inicio (ANSI).
const limpiarPantalla = "\033[2J\033[H"

// AnimarCambio resuelve el cambio mostrando las combinaciones exploradas: en
// cada paso dibuja la solución parcial, el resto por cambiar y la mejor
// combinación encontrada hasta ahora. Termina mostrando el cambio con la
// menor cantidad de piezas (o avisa si no existe combinación).
func AnimarCambio(cantidad int, denominaciones []int, w io.Writer, demora time.Duration) {
	pasos := 0

	var mejor []int

	actualizar := func(resto int, parcial []int, estado string) {
		pasos++
		fmt.Fprint(w, limpiarPantalla)
		fmt.Fprintf(w, "Cambio de %d con denominaciones %v\n\n", cantidad, denominaciones)
		fmt.Fprintf(w, "Combinación: %v\n", parcial)
		fmt.Fprintf(w, "Resto:       %d\n", resto)
		if estado != "" {
			fmt.Fprintf(w, "\n%s\n", estado)
		}
		if len(mejor) > 0 {
			fmt.Fprintf(w, "\nMejor hasta ahora: %v (%d piezas)\n", mejor, len(mejor))
		}
		fmt.Fprintf(w, "\nPasos: %d\n", pasos)
		time.Sleep(demora)
	}

	CambioDeMonedaConPaso(cantidad, denominaciones, ObservadorCambio{
		Intento: func(resto int, parcial []int) {
			actualizar(resto, parcial, "")
		},
		Solucion: func(solucion, mejorActual []int) {
			mejor = append([]int(nil), mejorActual...)
			actualizar(0, solucion, fmt.Sprintf("✓ Combinación exacta con %d piezas", len(solucion)))
		},
	})

	fmt.Fprint(w, limpiarPantalla)
	fmt.Fprintf(w, "Cambio de %d con denominaciones %v\n\n", cantidad, denominaciones)
	if len(mejor) > 0 {
		fmt.Fprintf(w, "El cambio con menos piezas: %v (%d piezas)\n", mejor, len(mejor))
	} else {
		fmt.Fprintf(w, "No hay combinación que dé %d con %v\n", cantidad, denominaciones)
	}
	fmt.Fprintf(w, "Pasos: %d\n", pasos)
}
