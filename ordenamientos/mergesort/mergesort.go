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

	// Cada mitad se ordena recursivamente y después se fusionan. Las
	// divisiones array[:middle] y array[middle:] son vistas, no copias.
	left := Mergesort(array[:middle])  // ordena la mitad izquierda
	right := Mergesort(array[middle:]) // ordena la mitad derecha

	return merge(left, right)
}

// merge combina dos slices ya ordenados en uno solo ordenado.
func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	array := make([]int, size)

	// i y j son los índices que avanzan por left y right.
	for k := 0; k < size; k++ {
		switch {
		case i > len(left)-1:
			// Se acabó la mitad izquierda: queda copiar el resto de right.
			array[k] = right[j]
			j++
		case j > len(right)-1:
			// Se acabó la mitad derecha: queda copiar el resto de left.
			array[k] = left[i]
			i++
		case left[i] <= right[j]:
			// Con elementos iguales se toma el de la izquierda: así el
			// ordenamiento es estable (igual que la versión visual).
			array[k] = left[i]
			i++
		default:
			array[k] = right[j]
			j++
		}
	}
	return array
}
