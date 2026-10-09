# Cambio: variante recursiva

## 1. El problema

Dado un monto y denominaciones de monedas, encontrar la **menor cantidad de
monedas** que suman exactamente el monto.

## 2. La estrategia

Fuerza bruta recursiva: por cada moneda que entra en el monto, se resta y se
vuelve a llamar. El caso base es `monto == 0` (0 monedas). El mismo subproblema
se calcula muchísimas veces: ese es el desperdicio que arreglan las variantes
con programación dinámica.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `Cambio(monto int, coins []int) int` | Mínima cantidad de monedas para `monto`. |

## 4. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./programacion_dinamica/cambio/recursivo/
```

## 6. Para experimentar

- Subí el monto en `main.go`: los tiempos explotan (exponencial).
- Agregá una moneda de 0: la recursión nunca termina.

## 7. Tests

```bash
go test ./programacion_dinamica/cambio/recursivo/
```

## 8. Complejidad

Exponencial en tiempo: en cada paso se prueban todas las monedas y el árbol de
llamadas crece sin control. El espacio es O(m) por la pila de recursión.