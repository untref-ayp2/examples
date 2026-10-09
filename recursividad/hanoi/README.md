# Torres de Hanoi

## 1. El problema

Hay tres varas (`A`, `B` y `C`) y `n` discos de tamaños distintos.
Inicialmente todos los discos están apilados en la vara `A`, del más grande
(abajo) al más chico (arriba). Hay que mover toda la pila a la vara `C`
respetando dos reglas: solo se mueve un disco a la vez y nunca puede quedar un
disco grande encima de uno chico.

## 2. La estrategia

Para mover `n` discos de `A` a `C` usando `B` como auxiliar:

1. mover los `n-1` discos de arriba de `A` a `B` (usando `C` como auxiliar);
2. mover el disco más grande de `A` a `C`;
3. mover los `n-1` discos de `B` a `C` (usando `A` como auxiliar).

El caso base es `n = 1`: el disco se mueve directo de origen a destino.

## 3. Qué hace el código

La lógica **no imprime**: cada movimiento se entrega a un callback (patrón
observador, igual que en los ejemplos de backtracking del taller):

| Función | Qué hace |
|---|---|
| `Mover(n int, origen string, aux string, destino string, onMove func(disco int, origen string, destino string))` | Genera la secuencia de movimientos para `n` discos. Por cada movimiento invoca a `onMove` con el disco movido y las varas de origen y destino. |

Separar la generación de la muestra permite acumular los movimientos en una
lista, contarlos o dibujarlos sin tocar la lógica.

## 4. Visualización

La capa `hanoi_visual.go` dibuja las tres varas en un marco fijo (UTF-8,
igual que el ejemplo de sudoku): cada dibujo borra la pantalla y pinta el
estado en el mismo lugar. Los discos son barras de `█`, la vara vacía se ve
como `│`, el disco que acaba de moverse se resalta en cian y un contador
acumula los movimientos. El estado se modela con `estadoHanoi` (cada vara es
una pila de discos) y se pinta con:

| Función | Qué hace |
|---|---|
| `nuevoEstadoHanoi(n int) *estadoHanoi` | Crea el estado inicial con los `n` discos en la vara `A`. |
| `(*estadoHanoi).mover(disco, origen, destino int)` | Pasa el disco superior de `origen` a `destino` y avanza el contador. |
| `(*estadoHanoi).Tablero() string` | Devuelve el estado actual con bordes, disco en movimiento resaltado y contador. |
| `(*estadoHanoi).dibujar(w io.Writer, delay time.Duration)` | Borra la pantalla, escribe `Tablero` en `w` y espera `delay` (no imprime por su cuenta). |

El `main.go` conecta ambas partes: el callback de `Mover` actualiza el estado
y lo redibuja, así se ve la pila acomodarse paso a paso en la misma posición.

## 5. Cómo correrlo

```bash
go run ./recursividad/hanoi/
```

En `main.go` se puede ajustar el número de discos (`n`, probá de 3 a 8) y la
velocidad (`delay`).

## 6. Para experimentar

- Cambiá `n` en `main.go`: con 3 o 4 discos se sigue cómodamente; con 8 ya hay
  255 movimientos.
- Bajá el `delay` a 0 y mirá el contador del dibujo: ¿el total coincide con
  `2^n - 1`?

## 7. Tests

```bash
go test ./recursividad/hanoi/
```

## 8. Complejidad

`O(2^n)`: cada disco duplica la cantidad de movimientos, que es exactamente
`2^n - 1`. Por eso con valores grandes el programa se vuelve impracticable:
es el costo de un algoritmo exponencial.
