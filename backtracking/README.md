# Backtracking

En esta carpeta se encuentran los ejemplos de backtracking que se ven en clase y
se resuelven en el aula. Cada ejemplo vive en su propia carpeta y se puede
ejecutar y testear de forma independiente.

## Ejemplos

| Ejemplo | Descripción | Cómo correrlo |
|---|---|---|
| [cambio](cambio/) | Cambio de moneda por backtracking: explora las combinaciones de denominaciones. | `go run ./backtracking/cambio/` |
| [ratmaze](ratmaze/) | Laberinto: encuentra un camino desde la entrada hasta la salida. | `go run ./backtracking/ratmaze/` |
| [reinas](reinas/) | Problema de las n reinas: ubica n reinas sin que se ataquen. | `go run ./backtracking/reinas/` |
| [sudoku](sudoku/) | Resolución de un tablero de Sudoku de 9×9. | `go run ./backtracking/sudoku/` |

Los ejemplos de `ratmaze`, `reinas` y `sudoku` incluyen una capa visual ASCII que
muestra paso a paso cómo se explora y se retrocede.

## Tests

```bash
go test ./backtracking/...
```
