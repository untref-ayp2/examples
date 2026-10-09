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

func TestDibujarLineaTiempoVacia(t *testing.T) {
	var buf bytes.Buffer
	dibujarLineaTiempo(&buf, nil, 0)

	if !strings.Contains(buf.String(), "No hay actividades") {
		t.Errorf("con lista vacía debe avisarlo; salida:\n%s", buf.String())
	}
}
