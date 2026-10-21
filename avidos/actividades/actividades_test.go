package main

import "testing"

func TestSelectorActividadesIterativo(t *testing.T) {
	actividades := []Actividad{
		{"a1", 1, 4},
		{"a2", 3, 5},
		{"a3", 0, 6},
		{"a4", 5, 7},
		{"a5", 3, 9},
		{"a6", 5, 9},
		{"a7", 6, 10},
		{"a8", 8, 11},
		{"a9", 8, 12},
		{"a10", 2, 14},
		{"a11", 12, 15},
	}

	seleccionadas := SelectorActividadesIterativo(actividades)

	esperado := []Actividad{
		{"a1", 1, 4},
		{"a4", 5, 7},
		{"a8", 8, 11},
		{"a11", 12, 15},
	}
	if len(seleccionadas) != 4 {
		t.Fatalf("se esperaban 4 actividades seleccionadas; se obtuvieron %d: %v", len(seleccionadas), seleccionadas)
	}
	// Se compara campo por campo con un helper propio en vez de
	// reflect.DeepEqual: así el test no depende de reflexión y queda claro
	// qué se está comparando.
	if !mismasActividades(seleccionadas, esperado) {
		t.Errorf("SelectorActividadesIterativo(...) = %v; se esperaba %v", seleccionadas, esperado)
	}
}

// mismasActividades indica si dos listas de actividades tienen el mismo
// contenido, en el mismo orden.
func mismasActividades(a, b []Actividad) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
