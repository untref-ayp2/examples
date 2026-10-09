package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestDibujarRecorrido(t *testing.T) {
	distancia := 950
	estaciones := []int{200, 375, 550, 750}
	paradas := []int{375, 750}

	var buf bytes.Buffer
	dibujarRecorrido(&buf, distancia, estaciones, paradas, 0)

	salida := buf.String()
	if !strings.HasPrefix(salida, "0 ") {
		t.Errorf("el camino debe arrancar en la partida 0; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "D") {
		t.Errorf("el camino debe marcar el destino D; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "950") {
		t.Errorf("el destino debe rotular 950 km; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "*") {
		t.Errorf("las paradas de recarga deben marcarse con *; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "375") || !strings.Contains(salida, "750") {
		t.Errorf("las etiquetas deben rotular los kilómetros de las estaciones; salida:\n%s", salida)
	}
}
