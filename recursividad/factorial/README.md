# Factorial

## 1. El problema

Calcular el factorial de un número entero no negativo `n`, definido como
`n! = n × (n-1) × … × 2 × 1`, con la convención `0! = 1`.

## 2. La estrategia

El factorial tiene una definición recursiva natural:

```
0! = 1
n! = n × (n-1)!
```

El caso base (`n = 0`) corta la recursión, y el paso recursivo reduce el
problema a uno más chico: calcular `(n-1)!`.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `Factorial(n int) int` | El factorial de `n` para `n >= 0`. |

Devuelve un `int`; el programa principal itera de 0 a 20 y muestra cada valor.

## 4. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./recursividad/factorial/
```

## 6. Para experimentar

- Corré el programa hasta `n = 20` y mirá cómo crecen los números.
- Probá un `n` más grande (por ejemplo 21): el resultado ya no entra en un
  `int` de 64 bits y se desborda silenciosamente. Por eso el ejemplo se limita
  a `n <= 20`.

## 7. Tests

```bash
go test ./recursividad/factorial/
```

## 8. Complejidad

`O(n)`: hay una llamada recursiva por cada entero entre `n` y 0, así que la
cantidad de llamadas crece linealmente con `n`.
