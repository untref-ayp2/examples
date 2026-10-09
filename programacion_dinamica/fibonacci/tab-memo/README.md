# Fibonacci: tabulación y memoización (tab-memo)

## 1. El problema

Calcular el n-ésimo número de Fibonacci: `F(0) = 0`, `F(1) = 1` y
`F(n) = F(n-1) + F(n-2)`. Acá conviven las dos funciones eficientes del
apunte: una con tabulación y otra con memoización.

## 2. La estrategia

- `FibonacciTab`: **tabulación** (bottom-up). Llena un slice `dp` de 0 a n con
  `dp[i] = dp[i-1] + dp[i-2]`.
- `FibonacciMemo`: **memoización** (top-down). Recursiva con un mapa que guarda
  los valores ya calculados; si un subproblema ya está, se recupera sin
  recalcular.

Ambas devuelven lo mismo; la versión naïve sin caché es exponencial y está en
`recursivo/`.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `FibonacciTab(n int) int` | El n-ésimo número de Fibonacci por tabulación. |
| `FibonacciMemo(n int) int` | El n-ésimo número de Fibonacci por memoización. |

## 4. Visualización

El árbol de llamadas y la tabla se pueden ver con la capa visual en la raíz de
`fibonacci/`:

```bash
go run ./programacion_dinamica/fibonacci/
```

## 5. Cómo correrlo

```bash
go run ./programacion_dinamica/fibonacci/tab-memo/
```

## 6. Para experimentar

- Subí `n` en `main.go` (probá 30, 40, 45): las dos funciones son O(n) e
  instantáneas.

## 7. Tests

```bash
go test ./programacion_dinamica/fibonacci/tab-memo/
```

## 8. Complejidad

O(n) en tiempo y O(n) en espacio en ambas: la tabla en el caso de
`FibonacciTab`, el mapa (y la pila) en el de `FibonacciMemo`.