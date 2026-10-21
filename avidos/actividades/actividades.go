package main

// Actividad es una tarea con un momento de inicio y uno de finalización. Dos
// actividades se solapan si sus intervalos [Inicio, Fin) comparten tiempo.
type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

// SelectorActividadesIterativo devuelve la mayor cantidad de actividades que
// no se solapan, recorriendo la lista en una sola pasada.
//
// Estrategia ávida: se toma siempre la actividad que termina antes, porque
// deja libre la mayor parte de la línea de tiempo para las que siguen. Es
// óptima: en cualquier selección, reemplazar la primera actividad elegida por
// la que termina antes no elimina ninguna posibilidad futura.
//
// Precondición: la lista no está vacía y las actividades están ordenadas de
// menor a mayor por tiempo de finalización.
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
