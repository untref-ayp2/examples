# Inserción (Insertion Sort)

## 1. El problema

Ordenar un slice de enteros de manera ascendente. El ejemplo parte de
`[5 2 4 6 1 3]` y muestra cada paso del algoritmo hasta dejarlo en
`[1 2 3 4 5 6]`.

## 2. La estrategia

Construye la lista final ordenada un elemento a la vez, insertando cada nuevo
elemento en su posición correcta entre los ya ordenados. Es similar a ordenar
cartas en la mano: se toma la carta siguiente y se la corre hacia atrás hasta
ubicarla donde corresponde. Es estable, in-place y online (procesa un
elemento por vez, sin necesitar toda la entrada).

## 3. Qué hace el código

| Función | Qué hace |
|---|---|
| `OrdenarInsercion(arr []int)` | Ordena el slice en orden ascendente, in-place: modifica el arreglo que recibe y no devuelve nada. |

## 4. Visualización

`insercion_visual.go` anima el estado del arreglo paso a paso: en cada paso
dibuja una línea con el arreglo y una barra de `#` por elemento, proporcional
a su valor. Los elementos que se acaban de desplazar e insertar quedan
marcados con `*`.

| Función | Qué hace |
|---|---|
| `animarInsercion(w io.Writer, arr []int, delay time.Duration)` | Anima el ordenamiento escribiendo cada estado en `w`, con una pausa de `delay` entre pasos. Trabaja sobre una copia: no modifica `arr`. |

## 5. Cómo correrlo

```bash
go run ./ordenamientos/insercion/
```

## 6. Para experimentar

- Cambiá el arreglo en `main.go`: con valores repetidos, con un arreglo casi
  ordenado (ahí el algoritmo hace muy pocos desplazamientos) o con uno en
  orden inverso (el peor caso).
- Bajá el `delay` a 0 para ver la animación de una sola pasada; subilo si
  querés seguir cada inserción con calma.

## 7. Tests

```bash
go test ./ordenamientos/insercion/
```

Cubren el arreglo vacío, un solo elemento, arreglos desordenados y con
valores repetidos.

## 8. Complejidad

`O(n²)` en el peor caso (arreglo en orden inverso) y en el promedio, con `n`
la cantidad de elementos. En el mejor caso (arreglo ya ordenado) cada
inserción no desplaza nada y el costo es `O(n)`. La memoria es `O(1)` por ser
in-place.