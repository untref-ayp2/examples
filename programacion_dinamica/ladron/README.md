# El ladrón

## 1. El problema

Un ladrón recorre una calle con `houses` casas, cada una con un botín. No puede
asaltar dos casas **contiguas** (lo descubrirían). El objetivo es elegir el
subconjunto de casas de **máximo botín total** sin que haya dos elegidas
vecinas.

## 2. Las variantes

| Variante | Carpeta | Enfoque |
|---|---|---|
| Recursivo | `recursivo/` | Prueba por fuerza bruta: robar la casa actual y saltar, o no robarla. |
| Top-down | `top-down/` | Igual que la recursiva, guardando resultados en un mapa (memoización). |
| Bottom-up | `bottom-up/` | Tabla `dp[i]` = máximo botín considerando las primeras i+1 casas. |
| Bottom-up optimizado | `bottom-up-optimizado/` | Solo las dos últimas posiciones: O(1) de espacio. |

## 3. Comparación

Todas las variantes exponen `Rob(houses []int) int`.

| Variante | Tiempo | Espacio | Observaciones |
|---|---|---|---|
| Recursivo | O(2^n) | O(n) (pila) | Se vuelve impracticable en calles largas. |
| Top-down | O(n) | O(n) | Solo calcula los subproblemas que necesita. |
| Bottom-up | O(n) | O(n) | Asume al menos dos casas. |
| Bottom-up optimizado | O(n) | O(1) | Asume al menos dos casas. |

Las variantes `bottom-up` asumen calles de al menos dos casas: el caso de una
calle vacía o de una sola casa queda cubierto por `recursivo/` y `top-down/`.

## 4. Cómo correrlo

```bash
go run ./programacion_dinamica/ladron/recursivo/
go run ./programacion_dinamica/ladron/top-down/
go run ./programacion_dinamica/ladron/bottom-up/
go run ./programacion_dinamica/ladron/bottom-up-optimizado/
```

## 5. Para experimentar

- Probá calles vacías y de una sola casa en `recursivo/` y `top-down/`.
- Agregá casas en `main.go` y compará los tiempos: la recursiva explota, las
  demás son lineales.
- Cambiá los valores de una casa y mirá cómo cambia la elección óptima.

## 6. Tests

```bash
go test ./programacion_dinamica/ladron/...
```