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
	pivot := array[0]
	i := 1 // borde: todo lo que está antes de i es menor que el pivote
	for j := 1; j < len(array); j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	// el pivote va al borde de los menores y queda en su posición final
	array[0], array[i-1] = array[i-1], array[0]

	Quicksort(array[:i-1]) // sublista izquierda (menores del pivote)
	Quicksort(array[i:])   // sublista derecha (mayores o iguales)

	return array
}
