# Mergesort (Ordenamiento por mezcla)

## 1. El problema

Ordenar un slice de enteros de manera ascendente. El ejemplo parte de
`[5 9 2 8 1 7 4 3 10 6]` y muestra cada paso del algoritmo hasta dejarlo en
`[1 2 3 4 5 6 7 8 9 10]`.

## 2. La estrategia

Divide y conquista: divide el arreglo en dos mitades, ordena cada mitad
recursivamente y fusiona (intercala) las dos mitades ya ordenadas en una sola
lista ordenada. El caso base es un arreglo con menos de dos elementos, que ya
está ordenado por definición. Como la fusión de dos listas ordenadas cuesta
lineal, el costo total queda `O(n log n)`.

## 3. Qué hace el código

| Función | Qué hace |
|---|---|
| `Mergesort(array []int) []int` | Devuelve un nuevo slice con los elementos de `array` ordenados. No modifica el arreglo de entrada. |
| `merge(left, right []int) []int` | Intercala dos slices ya ordenados en uno solo ordenado (función auxiliar usada por `Mergesort`). |

## 4. Visualización

`mergesort_visual.go` anima el estado del arreglo paso a paso: en cada paso
dibuja una línea con el arreglo y una barra de `#` por elemento, proporcional
a su valor. El segmento que se acaba de fusionar queda marcado con `*`.

| Función | Qué hace |
|---|---|
| `animarMergesort(w io.Writer, array []int, delay time.Duration)` | Anima el ordenamiento escribiendo cada estado en `w`, con una pausa de `delay` entre pasos. Trabaja sobre una copia: no modifica `array`. |

## 5. Cómo correrlo

```bash
go run ./ordenamientos/mergesort/
```

## 6. Para experimentar

- Cambiá el arreglo en `main.go`: se ve cómo los segmentos crecen y se van
  fusionando hasta cubrir el arreglo completo.
- Probá con un arreglo de otro largo (impar, por ejemplo) para ver cómo
  quedan las mitades desparejas.
- Bajá el `delay` a 0 para ver la animación de una sola pasada.

## 7. Tests

```bash
go test ./ordenamientos/mergesort/
```

Cubren el arreglo vacío, un solo elemento y arreglos con cantidad par e impar
de elementos.

## 8. Complejidad

`O(n log n)` en todos los casos (mejor, promedio y peor), con `n` la cantidad
de elementos: dividir cuesta `O(log n)` niveles y cada nivel fusiona `n`
elementos. La memoria es `O(n)` por los arreglos auxiliares que se crean en
las fusiones.