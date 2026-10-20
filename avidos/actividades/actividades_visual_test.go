package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestDibujarLineaTiempo(t *testing.T) {
	actividades := []Actividad{
		{"a1", 1, 4},
		{"a4", 5, 7},
		{"a8", 8, 11},
		{"a11", 12, 15},
	}

	var buf bytes.Buffer
	dibujarLineaTiempo(&buf, actividades, 0)

	salida := buf.String()
	if !strings.Contains(salida, "a1") || !strings.Contains(salida, "a11") {
		t.Errorf("el dibujo debe incluir los nombres de las actividades; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "[====]") {
		t.Errorf("la barra de a1 (de 1 a 4) debe ocupar 3 unidades; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "[==]") {
		t.Errorf("la barra de a4 (de 5 a 7) debe ocupar 2 unidades; salida:\n%s", salida)
	}
	if !strings.Contains(salida, "14") {
		t.Errorf("el eje debe llegar hasta la última unidad ocupada (14); salida:\n%s", salida)
	}
}

func TestDibujarLineaTiempoActividadDeDuracionUno(t *testing.T) {
	// una actividad que ocupa una sola unidad de tiempo no debe romper la
	// alineación de la fila: su barra también tiene que medir 2 caracteres
	actividades := []Actividad{
		{"corta", 2, 3},
		{"larga", 0, 5},
	}

	var buf bytes.Buffer
	dibujarLineaTiempo(&buf, actividades, 0)

	lineas := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	// la primera línea es el eje de tiempos; las dos siguientes son las filas
	if len(lineas) != 3 {
		t.Fatalf("se esperaban 3 líneas (eje + 2 filas); salida:\n%s", buf.String())
	}
	if len(lineas[1]) != len(lineas[2]) {
		t.Errorf("la fila de %q mide %d caracteres y la de %q mide %d; deben medir lo mismo:\n%s",
			"corta", len(lineas[1]), "larga", len(lineas[2]), buf.String())
	}
	if !strings.Contains(lineas[1], "[]") {
		t.Errorf("la barra de una actividad de duración 1 debe dibujarse como []; salida:\n%s", buf.String())
	}
}

func TestDibujarLineaTiempoVacia(t *testing.T) {
	var buf bytes.Buffer
	dibujarLineaTiempo(&buf, nil, 0)

	if !strings.Contains(buf.String(), "No hay actividades") {
		t.Errorf("con lista vacía debe avisarlo; salida:\n%s", buf.String())
	}
}
