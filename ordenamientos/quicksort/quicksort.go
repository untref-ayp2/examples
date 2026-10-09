package main

// Quicksort ordena el slice de enteros de manera ascendente usando el
// algoritmo de partición del mismo nombre: elige un pivote, particiona el
// arreglo dejándolo en su posición final y repite recursivamente con las dos
// sublistas. Los intercambios suceden sobre el mismo slice (in-place) y se
// devuelve ese mismo slice ya ordenado.
// Complejidad: O(n log n) promedio, O(n²) peor caso.
func Quicksort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	pivot := array[0]   // primer elemento como pivote
	i := 1              // avanza desde el segundo elemento
	j := len(array) - 1 // retrocede desde el último elemento

	// particionamiento: menores del pivote a la izquierda, mayores a la derecha
	for i < j {
		for array[i] < pivot && i < len(array)-1 {
			i++
		}
		for array[j] > pivot && j > 0 {
			j--
		}
		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	// el pivote queda en su posición final
	if array[j] < pivot {
		array[0], array[j] = array[j], array[0]
	}

	Quicksort(array[:j])   // sublista izquierda
	Quicksort(array[j+1:]) // sublista derecha

	return array
}
