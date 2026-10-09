# Quicksort (Ordenamiento rápido)

## 1. El problema

Ordenar un slice de enteros de manera ascendente. El ejemplo parte de
`[5 9 2 8 1 7 4 3 10 6]` y muestra cada paso del algoritmo hasta dejarlo en
`[1 2 3 4 5 6 7 8 9 10]`.

## 2. La estrategia

Divide y conquista con partición: elige un pivote (en la implementación del
ejemplo, el primer elemento), particiona el segmento en cuestión dejando al
pivote en su posición final — los menores a la izquierda y los mayores a la
derecha — y repite recursivamente con las dos sublistas. Es el algoritmo de
ordenamiento por comparación más usado en la práctica por su buen promedio y
por ordenar in-place.

## 3. Qué hace el código

| Función | Qué hace |
|---|---|
| `Quicksort(array []int) []int` | Ordena el slice en orden ascendente. Los intercambios suceden sobre el mismo slice (in-place) y se devuelve ese mismo slice ya ordenado. |

## 4. Visualización

`quicksort_visual.go` anima el estado del arreglo paso a paso: en cada paso
dibuja una línea con el arreglo y una barra de `#` por elemento, proporcional
a su valor. La posición donde queda el pivote (su posición final) se marca
con `*`.

| Función | Qué hace |
|---|---|
| `animarQuicksort(w io.Writer, array []int, delay time.Duration)` | Anima el ordenamiento escribiendo cada estado en `w`, con una pausa de `delay` entre pasos. Trabaja sobre una copia: no modifica `array`. |

## 5. Cómo correrlo

```bash
go run ./ordenamientos/quicksort/
```

## 6. Para experimentar

- Cambiá el arreglo en `main.go`: se ve cómo cada partición deja un elemento
  en su posición final y el problema se achica hacia los costados.
- Probá con un arreglo ya ordenado o en orden inverso: ahí el pivote (el
  primero) queda siempre en un extremo y el comportamiento decae a `O(n²)`.
- Bajá el `delay` a 0 para ver la animación de una sola pasada.

## 7. Tests

```bash
go test ./ordenamientos/quicksort/
```

Cubren el arreglo vacío, un solo elemento y arreglos con cantidad par e impar
de elementos.

## 8. Complejidad

`O(n log n)` promedio, con `n` la cantidad de elementos; `O(n²)` en el peor
caso, cuando el pivote elegido deja siempre particiones desparejas (por
ejemplo, arreglos ya ordenados con el primer elemento como pivote). La
memoria es `O(log n)` promedio por la profundidad de la recursión (in-place,
sin arreglos auxiliares).