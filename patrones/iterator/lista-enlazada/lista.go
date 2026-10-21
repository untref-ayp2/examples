package main

// Este archivo muestra el patrón Iterator sobre una lista enlazada: el
// cliente recorre los nodos con Siguiente/Valor sin manejar punteros. La
// lista mantiene su estructura de nodos oculta.

// Nodo es un eslabón de la lista enlazada: guarda un valor y un puntero al
// nodo siguiente (nil en el último).
type Nodo struct {
	valor int
	sig   *Nodo
}

// Lista es la colección.
type Lista struct {
	cabeza *Nodo
}

// AgregarAlFinal crea un nodo con el valor dado y lo engancha al final de la
// lista. Si la lista está vacía, pasa a ser la cabeza.
func (l *Lista) AgregarAlFinal(valor int) {
	nuevo := &Nodo{valor: valor}
	if l.cabeza == nil {
		l.cabeza = nuevo
		return
	}
	actual := l.cabeza
	for actual.sig != nil {
		actual = actual.sig
	}
	actual.sig = nuevo
}

// Iterador permite recorrer la lista sin exponer su estructura.
type Iterador struct {
	actual  *Nodo
	primera bool
}

// Iterador devuelve un iterador nuevo posicionado antes del primer nodo.
func (l *Lista) Iterador() *Iterador {
	return &Iterador{actual: l.cabeza, primera: true}
}

// Siguiente avanza al siguiente nodo y devuelve true si hay uno.
// La primera llamada posiciona el iterador en la cabeza.
func (it *Iterador) Siguiente() bool {
	if it.actual == nil {
		return false
	}
	if !it.primera {
		it.actual = it.actual.sig
		if it.actual == nil {
			return false
		}
	}
	it.primera = false
	return true
}

// Valor devuelve el valor del nodo actual.
//
// Precondición: solo debe llamarse después de que Siguiente haya devuelto
// true.
func (it *Iterador) Valor() int {
	return it.actual.valor
}
