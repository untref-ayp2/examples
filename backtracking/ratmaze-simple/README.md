# Laberinto de la rata (versión simple)

## 1. El problema

Dado un laberinto en forma de matriz donde `true` es celda libre y `false` es
pared, encontrar un camino que vaya desde la celda superior izquierda hasta la
inferior derecha moviéndose en las cuatro direcciones.

## 2. La estrategia

Backtracking: desde la celda actual se prueba avanzar en cada dirección; si
alguna lleva al destino, ese camino se conserva y la búsqueda corta. Si ninguna
dirección funciona, se desmarca la celda (**vuelta atrás**) y se devuelve
`false` a quien la había marcado.

El orden en que se prueban las direcciones (derecha, abajo, izquierda, arriba)
define cuál de todos los caminos posibles se encuentra: el algoritmo no
garantiza el más corto. En el laberinto del `main` el camino hallado mide 21
celdas y sí resulta ser el más corto (coincide con el que daría una búsqueda en
anchura), pero eso es casualidad de ese mapa: con otro orden de direcciones o
otro laberinto puede devolver uno más largo.

## 3. Qué hace el código

| Función | Qué hace |
|---|---|
| `NuevoRatMaze(maze [][]bool) *RatMaze` | Crea el resolutor para el laberinto dado. |
| `(*RatMaze).Resolver() bool` | Busca un camino y deja marcadas sus celdas en `solucion`. Devuelve si encontró alguno. |
| `(*RatMaze).Tablero() string` | Impresión estática del laberinto: camino con `o`, paredes con `#` y libres con `.`. |
| `(*RatMaze).backtracking(x, y int) bool` | La recursión del esquema: `esSolucion`, poda (`esFactible`), `registrar`, extender y `borrar`. |

Es el mismo esquema de la versión animada en [`../ratmaze/`](../ratmaze/), sin
el callback `PasoLaberinto`: acá la búsqueda no avisa cada paso, solo deja el
resultado.

## 4. Cómo correrlo

```bash
go run ./backtracking/ratmaze-simple/
```

## 5. Para experimentar

- Tapá la única salida del laberinto y mirá cómo `Resolver` devuelve `false`.
- Cambiá el orden de las cuatro llamadas recursivas y compará los caminos.
- Después de `Resolver`, mirá `visitado`: las celdas marcadas que no están en
  `solucion` son las que se exploraron y se abandonaron.

## 6. Tests

```bash
go test ./backtracking/ratmaze-simple/
```

## 7. Complejidad

Cuando hay solución y las direcciones encuentran el destino rápido, el costo es
del orden de la cantidad de celdas: `O(filas × columnas)`.

En el peor caso, en cambio, no alcanza con un orden polinomial: si el destino
está encerrado (`Resolver` devuelve `false`), la búsqueda recorre todos los
caminos simples antes de descartarlos y el árbol crece exponencialmente. Por
ejemplo, con un tablero abierto de `n × n` cuyo destino es una pared, la
cantidad de llamadas a `backtracking` crece así: 3×3 → 157, 4×4 → 4349,
5×5 → 326397. La memoria sí es acotada: `O(filas × columnas)` por los dos
tableros auxiliares (`solucion` y `visitado`) más el alto de la recursión.
