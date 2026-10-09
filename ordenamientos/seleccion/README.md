# Selección (Selection Sort)

## 1. El problema

Ordenar un slice de enteros de manera ascendente. El ejemplo parte de
`[64 25 12 22 11]` y muestra cada paso del algoritmo hasta dejarlo en
`[11 12 22 25 64]`.

## 2. La estrategia

Encuentra repetidamente el elemento mínimo del resto no ordenado y lo coloca
al principio de la lista: en cada paso el mínimo de la parte todavía
desordenada pasa a la primera posición libre. Es simple de entender, pero
siempre hace la misma cantidad de comparaciones. Es in-place y **no** es
estable (puede alterar el orden relativo de elementos iguales).

## 3. Qué hace el código

| Función | Qué hace |
|---|---|
| `OrdenarSeleccion(arr []int)` | Ordena el slice en orden ascendente, in-place: modifica el arreglo que recibe y no devuelve nada. |

## 4. Visualización

`seleccion_visual.go` anima el estado del arreglo paso a paso: en cada paso
dibuja una línea con el arreglo y una barra de `#` por elemento, proporcional
a su valor. El mínimo recién colocado y la posición donde estaba quedan
marcados con `*`.

| Función | Qué hace |
|---|---|
| `animarSeleccion(w io.Writer, arr []int, delay time.Duration)` | Anima el ordenamiento escribiendo cada estado en `w`, con una pausa de `delay` entre pasos. Trabaja sobre una copia: no modifica `arr`. |

## 5. Cómo correrlo

```bash
go run ./ordenamientos/seleccion/
```

## 6. Para experimentar

- Cambiá el arreglo en `main.go`: fijate que la cantidad de pasos no cambia
  aunque el arreglo ya esté ordenado (siempre busca el mínimo del resto).
- Bajá el `delay` a 0 para ver la animación de una sola pasada; subilo si
  querés seguir cada selección con calma.

## 7. Tests

```bash
go test ./ordenamientos/seleccion/
```

Cubren el arreglo vacío, un solo elemento, arreglos desordenados y con
valores repetidos.

## 8. Complejidad

`O(n²)` en todos los casos (mejor, promedio y peor), con `n` la cantidad de
elementos: siempre recorre el resto no ordenado para encontrar el mínimo. La
memoria es `O(1)` por ser in-place.
