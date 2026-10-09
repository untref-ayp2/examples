# Recursividad

En esta carpeta se encuentran los ejemplos de recursividad que se ven en clase
y se resuelven en el aula. Cada ejemplo vive en su propia carpeta y se puede
ejecutar y testear de forma independiente.

## Ejemplos

| Ejemplo | Descripción | Cómo correrlo |
|---|---|---|
| [bases](bases/) | Descomposición de un número en base 10 o en una base arbitraria. | `go run ./recursividad/bases/` |
| [factorial](factorial/) | Cálculo del factorial de `n` de forma recursiva. | `go run ./recursividad/factorial/` |
| [hanoi](hanoi/) | Torres de Hanoi: mueve los discos paso a paso, con visualización ASCII. | `go run ./recursividad/hanoi/` |
| [palindromo](palindromo/) | Detección de palíndromos en textos. | `go run ./recursividad/palindromo/` |
| [par-impar](par-impar/) | Determinación de paridad mediante recursión. | `go run ./recursividad/par-impar/` |

## Tests

```bash
go test ./recursividad/...
```

Cada ejemplo tiene su propio `README.md` con el detalle del algoritmo, los
parámetros para experimentar y la complejidad.
