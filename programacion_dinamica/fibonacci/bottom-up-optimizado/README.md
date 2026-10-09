# Fibonacci: variante bottom-up optimizada

## 1. El problema

Calcular el n-ésimo número de Fibonacci: `F(0) = 0`, `F(1) = 1` y
`F(n) = F(n-1) + F(n-2)`.

## 2. La estrategia

Para calcular `F(n)` solo hacen falta los dos valores anteriores. En vez de
guardar toda la tabla, se mantienen dos variables (`prev2` y `prev`) que se van
desplazando en cada paso. Mismo razonamiento bottom-up, O(1) de espacio.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `Fibonacci(n int) int` | El n-ésimo número de Fibonacci. |

## 4. Visualización

No tiene capa visual propia: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./programacion_dinamica/fibonacci/bottom-up-optimizado/
```

## 6. Para experimentar

- Subí `n` en `main.go` (probá 30, 40, 45): instantáneo y sin crecer el
  espacio.

## 7. Tests

```bash
go test ./programacion_dinamica/fibonacci/bottom-up-optimizado/
```

## 8. Complejidad

O(n) en tiempo y O(1) en espacio: solo se conservan las dos últimas
posiciones.
