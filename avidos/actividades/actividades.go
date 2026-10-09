package main

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

// SelectorActividadesIterativo devuelve la mayor cantidad de actividades que
// no se solapan, recorriendo la lista en una sola pasada.
//
// Precondición: las actividades están ordenadas de menor a mayor por tiempo de
// finalización.
func SelectorActividadesIterativo(actividades []Actividad) []Actividad {
	var seleccionadas []Actividad
	n := len(actividades)
	seleccionadas = append(seleccionadas, actividades[0])
	k := 0
	for i := 1; i < n; i++ {
		if actividades[i].Inicio >= actividades[k].Fin {
			seleccionadas = append(seleccionadas, actividades[i])
			k = i
		}
	}
	return seleccionadas
}
