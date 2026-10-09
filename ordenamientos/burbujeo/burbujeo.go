package main

// OrdenarBurbujeo ordena el slice en orden ascendente con el algoritmo de
// burbujeo: compara elementos adyacentes y los intercambia si están en el
// orden incorrecto, de modo que los más grandes "burbujean" hacia el final.
// Optimización: se detiene si en una pasada completa no hubo intercambios.
// Es estable e in-place.
// Complejidad: O(n²) peor caso, O(n) mejor caso.
func OrdenarBurbujeo(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		intercambiado := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				intercambiado = true
			}
		}
		if !intercambiado {
			break
		}
	}
}
