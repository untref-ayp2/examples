# Cambio: variante bottom-up

## 1. El problema

Dado un monto y denominaciones de monedas, encontrar la **menor cantidad de
monedas** que suman exactamente el monto.

## 2. La estrategia

**Tabulación**: se llena una tabla `dp[x]` = mínimas monedas para el monto `x`,
desde `0` hasta el monto pedido. Para cada `x` se prueba cada moneda y se toma
`1 + dp[x - coin]` mínimo. La respuesta es `dp[monto]`.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `Cambio(monto int, coins []int) int` | Mínima cantidad de monedas para `monto`. |

## 4. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./programacion_dinamica/cambio/bottom-up/
```

## 6. Para experimentar

- Subí el monto en `main.go`: sigue siendo O(m·n).
- Con `coins = {1, 5, 11}` y `monto = 97`, el ávido usaría 13 monedas; la DP
  encuentra el óptimo, 11.

## 7. Tests

```bash
go test ./programacion_dinamica/cambio/bottom-up/
```

## 8. Complejidad

O(m·n) en tiempo (m = monto, n = cantidad de monedas) y O(m) en espacio por la
tabla.