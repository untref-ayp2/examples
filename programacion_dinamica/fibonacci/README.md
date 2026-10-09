# Fibonacci

## 1. El problema

Calcular el n-ésimo número de la sucesión de Fibonacci: `F(0) = 0`, `F(1) = 1`
y `F(n) = F(n-1) + F(n-2)`. El árbol de llamadas de la versión recursiva
naïve crece exponencialmente; las versiones con programación dinámica lo
resuelven en tiempo lineal.

## 2. Las variantes

| Variante | Carpeta | Enfoque |
|---|---|---|
| Recursivo | `recursivo/` | Fuerza bruta: `fib(n-1) + fib(n-2)` sin guardar nada. |
| Top-down | `top-down/` | Recursivo con memoización en un mapa. |
| Bottom-up | `bottom-up/` | Tabla `dp[0..n]` llenada de 0 hacia arriba. |
| Bottom-up optimizado | `bottom-up-optimizado/` | Solo las dos últimas posiciones: O(1) de espacio. |
| Tab + Memo | `tab-memo/` | `FibonacciTab` y `FibonacciMemo`, las dos funciones del apunte, juntas. |

## 3. Comparación

Todas las variantes de `recursivo/`, `top-down/`, `bottom-up/` y
`bottom-up-optimizado/` exponen `Fibonacci(n int) int`; `tab-memo/` expone
`FibonacciTab(n int) int` y `FibonacciMemo(n int) int`.

| Variante | Tiempo | Espacio |
|---|---|---|
| Recursivo | O(2^n) | O(n) (pila) |
| Top-down | O(n) | O(n) (mapa) |
| Bottom-up | O(n) | O(n) (tabla) |
| Bottom-up optimizado | O(n) | O(1) |
| Tab + Memo | O(n) | O(n) |

## 4. Visualización

En la raíz de `fibonacci/` hay una capa visual (`fibonacci_visual.go` +
`main.go`) que dibuja el **árbol de llamadas** de `Fibonacci(n)` y la **tabla
bottom-up** `dp[i] = Fibonacci(i)` fila por fila.

```bash
go run ./programacion_dinamica/fibonacci/
```

## 5. Cómo correr cada variante

```bash
go run ./programacion_dinamica/fibonacci/recursivo/
go run ./programacion_dinamica/fibonacci/top-down/
go run ./programacion_dinamica/fibonacci/bottom-up/
go run ./programacion_dinamica/fibonacci/bottom-up-optimizado/
go run ./programacion_dinamica/fibonacci/tab-memo/
```

## 6. Para experimentar

- Subí `n` entre 30 y 45 en `main.go` y mirá los tiempos: la recursiva se
  dispara, las demás siguen instantáneas.
- Corré la visual con `n` chico (5 o 6) para ver cómo el árbol naive explota
  en llamadas repetidas.

## 7. Tests

```bash
go test ./programacion_dinamica/fibonacci/...
```
