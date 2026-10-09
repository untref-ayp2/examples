package main

// MinimasRecargas recibe la distancia a recorrer, la capacidad del tanque y
// las estaciones de servicio en orden de aparición en la ruta, y devuelve la
// cantidad mínima de recargas necesarias para recorrer la distancia.
func MinimasRecargas(distancia int, capacidad int, estaciones []int) int {
	return len(paradasDeRecarga(distancia, capacidad, estaciones))
}

// paradasDeRecarga aplica la estrategia ávida y devuelve las estaciones donde
// se recarga, en orden. La idea es llegar lo más lejos posible con cada
// recarga: desde donde estoy, avanzo hasta la última estación alcanzable con
// el tanque actual y paro ahí.
func paradasDeRecarga(distancia int, capacidad int, estaciones []int) []int {
	var paradas []int
	ubicacion := 0

	for ubicacion+capacidad < distancia {
		// alcanza el tanque para llegar al final: no hace falta recargar
		if len(estaciones) == 0 {
			break
		}

		// me muevo hasta la última estación a la que puedo llegar con el
		// tanque actual
		ultimaParada := ubicacion
		for len(estaciones) > 0 && capacidad >= estaciones[0]-ubicacion {
			ultimaParada = estaciones[0]
			estaciones = estaciones[1:]
		}

		// no había ninguna estación alcanzable: no se puede avanzar
		if ultimaParada == ubicacion {
			break
		}

		paradas = append(paradas, ultimaParada)
		ubicacion = ultimaParada
	}

	return paradas
}
