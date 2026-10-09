package main

// Solución al problema del cambio de moneda con backtracking.
// cantidad: cantidad de dinero a cambiar.
// denominaciones: monedas disponibles para el cambio.
// Devuelve un slice con las monedas que dan el cambio usando la menor
// cantidad de piezas posibles.

// ObservadorCambio recibe los eventos de la búsqueda. Cualquiera de los dos
// callbacks puede ser nil.
type ObservadorCambio struct {
	// Intento se invoca antes de probar cada moneda de una denominación:
	// resto es lo que queda por cambiar con la solución parcial actual.
	Intento func(resto int, parcial []int)
	// Solucion se invoca al encontrar una combinación que da el monto
	// exacto. solucion es la combinación encontrada y mejor la mejor
	// solución hallada hasta ahora.
	Solucion func(solucion, mejor []int)
}

// Función recursiva que realiza el backtracking para encontrar la solución.
func backtracking(cantidad int, denominaciones []int, solucion []int, mejorSolucion *[]int, obs ObservadorCambio) {
	// Si ya encontramos una solución, comparamos si es mejor que la actual
	if cantidad == 0 { // Encontramos una solución
		if len(solucion) < len(*mejorSolucion) || len(*mejorSolucion) == 0 {
			// *mejorSolucion = solucion no funciona porque comparten memoria
			*mejorSolucion = make([]int, len(solucion))
			copy(*mejorSolucion, solucion)
		}
		if obs.Solucion != nil {
			obs.Solucion(solucion, *mejorSolucion)
		}
		return
	}
	// Poda: si esta rama ya tiene igual o más piezas que la mejor, no puede
	// mejorar la solución actual.
	if len(solucion) >= len(*mejorSolucion) && len(*mejorSolucion) != 0 {
		return
	}
	// Si no encontramos una solución, seguimos explorando todas las posibilidades
	for i := 0; i < len(denominaciones); i++ { // extender
		if denominaciones[i] <= cantidad { // es factible entregar una moneda de este tipo
			nuevaSolucion := append(solucion, denominaciones[i]) // registrar
			if obs.Intento != nil {
				obs.Intento(cantidad-denominaciones[i], nuevaSolucion)
			}
			backtracking(cantidad-denominaciones[i], denominaciones[i:], nuevaSolucion, mejorSolucion, obs)
		} // borrar implícito porque usé nuevaSolucion
	}
}

// CambioDeMonedaConPaso resuelve el cambio notificando cada evento a obs.
func CambioDeMonedaConPaso(cantidad int, denominaciones []int, obs ObservadorCambio) []int {
	var solucion []int
	var mejorSolucion []int

	backtracking(cantidad, denominaciones, solucion, &mejorSolucion, obs)
	return mejorSolucion
}

// CambioDeMoneda devuelve el cambio para cantidad usando la menor cantidad
// de piezas posible.
func CambioDeMoneda(cantidad int, denominaciones []int) []int {
	return CambioDeMonedaConPaso(cantidad, denominaciones, ObservadorCambio{})
}
