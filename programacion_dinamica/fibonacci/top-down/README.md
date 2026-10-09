# Fibonacci: variante top-down

## 1. El problema

Calcular el n-ésimo número de Fibonacci: `F(0) = 0`, `F(1) = 1` y
`F(n) = F(n-1) + F(n-2)`.

## 2. La estrategia

Como la variante recursiva, pero con **memoización**: un mapa guarda cada
`fib(k)` ya calculado. Se resuelve «de arriba hacia abajo»: desde `n` hacia los
casos base, y cada subproblema se calcula una sola vez.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `Fibonacci(n int) int` | El n-ésimo número de Fibonacci. |

## 4. Visualización

El árbol de llamadas (ahora sin recalcular) se puede ver con la capa visual en
la raíz de `fibonacci/`:

```bash
go run ./programacion_dinamica/fibonacci/
```

## 5. Cómo correrlo

```bash
go run ./programacion_dinamica/fibonacci/top-down/
```

## 6. Para experimentar

- Subí `n` en `main.go` (probá 30, 40, 45): siempre instantáneo (O(n)).

## 7. Tests

```bash
go test ./programacion_dinamica/fibonacci/top-down/
```

## 8. Complejidad

O(n) en tiempo (cada `fib(k)` se calcula una vez) y O(n) en espacio por el mapa
de memoización más la pila de recursión.