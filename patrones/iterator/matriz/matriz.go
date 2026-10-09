package main

// Matriz contiene los datos a recorrer en filas y columnas.
type Matriz struct {
	data [][]int
}

// El constructor de matriz crea una nueva matriz y la devuelve.
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

// La fábrica Iterador devuelve un iterador nuevo posicionado al inicio.
func (m *Matriz) Iterador() *Iterador {
	return &Iterador{matriz: m, primera: true}
}

// Siguiente avanza al siguiente elemento y devuelve true si hay más.
// La primera llamada posiciona el iterador en el primer elemento.
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
// Solo debe llamarse después de que Siguiente() haya devuelto true.
func (it *Iterador) Valor() int {
	return it.matriz.data[it.filaActual][it.columnaActual]
}
