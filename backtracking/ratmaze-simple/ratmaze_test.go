package main

import "testing"

// TestRatMazeVuelveAtrasDeVerdad comprueba la huella real de la vuelta atrás
// sobre un laberinto chico donde la búsqueda entra a un callejón sin salida:
//
//	       x=0    x=1    x=2
//	y=0  [ libre  libre  pared ]
//	y=1  [ libre  pared  pared ]
//	y=2  [ libre  libre  libre ]
//
// La rata sale de (0,0) y prueba las direcciones en el orden derecha, abajo,
// izquierda, arriba. Primero entra a (1,0), descubre que desde ahí no hay
// salida y la abandona; recién después baja por la izquierda y llega al
// destino. Ese abandono es el corazón del backtracking y queda registrado en
// los dos tableros del resolutor:
//
//   - solucion: la celda abandonada NO queda marcada (solucion es el camino
//     actual y esa celda dejó de pertenecerle).
//   - visitado: la celda abandonada SÍ queda marcada (visitado es la memoria
//     de lo ya explorado y no se borra nunca).
//
// Si borrar no deshiciera la marca, (1,0) seguiría en solucion y el test
// fallaría: por eso estas aserciones sí detectan la falta de vuelta atrás.
func TestRatMazeVuelveAtrasDeVerdad(t *testing.T) {
	maze := [][]bool{
		{true, true, false},
		{true, false, false},
		{true, true, true},
	}

	ratmaze := NuevoRatMaze(maze)
	if !ratmaze.Resolver() {
		t.Fatal("Resolver() = false, se esperaba encontrar un camino")
	}

	// (1,0) es el callejón sin salida: se explora y se abandona.
	if ratmaze.visitado[0][1] != true {
		t.Error("la celda (1,0) debería quedar marcada en visitado: se exploró")
	}
	if ratmaze.solucion[0][1] != false {
		t.Error("la celda (1,0) no debería quedar marcada en solucion: " +
			"se abandonó, y la vuelta atrás tiene que desmarcarla")
	}

	// (0,1), en cambio, es la que sí lleva al destino: tiene que quedar en
	// el camino.
	if ratmaze.solucion[1][0] != true {
		t.Error("la celda (0,1) debería quedar marcada en solucion: lleva al destino")
	}

	// Y ninguna pared puede estar marcada en ninguno de los dos tableros.
	if ratmaze.solucion[1][1] || ratmaze.solucion[0][2] {
		t.Error("las paredes (1,1) y (2,0) no deberían estar en solucion")
	}
	if ratmaze.visitado[1][1] || ratmaze.visitado[0][2] {
		t.Error("las paredes (1,1) y (2,0) no deberían estar en visitado")
	}

	// Vista global: si hubo al menos una vuelta atrás, visitado tiene más
	// celdas marcadas que solucion (la abandonada quedó sólo en visitado).
	// Con el camino real son 6 contra 5; sin deshacer, solucion arrastraría
	// la celda abandonada y quedarían 6 y 6.
	if contarMarcadas(ratmaze.visitado) <= contarMarcadas(ratmaze.solucion) {
		t.Errorf("visitado (%d celdas) debería tener más marcas que solucion "+
			"(%d celdas): hay al menos una celda explorada y abandonada",
			contarMarcadas(ratmaze.visitado), contarMarcadas(ratmaze.solucion))
	}
}

// contarMarcadas cuenta cuántas celdas tienen true en un tablero.
func contarMarcadas(tablero [][]bool) int {
	marcadas := 0
	for y := range tablero {
		for x := range tablero[y] {
			if tablero[y][x] {
				marcadas++
			}
		}
	}
	return marcadas
}
