# Ejemplo: Laberinto (Rat in a Maze)

## 1. El problema

Encontrar un camino desde la celda superior izquierda `(0, 0)` hasta la
inferior derecha de un laberinto, moviéndose en las cuatro direcciones por las
celdas libres. El laberinto es una matriz de booleanos: `true` es una celda
libre y `false` una pared.

## 2. La estrategia

Backtracking con el esquema general del apunte: desde la celda actual se
marcan todas las direcciones posibles (`extender`) y se avanza por la primera
que lleve a la salida; si ninguna funciona, se desmarca la celda (`borrar`) y
se vuelve atrás. La poda (`esFactible`) descarta las celdas fuera del
laberinto, las paredes y las celdas que ya forman parte del camino actual.

## 3. Qué hace el código

| Función | Qué hace |
|---|---|
| `NuevoRatMaze(maze [][]bool) *RatMaze` | Crea el laberinto. |
| `Resolver() bool` | Busca un camino desde `(0, 0)` hasta la inferior derecha y devuelve si encontró alguno. Las celdas que forman el camino quedan marcadas en `solucion`. |
| `ResolverConPaso(paso PasoLaberinto) bool` | Igual que `Resolver`, pero invoca `paso(x, y, avanzando)` en **cada** avance (`true`) y cada vuelta atrás (`false`). |
| `PasoLaberinto` | Tipo de función `func(x, y int, avanzando bool)` usada como gancho de la visualización; puede ser `nil`. |

El núcleo del algoritmo son las funciones del esquema:
`esSolucion(x, y)`, `esFactible(x, y)`, `registrar(x, y, paso)` y
`borrar(x, y, paso)`, con la recursión general en `backtracking(x, y, paso)`.

## 4. Visualización

`ratmaze_visual.go` es la capa de dibujo, separada del backtracking:

- `Tablero(actualX, actualY) string` devuelve el laberinto en ASCII: `#`
  pared, `.` libre, `o` camino actual, `x` celda explorada y abandonada y
  `@` el cursor.
- `Animar(w, demora)` dibuja cada paso del recorrido con una pausa y termina
  mostrando el camino encontrado (o avisando que no existe).
- `Dibujar(w)` escribe el tablero sin cursor.

## 5. Cómo correrlo

```bash
go run ./backtracking/ratmaze/
```

Anima un laberinto grande: se ve cómo la rata avanza, se mete en callejones
sin salida y vuelve atrás hasta encontrar el camino.

## 6. Para experimentar

El laberinto del ejemplo está en `main.go`, en la variable `maze`: cada
`true` es una celda libre y cada `false` una pared. El origen es siempre la
celda de arriba a la izquierda `(0, 0)` y el destino la de abajo a la
derecha.

```go
maze := [][]bool{
	{true, true, true, /* ... */},
	// ...
}
```

Armá tu propio laberinto y correlo con
`go run ./backtracking/ratmaze/`. Algunas ideas:

- **Agregá callejones sin salida**: pasillos que terminan en una pared. Se
  ve cómo la rata se mete, se queda sin opciones y vuelve atrás.
- **Tapá la salida** (poné `false` en la esquina inferior derecha) para
  comprobar que el programa avisa que no hay camino.
- **Agrandalo**: cuantas más celdas libres, más pasos y más larga la
  animación (y también más lo que tarda).

Mantené todas las filas del mismo largo. El último argumento de `Animar` (en
`main.go`) es la pausa entre pasos.

## 7. Tests

```bash
go test ./backtracking/ratmaze/
```

Verifican que la solución sea un camino simple y conexo desde el origen
hasta el destino, y el caso en que no existe camino. Agregá tus propios
laberintos y comprobá qué pasa.

## 8. Complejidad

Exponencial en el peor caso (cada celda libre puede generar varias ramas),
acotada en la práctica por la poda de celdas ya visitadas y por la
estructura del laberinto: los caminos sin salida se descartan rápido con la
vuelta atrás. La memoria es `O(m·n)` con `m×n` el tamaño del laberinto.
