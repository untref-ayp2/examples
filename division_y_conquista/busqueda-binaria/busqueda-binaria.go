package main

// BusquedaBinaria busca x en un arreglo ordenado de menor a mayor y devuelve
// el índice donde está, o -1 si no aparece.
//
// Divide el problema a la mitad en cada paso: compara x con el elemento del
// medio y descarta la mitad que no puede contenerlo. Por eso es mucho más
// rápida que la búsqueda lineal, pero exige que el arreglo esté ordenado.
//
// Precondición: array está ordenado de menor a mayor.
// Complejidad: O(log n).
func BusquedaBinaria(array []int, x int) int {
	return busquedaBinaria(array, 0, len(array)-1, x)
}

// busquedaBinaria es la recursión: busca x dentro de array[inicio..fin]
// (ambos inclusive).
//
// Caso base: inicio > fin significa que el rango quedó vacío, es decir que x
// no está en el arreglo; devuelve -1. Termina porque en cada llamada el rango
// se parte a la mitad y se descarta el medio: el rango se achica siempre.
func busquedaBinaria(array []int, inicio int, fin int, x int) int {
	if inicio > fin {
		return -1
	}

	medio := (inicio + fin) / 2

	// El elemento del medio es más grande que x: x (si está) sólo puede
	// estar a la izquierda, así que descarto medio y todo lo que sigue.
	if array[medio] > x {
		return busquedaBinaria(array, inicio, medio-1, x)
	}

	// Al revés: x sólo puede estar a la derecha; descarto medio y lo anterior.
	if array[medio] < x {
		return busquedaBinaria(array, medio+1, fin, x)
	}

	// No es mayor ni menor: es igual, y el medio es la posición buscada.
	return medio
}
