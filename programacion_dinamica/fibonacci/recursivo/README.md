# Fibonacci: variante recursiva

## 1. El problema

Calcular el n-ésimo número de Fibonacci: `F(0) = 0`, `F(1) = 1` y
`F(n) = F(n-1) + F(n-2)`.

## 2. La estrategia

Traducción directa de la recurrencia, sin guardar nada. El árbol de llamadas
crece exponencialmente y el mismo `fib(k)` se recalcula miles de veces: es el
caso más claro de por qué hace falta programación dinámica.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `Fibonacci(n int) int` | El n-ésimo número de Fibonacci. |

## 4. Visualización

El árbol de llamadas se puede ver con la capa visual en la raíz de
`fibonacci/`:

```bash
go run ./programacion_dinamica/fibonacci/
```

## 5. Cómo correrlo

```bash
go run ./programacion_dinamica/fibonacci/recursivo/
```

## 6. Para experimentar

- Subí `n` de a poco en `main.go` (30 o más) y mirá cómo explota el tiempo.
- Con `n = 35` ya se nota la espera; con 45 se vuelve impracticable.

## 7. Tests

```bash
go test ./programacion_dinamica/fibonacci/recursivo/
```

## 8. Complejidad

O(2^n) en tiempo por el árbol de llamadas duplicado; O(n) de espacio por la
pila de recursión.
