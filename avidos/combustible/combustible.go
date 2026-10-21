package main

// MinimasRecargas recibe la distancia a recorrer, la capacidad del tanque y
// las estaciones de servicio en orden de aparición en la ruta, y devuelve la
// cantidad mínima de recargas necesarias para recorrer la distancia.
//
// Precondición: las estaciones vienen ordenadas de menor a mayor y sus
// posiciones están dentro de la ruta; distancia > 0 y capacidad > 0.
//
// La estrategia ávida (parar siempre en la estación más lejana alcanzable)
// da la solución óptima: si en un tramo se para antes de la más lejana, esa
// parada es reemplazable por la más lejana sin perder ninguna otra opción, y
// el subproblema que queda es más chico.
func MinimasRecargas(distancia int, capacidad int, estaciones []int) int {
	return len(paradasDeRecarga(distancia, capacidad, estaciones))
}

// paradasDeRecarga aplica la estrategia ávida y devuelve las estaciones donde
// se recarga, en orden. La idea es llegar lo más lejos posible con cada
// recarga: desde la ubicación actual se avanza hasta la última estación
// alcanzable con el tanque y se para ahí.
func paradasDeRecarga(distancia int, capacidad int, estaciones []int) []int {
	var paradas []int
	ubicacion := 0

	for ubicacion+capacidad < distancia {
		// No quedan estaciones en la ruta: ya no se puede recargar, así
		// que la distancia no se puede cubrir (el bucle corta acá).
		if len(estaciones) == 0 {
			break
		}

		// Se avanza hasta la última estación alcanzable con el tanque
		// actual.
		ultimaParada := ubicacion
		for len(estaciones) > 0 && capacidad >= estaciones[0]-ubicacion {
			ultimaParada = estaciones[0]
			estaciones = estaciones[1:]
		}

		// No había ninguna estación alcanzable: no se puede avanzar.
		if ultimaParada == ubicacion {
			break
		}

		paradas = append(paradas, ultimaParada)
		ubicacion = ultimaParada
	}

	return paradas
}
