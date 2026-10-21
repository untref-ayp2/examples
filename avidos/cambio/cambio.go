package main

// Cambiar da cambio para cantidad usando el menor número de billetes y
// monedas posible, y devuelve un map con la cantidad de piezas de cada
// denominación (las denominaciones sin ninguna pieza no aparecen en el map).
//
// El ávido acá es óptimo porque las denominaciones son canónicas (cada una
// es divisible por las menores): entregar siempre la pieza más grande posible
// nunca obliga a usar más piezas que cualquier otra combinación.
//
// Precondición: cantidad >= 0 (con negativo no se entrega nada). Las
// denominaciones están fijas en el código, de mayor a menor.
func Cambiar(cantidad int) map[int]int {
	billetes := []int{10000, 2000, 1000, 500, 200, 100, 50, 20, 10, 5, 2, 1}
	cambio := make(map[int]int)
	for _, denominacion := range billetes {
		if cantidad >= denominacion {
			cambio[denominacion] = cantidad / denominacion
			cantidad = cantidad % denominacion
		}
	}
	return cambio
}
