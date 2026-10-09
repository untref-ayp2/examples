# División y conquista

En esta carpeta se encuentran los ejemplos de división y conquista que se han
visto en clase, y que se han resuelto en el aula. Cada ejemplo vive en su propia
carpeta y se puede ejecutar y testear de forma independiente.

## Ejemplos

| Ejemplo | Descripción | Cómo correrlo |
|---|---|---|
| [busqueda-binaria](busqueda-binaria/) | Búsqueda binaria sobre un arreglo ordenado, con estrategia recursiva de división y conquista. | `go run ./division_y_conquista/busqueda-binaria/` |
| [potencia](potencia/) | Cálculo de potencias por división y conquista. | `go run ./division_y_conquista/potencia/` |
| [suma](suma/) | Suma de los elementos de un slice por división y conquista. | `go run ./division_y_conquista/suma/` |

## Tests

```bash
go test ./division_y_conquista/...
```

Cada ejemplo tiene su propio `README.md` con el detalle del algoritmo, los
parámetros para experimentar y la complejidad.
