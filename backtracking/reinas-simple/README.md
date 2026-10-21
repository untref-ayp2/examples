# N reinas (versión simple)

## 1. El problema

Ubicar `n` reinas en un tablero de `n × n` de modo que ninguna ataque a otra:
no puede haber dos reinas en la misma fila, la misma columna ni la misma
diagonal.

## 2. La estrategia

Backtracking fila por fila: se ubica una reina en la fila actual probando las
columnas desde la 0. Si la columna no entra en conflicto con las reinas ya
puestas, se registra y se pasa a la fila siguiente. Cuando ninguna columna
sirve, se saca la reina de la fila anterior (**vuelta atrás**) y se sigue
probando.

La solución parcial es un slice: la posición `i` es la fila `i` y el valor es
la columna donde está su reina.

## 3. Qué hace el código

| Función | Qué hace |
|---|---|
| `NReinas(n int) []int` | Devuelve la primera solución, o `nil` si no hay. |
| `NReinasTodas(n int) [][]int` | Devuelve todas las soluciones. |
| `esSolucion(n int, parcial []int) bool` | Caso base: ya hay una reina por fila. |
| `esFactible(fila, columna int, parcial []int) bool` | Indica si la reina de (fila, columna) choca con alguna de las ya ubicadas. |
| `backtracking(n, fila int, parcial []int, soluciones *[][]int, primera bool)` | La recursión del esquema (extender, registrar, borrar). |

Es el mismo esquema de la versión animada en [`../reinas/`](../reinas/), sin el
`ObservadorReinas`: en vez de notificar cada intento, acumula las soluciones.

## 4. Cómo correrlo

```bash
go run ./backtracking/reinas-simple/
```

## 5. Para experimentar

- Cambiá `n` en `main.go`: con 8 hay 92 soluciones, y con 10 (724) y 12
  (14200) se empieza a notar el crecimiento. Como referencia medida con
  `NReinasTodas`: n=8 tarda décimas de milisegundo, n=10 unos pocos
  milisegundos y n=12 alrededor de 75 ms.
- Compará `NReinas` (corta en la primera) con `NReinasTodas` (agota todas).
- Probá `n = 2` o `n = 3`: no hay solución y ambas funciones lo reflejan
  (`nil` y slice vacío).

## 6. Tests

```bash
go test ./backtracking/reinas-simple/
```

## 7. Complejidad

Peor caso `O(n!)`: se prueban las combinaciones de columnas, aunque
`esFactible` poda las ramas inviables. El espacio es `O(n)` por la solución
parcial más lo que ocupen las soluciones guardadas.
