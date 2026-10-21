package main

// Este archivo muestra el patrón Iterator: la matriz expone Iterador, que
// recorre fila por fila sin que el cliente sepa cómo está guardada. Si la
// matriz cambiara su representación interna, el cliente que usa
// Siguiente/Valor no se enteraría.

// Matriz contiene los datos a recorrer en filas y columnas.
type Matriz struct {
	data [][]int
}

// NewMatriz crea una nueva matriz que usa data como datos: no lo copia.
func NewMatriz(data [][]int) *Matriz {
	return &Matriz{data: data}
}

// Iterador permite recorrer la matriz sin exponer su estructura interna.
type Iterador struct {
	matriz        *Matriz
	filaActual    int
	columnaActual int
	primera       bool
}

// Iterador devuelve un iterador nuevo posicionado antes del primer elemento:
// la primera llamada a Siguiente lo deja en el (0, 0).
func (m *Matriz) Iterador() *Iterador {
	return &Iterador{matriz: m, primera: true}
}

// Siguiente avanza al siguiente elemento y devuelve true si hay más.
// La primera llamada posiciona el iterador en el primer elemento.
//
// Recorre por filas: al terminar una fila salta a la primera columna de la
// siguiente.
func (it *Iterador) Siguiente() bool {
	if !it.primera {
		it.columnaActual++
		if it.columnaActual >= len(it.matriz.data[it.filaActual]) {
			it.filaActual++
			it.columnaActual = 0
		}
	}
	it.primera = false
	return it.filaActual < len(it.matriz.data) &&
		it.columnaActual < len(it.matriz.data[it.filaActual])
}

// Valor devuelve el valor del elemento actual.
//
// Precondición: solo debe llamarse después de que Siguiente haya devuelto
// true; si no, la posición actual no es válida.
func (it *Iterador) Valor() int {
	return it.matriz.data[it.filaActual][it.columnaActual]
}
