# Iterator

## 1. El problema

El patrón *Iterator* permite recorrer los elementos de una colección sin
exponer su estructura interna. La interfaz unificada tiene dos métodos:

- `Siguiente() bool`: avanza al siguiente elemento y devuelve `true` si hay
  más. La primera llamada posiciona el iterador en el primer elemento.
- `Valor() int`: devuelve el valor del elemento actual. Solo debe llamarse
  después de que `Siguiente()` haya devuelto `true`.

En esta carpeta hay **dos implementaciones** de la misma interfaz, sobre
estructuras internas muy distintas:

| Subcarpeta | Estructura | Colección | Fábrica |
|---|---|---|---|
| `matriz` | `[][]int` (filas y columnas) | `Matriz` | `Matriz.Iterador() *Iterador` |
| `lista-enlazada` | `Nodo` con puntero al siguiente | `Lista` | `Lista.Iterador() *Iterador` |

## 2. La estrategia

Ambas implementaciones comparten el mismo patrón de uso:

```go
it := coleccion.Iterador()
for it.Siguiente() {
    fmt.Println(it.Valor())
}
```

El iterador guarda la posición actual dentro de la colección y una bandera
`primera` para distinguir la primera llamada a `Siguiente()` (que solo
posiciona) de las siguientes (que avanzan). En la matriz se avanza columna a
columna y, al terminar la fila, se pasa a la siguiente; en la lista se avanza
de nodo en nodo.

Gracias a la interfaz común, el código que recorre una `Matriz` es idéntico
al que recorre una `Lista`: solo cambia la colección que entrega el iterador.

## 3. Qué hace el código

| Firma | Rol |
|---|---|
| `Matriz.Iterador() *Iterador` | Fábrica: devuelve un iterador nuevo posicionado al inicio de la matriz. |
| `Lista.Iterador() *Iterador` | Fábrica: devuelve un iterador nuevo posicionado al inicio de la lista. |
| `Iterador.Siguiente() bool` | Avanza y devuelve `true` si hay elemento; `false` si no quedan más. |
| `Iterador.Valor() int` | Devuelve el valor del elemento actual. |
| `Lista.AgregarAlFinal(int)` | Agrega un valor al final de la lista enlazada. |

## 4. Cómo correrlos

```bash
go run ./patrones/iterator/matriz/
go run ./patrones/iterator/lista-enlazada/
```

## 5. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 6. Para experimentar

- Cambiá los datos de la matriz en `main.go`: probá un rectángulo (filas de
  distinto largo), el caso vacío (`NewMatriz(nil)`) y una matriz de una sola
  fila o columna.
- Cambiá los valores y el orden de `AgregarAlFinal` en la lista y verificá
  que el recorrido sigue el orden de inserción.
- Intercambiá el cuerpo del `for`: en ambos casos es el mismo, porque la
  interfaz es la misma.

## 7. Complejidad

—: el patrón agrega una capa de indirección constante; no cambia el orden de
crecimiento del recorrido de la colección.
