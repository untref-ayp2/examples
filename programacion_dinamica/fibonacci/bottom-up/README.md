# Fibonacci: variante bottom-up

## 1. El problema

Calcular el n-ésimo número de Fibonacci: `F(0) = 0`, `F(1) = 1` y
`F(n) = F(n-1) + F(n-2)`.

## 2. La estrategia

**Tabulación**: se llena una tabla `dp[0..n]` de abajo hacia arriba, con
`dp[i] = dp[i-1] + dp[i-2]`. La respuesta es `dp[n]`. Cada celda se usa una
sola vez y no hay recursión.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `Fibonacci(n int) int` | El n-ésimo número de Fibonacci. |

## 4. Visualización

El llenado de la tabla se puede ver con la capa visual en la raíz de
`fibonacci/`:

```bash
go run ./programacion_dinamica/fibonacci/
```

## 5. Cómo correrlo

```bash
go run ./programacion_dinamica/fibonacci/bottom-up/
```

## 6. Para experimentar

- Subí `n` en `main.go` (probá 30, 40, 45): siempre instantáneo (O(n)).

## 7. Tests

```bash
go test ./programacion_dinamica/fibonacci/bottom-up/
```

## 8. Complejidad

O(n) en tiempo y O(n) en espacio por la tabla. Si el espacio importa, la
variante `bottom-up-optimizado/` baja el espacio a O(1).