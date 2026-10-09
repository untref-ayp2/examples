# Ordenamientos

En esta carpeta se encuentran los ejemplos de algoritmos de ordenamiento que se
ven en clase y se resuelven en el aula. Cada ejemplo vive en su propia carpeta y
se puede ejecutar y testear de forma independiente.

## Ejemplos

| Ejemplo | Descripción | Cómo correrlo |
|---|---|---|
| [burbujeo](burbujeo/) | Ordenamiento por burbujeo, con el estado del arreglo paso a paso. | `go run ./ordenamientos/burbujeo/` |
| [insercion](insercion/) | Ordenamiento por inserción, con el estado del arreglo paso a paso. | `go run ./ordenamientos/insercion/` |
| [seleccion](seleccion/) | Ordenamiento por selección, con el estado del arreglo paso a paso. | `go run ./ordenamientos/seleccion/` |
| [mergesort](mergesort/) | Ordenamiento por mezcla (recursivo). | `go run ./ordenamientos/mergesort/` |
| [quicksort](quicksort/) | Ordenamiento rápido (recursivo). | `go run ./ordenamientos/quicksort/` |

## Tests

```bash
go test ./ordenamientos/...
```

Cada ejemplo tiene su propio `README.md` con el detalle del algoritmo, los
parámetros para experimentar y la complejidad.
