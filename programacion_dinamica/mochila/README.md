# Mochila 0/1

## 1. El problema

Tenés `items` objetos, cada uno con un `Peso` y un `Valor`, y una mochila con
`capacidad` máxima. Cada objeto se toma una sola vez o no se toma (mochila
**0/1**; no se puede fraccionar). El objetivo es elegir los objetos que
maximicen el valor total sin superar la capacidad.

## 2. La estrategia

Programación dinámica con la recurrencia clásica: para cada cantidad de ítems
`i` y cada capacidad `w`, el máximo valor es el de `MochilaTab(n-1, w)` si el
ítem no entra, o el mejor entre tomarlo sin su peso y no tomarlo. Se implementa
con tabulación (bottom-up), memoización (top-down) y una versión que además
reconstruye la selección recorriendo la tabla hacia atrás.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `MochilaTab(items []Item, capacidad int) int` | Valor máximo por tabulación. |
| `MochilaMemo(items []Item, capacidad int) int` | Valor máximo por memoización. |
| `MochilaTabConItems(items []Item, capacidad int) (int, []Item)` | Valor máximo y los ítems elegidos. |

`Item` es un `struct` con `Peso int` y `Valor int`.

## 4. Visualización

La capa `mochila_visual.go` dibuja el **llenado de la tabla DP** fila por fila:
una fila por cada cantidad de ítems considerada y una columna por capacidad.

| Función | Qué hace |
|---|---|
| `dibujarTablaMochila(w io.Writer, items []Item, capacidad int, delay time.Duration)` | Escribe la tabla en `w` y espera `delay` entre filas (no imprime por su cuenta). |

## 5. Cómo correrlo

```bash
go run ./programacion_dinamica/mochila/
```

## 6. Para experimentar

- Probá `capacidad = 0` o sin ítems: el valor máximo es 0.
- Cambiá los pesos y valores en `main.go` y mirá cómo cambia la selección.
- Bajá el `delay` a 0 para ver toda la tabla de una.

## 7. Tests

```bash
go test ./programacion_dinamica/mochila/
```

## 8. Complejidad

O(n·W) en tiempo y espacio (n = cantidad de ítems, W = capacidad), tanto en
tabulación como en memoización.