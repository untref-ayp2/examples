# Algoritmos ávidos

En esta carpeta se encuentran los ejemplos de algoritmos ávidos (*greedy*) que
se ven en clase y se resuelven en el aula. Cada ejemplo vive en su propia carpeta
y se puede ejecutar y testear de forma independiente.

## Ejemplos

| Ejemplo | Descripción | Cómo correrlo |
|---|---|---|
| [actividades](actividades/) | Selección de la mayor cantidad de actividades compatibles, con una línea de tiempo de las elegidas. | `go run ./avidos/actividades/` |
| [cambio](cambio/) | Cambio de moneda con la menor cantidad de monedas posibles. | `go run ./avidos/cambio/` |
| [combustible](combustible/) | Mínima cantidad de recargas para llegar a destino, con el recorrido de las estaciones. | `go run ./avidos/combustible/` |

## Tests

```bash
go test ./avidos/...
```

Cada ejemplo tiene su propio `README.md` con el detalle del algoritmo, los
parámetros para experimentar y la complejidad.
