# Ejemplo: Sudoku

## 1. El problema

Resolver un tablero de Sudoku 9 × 9: completar las celdas vacías con los
dígitos del 1 al 9 sin repetir ninguno en la misma fila, la misma columna ni
el mismo cuadrante de 3 × 3. Las celdas vacías se representan con `0`.

## 2. La estrategia

Backtracking con el esquema general del apunte: se busca la próxima celda
vacía y se prueban los valores del 1 al 9 (`extender`). Un valor es factible
si no se repite en su fila, su columna ni su cuadrante (`esFactible`); si
funciona se registra (`registrar`) y se resuelve el resto del tablero; si
ningún valor sirve, se deshace la asignación (`borrar`) y se vuelve atrás.

## 3. Qué hace el código

| Función | Qué hace |
|---|---|
| `NuevoSudoku(tablero [9][9]int) *Sudoku` | Crea el tablero. |
| `Resolver() bool` | Busca **todas** las soluciones y las deja en `soluciones`. Devuelve si existe al menos una. El tablero original no se modifica. |

El núcleo del algoritmo son las funciones del esquema:
`esSolucion()`, `esFactible(x, y, n)`, `registrar(x, y, n)` y
`borrar(x, y)`, con la recursión general en
`backtracking(paso, primera)` — el callback `paso` permite observar cada
asignación desde afuera sin que el algoritmo conozca la terminal.

## 4. Visualización

`sudoku_visual.go` es la capa de dibujo, completamente separada del
backtracking:

- `Tablero(x, y, pasos) string` dibuja el tablero con bordes, muestra en
  gris las celdas vacías y resalta la celda actual.
- `Animar(w, demora)` resuelve mostrando cada paso con una pausa y termina
  mostrando la solución encontrada (o avisa si el tablero no tiene solución).
- `Dibujar(w, x, y, pasos)` escribe el tablero en `w`.

## 5. Cómo correrlo

```bash
go run ./backtracking/sudoku/
```

Corre en la terminal una animación de la resolución paso a paso.

## 6. Para experimentar

El tablero del ejemplo está en `main.go`, en la variable `tablero`. Las
celdas vacías se escriben con `0` y el resto son las pistas:

```go
tablero := [9][9]int{
	{4, 0, 0, 1, 9, 5, 0, 6, 8},
	// ...
}
```

Cambiá las pistas y volvé a correr
`go run ./backtracking/sudoku/`. Algunas ideas:

- **Poné más ceros**: quedan más celdas libres, más soluciones y muchos más
  pasos; se ve mejor cómo el algoritmo prueba, se arrepiente y vuelve atrás.
- **Armá un tablero sin salida**: la animación busca, no encuentra y avisa
  que el tablero no tiene solución.
- **Dejá el tablero casi resuelto**: la solución aparece en pocos pasos,
  útil para ver el final sin esperar.

El último argumento de `Animar` (en `main.go`) es la pausa entre pasos:
bajala para que sea más rápido o subila para mirarlo con calma.

## 7. Tests

```bash
go test ./backtracking/sudoku/
```

Verifican que las soluciones respeten las reglas del sudoku y las pistas
originales, con tableros de una, dos y muchísimas soluciones (¡32 604!).
Agregá tus propios tableros y averiguá cuántas soluciones tienen.

## 8. Complejidad

Exponencial en el peor caso (cada celda vacía puede probar hasta 9 valores),
pero la poda por fila, columna y cuadrante recorta muchísimo el árbol: los
tableros habituales se resuelven en un tiempo práctico. La memoria es
`O(9²)` por el tablero más la cantidad de soluciones guardadas.
