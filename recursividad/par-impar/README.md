# Par e impar

## 1. El problema

Decidir si un número entero no negativo `n` es par o impar **sin usar la
operación módulo**, solo con suma, resta y recursión.

## 2. La estrategia

Dos funciones que se llaman entre sí (recursión mutua):

- `0` es par y no es impar.
- Para cualquier otro `n`, `n` es par si `n-1` es impar, y `n` es impar si
  `n-1` es par.

Cada llamada baja `n` de a uno hasta llegar al caso base `n = 0`.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `esPar(n int) bool` | `true` si `n` es par, `false` en caso contrario. |
| `esImpar(n int) bool` | `true` si `n` es impar, `false` en caso contrario. |

Están en minúsculas porque son internas del ejemplo; el programa principal
imprime algunos casos.

## 4. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./recursividad/par-impar/
```

## 6. Para experimentar

- Probá `esPar` y `esImpar` con números par e impar chicos.
- Investigá el caso `n = 0`: `esPar(0)` es `true` y `esImpar(0)` es `false`.
- Negativos: ¿qué pasa con `esPar(-1)`? La recursión nunca llega al caso base
  `n = 0`, baja para siempre y termina desbordando la pila.

## 7. Tests

```bash
go test ./recursividad/par-impar/
```

## 8. Complejidad

`O(n)`: entre `esPar` y `esImpar` hay una llamada por cada entero entre `n` y
0, así que el trabajo crece linealmente con `n`.
