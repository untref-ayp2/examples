# Cambio de moneda

## 1. El problema

Dado un monto `monto` y un conjunto de denominaciones `coins`, encontrar la
**menor cantidad de monedas** que suman exactamente el monto. Las monedas se
pueden usar todas las veces que sea necesario.

## 2. Las variantes

El problema se resuelve de tres formas, una por carpeta:

| Variante | Carpeta | Enfoque |
|---|---|---|
| Recursivo | `recursivo/` | Prueba todas las combinaciones por fuerza bruta. |
| Top-down | `top-down/` | Igual que la recursiva, pero guardando cada resultado en un mapa (memoización). |
| Bottom-up | `bottom-up/` | Llena una tabla `dp[x]` = mínimas monedas para el monto `x`, de 0 a `monto`. |

## 3. Comparación

Las tres exponen la misma función `Cambio(monto int, coins []int) int`.

| Variante | Tiempo | Espacio | Observaciones |
|---|---|---|---|
| Recursivo | Exponencial | O(m) (pila) | Se vuelve impracticable con montos grandes. |
| Top-down | O(m·n) | O(m) | Solo calcula los subproblemas que necesita. |
| Bottom-up | O(m·n) | O(m) | Llena toda la tabla, use o no cada celda. |

Ojo con una moneda de **valor 0**: en las tres variantes el monto nunca baja y
entran en un ciclo infinito. Es un buen experimento para ver por qué la
definición de los subproblemas importa.

## 4. Cómo correrlo

```bash
go run ./programacion_dinamica/cambio/recursivo/
go run ./programacion_dinamica/cambio/top-down/
go run ./programacion_dinamica/cambio/bottom-up/
```

## 5. Para experimentar

- Subí el monto en `main.go` y compará cuánto tarda cada variante.
- Con `coins = {1, 5, 11}` y `monto = 97`: el ávido que siempre toma la de 11
  usa 13 monedas; la DP encuentra el óptimo, 11 (siete de 11 más cuatro de 5).
- Agregá una moneda de 0 y mirá cómo se cuelgan las tres.

## 6. Tests

```bash
go test ./programacion_dinamica/cambio/...
```
