# Backtracking

En esta carpeta se encuentran los ejemplos de backtracking que se ven en clase y
se resuelven en el aula. Cada ejemplo vive en su propia carpeta y se puede
ejecutar y testear de forma independiente.

## Ejemplos

Cada ejemplo tiene dos versiones: una **animada**, que muestra la búsqueda paso
a paso en la terminal, y una **simple**, sin animación y sin callbacks, pensada
para estudiar el algoritmo y su esquema (caso base, poda, extender, registrar y
borrar).

| Ejemplo | Versión animada | Versión simple | Descripción |
|---|---|---|---|
| Cambio de moneda | [cambio](cambio/) | [cambio-simple](cambio-simple/) | Explora las combinaciones de denominaciones. |
| Laberinto | [ratmaze](ratmaze/) | [ratmaze-simple](ratmaze-simple/) | Encuentra un camino desde la entrada hasta la salida. |
| N reinas | [reinas](reinas/) | [reinas-simple](reinas-simple/) | Ubica n reinas sin que se ataquen. |
| Sudoku | [sudoku](sudoku/) | [sudoku-simple](sudoku-simple/) | Resuelve un tablero de 9×9. |

Las versiones animadas (`cambio`, `ratmaze`, `reinas` y `sudoku`) incluyen una
capa visual ASCII que muestra paso a paso cómo se explora y se retrocede; las
simples imprimen el resultado de una sola vez.

## Tests

```bash
go test ./backtracking/...
```
