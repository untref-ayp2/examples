# Programación dinámica

En esta carpeta se encuentran los ejemplos de programación dinámica que se ven
en clase y se resuelven en el aula. Cada ejemplo vive en su propia carpeta y se
puede ejecutar y testear de forma independiente.

La mayoría de los ejemplos presenta **una carpeta por variante**: la solución
recursiva ingenua, la de *top-down* con memoización, la de *bottom-up* con
tabulación y —cuando corresponde— una versión *bottom-up optimizada* que reduce
el espacio auxiliar. Comparar las variantes es el objetivo del tema.

## Ejemplos

| Ejemplo | Descripción | Cómo correrlo |
|---|---|---|
| [cambio](cambio/) | Cambio de moneda con la menor cantidad de monedas. Variantes: recursivo, top-down y bottom-up. | `go run ./programacion_dinamica/cambio/recursivo/`<br>`go run ./programacion_dinamica/cambio/top-down/`<br>`go run ./programacion_dinamica/cambio/bottom-up/` |
| [fibonacci](fibonacci/) | Cálculo del n-ésimo número de Fibonacci. Variantes: recursivo, top-down, bottom-up y bottom-up optimizado. | `go run ./programacion_dinamica/fibonacci/recursivo/`<br>`go run ./programacion_dinamica/fibonacci/top-down/`<br>`go run ./programacion_dinamica/fibonacci/bottom-up/`<br>`go run ./programacion_dinamica/fibonacci/bottom-up-optimizado/` |
| [ladron](ladron/) | Máximo botín que se puede robar de una calle sin asaltar dos casas contiguas. Variantes: recursivo, top-down, bottom-up y bottom-up optimizado. | `go run ./programacion_dinamica/ladron/recursivo/`<br>`go run ./programacion_dinamica/ladron/top-down/`<br>`go run ./programacion_dinamica/ladron/bottom-up/`<br>`go run ./programacion_dinamica/ladron/bottom-up-optimizado/` |
| [mochila](mochila/) | Problema de la mochila 0/1, con visualización del llenado de la tabla DP. | `go run ./programacion_dinamica/mochila/` |

## Tests

```bash
go test ./programacion_dinamica/...
```
