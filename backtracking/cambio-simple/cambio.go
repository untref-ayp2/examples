package main

// Este archivo contiene el backtracking del problema del cambio de moneda,
// sin animación: la búsqueda no notifica nada, solo devuelve el resultado.
//
// El esquema es el mismo que el de la versión animada (carpeta cambio/):
// caso base, poda, extender, registrar y borrar. Acá el "borrar" es
// implícito: cada rama construye su propio slice con append en vez de
// modificar el actual.

// CambioDeMoneda devuelve el cambio para cantidad usando la menor cantidad
// de piezas posible —o nil si ninguna combinación de denominaciones suma
// exactamente cantidad.
//
// Precondiciones: denominaciones no está vacía y todas sus monedas son
// positivas. cantidad es >= 0 (con cantidad 0 el cambio es vacío y válido).
func CambioDeMoneda(cantidad int, denominaciones []int) []int {
	var mejorSolucion []int
	backtracking(cantidad, denominaciones, nil, &mejorSolucion)
	return mejorSolucion
}

// backtracking es la recursión general del esquema.
//
// Parámetros: cantidad es lo que queda por cambiar en esta rama;
// denominaciones son las monedas que todavía se pueden usar (de ahí en
// adelante del slice original); solucion es la combinación parcial y
// mejorSolucion apunta a la mejor combinación encontrada hasta ahora (por
// eso se recibe como puntero: la recursión la va actualizando).
//
// Termina porque en cada llamada recursiva cantidad disminuye en al menos 1
// (las denominaciones son positivas), y corta al llegar al caso base
// cantidad == 0 o al activarse la poda.
func backtracking(cantidad int, denominaciones []int, solucion []int, mejorSolucion *[]int) {
	// Caso base (esSolucion): la combinación actual da el monto exacto.
	if cantidad == 0 {
		if len(solucion) < len(*mejorSolucion) || len(*mejorSolucion) == 0 {
			// Hay que copiar: solucion se sigue compartiendo con las
			// ramas que quedan y se pisaría si se guardara el slice.
			*mejorSolucion = make([]int, len(solucion))
			copy(*mejorSolucion, solucion)
		}
		return
	}

	// Poda: esta rama ya tiene igual o más piezas que la mejor conocida,
	// así que no puede mejorarla; seguir explorándola es trabajo perdido.
	if len(*mejorSolucion) != 0 && len(solucion) >= len(*mejorSolucion) {
		return
	}

	// Extender: probar cada denominación que todavía no se haya usado
	// (denominaciones[i:] evita repetir combinaciones en distinto orden).
	for i := 0; i < len(denominaciones); i++ {
		if denominaciones[i] > cantidad { // esFactible: la moneda no se pasa del resto
			continue
		}
		// Registrar: agregar la moneda a la combinación parcial. Con
		// append se crea una nueva rama, así el "borrar" es implícito.
		nuevaSolucion := append(solucion, denominaciones[i])
		backtracking(cantidad-denominaciones[i], denominaciones[i:], nuevaSolucion, mejorSolucion)
	}
}
