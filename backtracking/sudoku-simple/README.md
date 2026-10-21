# Sudoku (versión simple)

## 1. El problema

Completar un tablero de Sudoku de 9×9: en cada fila, cada columna y cada
cuadrante de 3×3 tienen que estar los dígitos del 1 al 9 sin repetirse. Los
ceros son las celdas vacías.

## 2. La estrategia

Backtracking sobre la primera celda vacía: se prueban los valores del 1 al 9;
si el valor no choca con la fila, la columna ni el cuadrante, se registra y se
resuelve el resto del tablero recursivamente. Si ninguno funciona, se deshace
la asignación (**vuelta atrás**) y se prueba otra rama.

El algoritmo agota todas las ramas para juntar **todas** las soluciones; la
función `Resolver` solo informa si hay al menos una.

## 3. Qué hace el código

| Función | Qué hace |
|---|---|
| `NuevoSudoku(tablero [9][9]int) *Sudoku` | Crea el resolutor con una copia del tablero. |
| `(*Sudoku).Resolver() bool` | Busca todas las soluciones, las guarda en `soluciones` y devuelve si hay alguna. |
| `(*Sudoku).Tablero() string` | Impresión estática del tablero con bordes. |
| `(*Sudoku).backtracking() bool` | La recursión del esquema: `esSolucion`, `proximaVacia`, `esFactible`, `registrar`, `borrar` y extender. |

Es el mismo esquema de la versión animada en [`../sudoku/`](../sudoku/), sin el
callback que notifica cada asignación.

## 4. Cómo correrlo

```bash
go run ./backtracking/sudoku-simple/
```

## 5. Para experimentar

- Sacá alguna pista del tablero de `main.go` y mirá cómo cambia la cantidad de
  soluciones: el tablero tiene solución única, pero no todas las pistas son
  indispensables. Por ejemplo, si sacás la de la fila 1, columna 6 (un `5`) o la
  de la fila 2, columna 7 (un `7`), van a aparecer soluciones nuevas; si sacás
  la de la esquina superior izquierda (un `4`), el tablero sigue teniendo una
  sola solución.
- Armá un tablero imposible de completar (por ejemplo, en la esquina
  superior derecha de `{1, 2, 3, 4, 5, 6, 7, 8, 0}` no puede ir ningún
  dígito) y comprobá que `Resolver` devuelve `false`.
- Compará el tiempo de `Resolver` con el de la versión animada: la animación no
  cambia el algoritmo, solo agrega la espera entre pasos.

## 6. Tests

```bash
go test ./backtracking/sudoku-simple/
```

## 7. Complejidad

En el peor caso exponencial: se prueban hasta 9 valores por cada una de las 81
celdas. `esFactible` y el orden de llenado (siempre la primera celda vacía)
poda la mayor parte del árbol en tableros reales.
