# Selección de actividades

## 1. El problema

Tenés un conjunto de actividades que se desarrollan en intervalos de tiempo
(por ejemplo, charlas en una jornada). Cada actividad tiene un nombre, una hora
de inicio y una hora de fin. Las actividades **se solapan** cuando sus
intervalos se superponen, y dos actividades solapadas no se pueden hacer ambas.
El objetivo es elegir la mayor cantidad posible de actividades compatibles
(que no se solapen).

## 2. La estrategia

Algoritmo ávido: se recorre la lista de actividades **ordenada por tiempo de
finalización** y se va tomando la primera actividad compatible con la última
elegida. Una actividad es compatible si su inicio es mayor o igual que el fin
de la última elegida.

Elegir primero la que termina más temprano deja el mayor margen posible para
las siguientes, y se puede demostrar que esta regla simple da el óptimo: la
selección greedy es la de mayor cardinalidad.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `SelectorActividadesIterativo(actividades []Actividad) []Actividad` | La mayor cantidad de actividades compatibles, en orden. |

`Actividad` es un `struct` con tres campos: `Nombre string`, `Inicio int` y
`Fin int`. La función supone que `actividades` llega ordenada de menor a mayor
por tiempo de finalización (es la precondición del algoritmo) y no modifica el
slice original.

## 4. Visualización

La capa `actividades_visual.go` dibuja una **línea de tiempo** en ASCII: cada
actividad elegida es una fila y su barra `[====]` ocupa las unidades de tiempo
en las que transcurre, con el eje de tiempos arriba.

| Función | Qué hace |
|---|---|
| `dibujarLineaTiempo(w io.Writer, actividades []Actividad, delay time.Duration)` | Escribe la línea de tiempo en `w` y espera `delay` (no imprime por su cuenta). |

## 5. Cómo correrlo

```bash
go run ./avidos/actividades/
```

## 6. Para experimentar

- Agregá o quitá actividades en `main.go`. Recordá mantenerlas **ordenadas por
  finalización**, que es la precondición del algoritmo.
- Cambiá los horarios de una actividad para que se superponga con otra y fijate
  cómo cambia la selección.
- Probá con actividades "encadenadas" (una termina justo cuando arranca la
  otra): son compatibles y entran todas.
- Bajá el `delay` a 0 para ver el resultado de una sola pasada.

## 7. Tests

```bash
go test ./avidos/actividades/
```

## 8. Complejidad

`O(n log n)`: el costo lo domina ordenar las actividades por tiempo de
finalización; la selección en sí es una sola pasada, `O(n)`, porque cada
actividad se visita una única vez.