# Burbujeo (Bubble Sort)

## 1. El problema

Ordenar un slice de enteros de manera ascendente. El ejemplo parte de
`[5 1 4 2 8]` y muestra cada paso del algoritmo hasta dejarlo en
`[1 2 4 5 8]`.

## 2. La estrategia

Recorre repetidamente la lista, compara elementos adyacentes y los
intercambia si están en el orden incorrecto. Las "burbujas" de elementos más
grandes se mueven hacia el final, de ahí el nombre. Optimización: se detiene
si en una pasada completa no hubo intercambios, porque eso significa que el
arreglo ya está ordenado. Es estable e in-place.

## 3. Qué hace el código

| Función | Qué hace |
|---|---|
| `OrdenarBurbujeo(arr []int)` | Ordena el slice en orden ascendente, in-place: modifica el arreglo que recibe y no devuelve nada. |

## 4. Visualización

`burbujeo_visual.go` anima el estado del arreglo paso a paso: en cada paso
dibuja una línea con el arreglo y una barra de `#` por elemento, proporcional
a su valor. Los dos elementos que se acaban de intercambiar quedan marcados
con `*`.

| Función | Qué hace |
|---|---|
| `animarBurbujeo(w io.Writer, arr []int, delay time.Duration)` | Anima el ordenamiento escribiendo cada estado en `w`, con una pausa de `delay` entre pasos. Trabaja sobre una copia: no modifica `arr`. |

## 5. Cómo correrlo

```bash
go run ./ordenamientos/burbujeo/
```

## 6. Para experimentar

- Cambiá el arreglo en `main.go`: con valores repetidos, con un arreglo ya
  ordenado (ahí se nota el corte por falta de intercambios) o con uno en
  orden inverso (el peor caso).
- Bajá el `delay` a 0 para ver la animación de una sola pasada; subilo si
  querés seguir cada comparación con calma.

## 7. Tests

```bash
go test ./ordenamientos/burbujeo/
```

Cubren el arreglo vacío, un solo elemento, arreglos desordenados y con
valores repetidos.

## 8. Complejidad

`O(n²)` en el peor caso (y en el promedio), con `n` la cantidad de elementos.
Gracias al corte cuando una pasada no produce intercambios, pasa a `O(n)` en
el mejor caso (arreglo ya ordenado). La memoria es `O(1)` por ser in-place.
