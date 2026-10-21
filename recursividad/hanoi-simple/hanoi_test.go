package main

import (
	"fmt"
	"strings"
	"testing"
)

// TestMoverConTresDiscos verifica la secuencia exacta de movimientos para la
// torre clásica de 3 discos.
func TestMoverConTresDiscos(t *testing.T) {
	movimientos := Mover(3, "A", "B", "C")

	esperado := "1:A->C,2:A->B,1:C->B,3:A->C,1:B->A,2:B->C,1:A->C"

	var partes []string
	for _, m := range movimientos {
		partes = append(partes, fmt.Sprintf("%d:%s->%s", m.Disco, m.Origen, m.Destino))
	}
	got := strings.Join(partes, ",")

	if got != esperado {
		t.Errorf("secuencia de movimientos = %s; se esperaba %s", got, esperado)
	}
}

// TestCantidadDeMovimientos verifica que con n discos se generan 2^n - 1
// movimientos.
func TestCantidadDeMovimientos(t *testing.T) {
	casos := []struct {
		n      int
		quiero int
	}{
		{1, 1}, {2, 3}, {3, 7}, {4, 15}, {5, 31},
	}

	for _, caso := range casos {
		movimientos := Mover(caso.n, "A", "B", "C")
		if len(movimientos) != caso.quiero {
			t.Errorf("Mover(%d, ...) generó %d movimientos; se esperaban %d",
				caso.n, len(movimientos), caso.quiero)
		}
	}
}

// TestMoverTerminaEnLaVaraDestino comprueba que el disco más grande termina
// en la vara destino: se busca el movimiento que lo lleva allí, porque el
// disco grande se mueve una sola vez, a mitad de la secuencia (después van los
// n-1 discos de arriba encima suyo) y nunca es el último movimiento.
func TestMoverTerminaEnLaVaraDestino(t *testing.T) {
	movimientos := Mover(4, "A", "B", "C")

	var trasladoDelGrande *Movimiento
	for i := range movimientos {
		if movimientos[i].Disco == 4 {
			trasladoDelGrande = &movimientos[i]
		}
	}

	if trasladoDelGrande == nil {
		t.Fatal("no hay ningún movimiento del disco 4")
	}
	if trasladoDelGrande.Destino != "C" {
		t.Errorf("el disco 4 debería terminar en C, fue %+v", *trasladoDelGrande)
	}
}

// TestMoverNoUsaLaVaraAuxiliarComoDestinoFinal comprueba que ningún
// movimiento deje un disco grande encima de uno chico: la simulación aplica
// las reglas del problema y al final todos los discos quedan en la destino.
func TestMoverNoUsaLaVaraAuxiliarComoDestinoFinal(t *testing.T) {
	var origen, aux, destino []int
	origen = []int{3, 2, 1} // arriba de todo está el disco 1

	for _, m := range Mover(3, "A", "B", "C") {
		var desde, hacia *[]int
		switch m.Origen {
		case "A":
			desde = &origen
		case "B":
			desde = &aux
		case "C":
			desde = &destino
		}
		switch m.Destino {
		case "A":
			hacia = &origen
		case "B":
			hacia = &aux
		case "C":
			hacia = &destino
		}

		// El disco que se mueve tiene que ser el de arriba de su vara.
		if len(*desde) == 0 || (*desde)[len(*desde)-1] != m.Disco {
			t.Fatalf("el movimiento %+v no mueve el disco de arriba", m)
		}
		*desde = (*desde)[:len(*desde)-1]

		// Nunca un disco grande puede quedar encima de uno chico.
		if len(*hacia) > 0 && (*hacia)[len(*hacia)-1] < m.Disco {
			t.Fatalf("el movimiento %+v pone el disco %d encima de uno más chico", m, m.Disco)
		}
		*hacia = append(*hacia, m.Disco)
	}

	if len(destino) != 3 || len(origen) != 0 || len(aux) != 0 {
		t.Errorf("al final los discos deberían estar todos en C: A=%v B=%v C=%v",
			origen, aux, destino)
	}
}
