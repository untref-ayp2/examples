package main

// Mergesort ordena el slice de enteros de manera ascendente usando el
// algoritmo divide y conquista del mismo nombre: divide el arreglo en dos
// mitades, las ordena recursivamente y las fusiona. Es estable.
// No modifica el arreglo original: devuelve un nuevo slice ordenado.
// Complejidad: O(n log n) en todos los casos.
func Mergesort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	middle := len(array) / 2

	left := Mergesort(array[:middle])  // ordena la mitad izquierda
	right := Mergesort(array[middle:]) // ordena la mitad derecha

	return merge(left, right)
}

// merge combina dos slices ya ordenados en uno solo ordenado.
func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	array := make([]int, size)

	for k := 0; k < size; k++ {
		switch {
		case i > len(left)-1:
			array[k] = right[j]
			j++
		case j > len(right)-1:
			array[k] = left[i]
			i++
		case left[i] <= right[j]:
			// con elementos iguales tomo el de la izquierda: así el
			// ordenamiento es estable (igual que la versión visual)
			array[k] = left[i]
			i++
		default:
			array[k] = right[j]
			j++
		}
	}
	return array
}
