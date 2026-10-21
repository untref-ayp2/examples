# Torres de Hanoi (versión simple)

## 1. El problema

Hay tres varas (`A`, `B` y `C`) y `n` discos de tamaños distintos. Al principio
todos los discos están apilados en `A`, del más grande (abajo) al más chico
(arriba). Hay que mover toda la pila a `C` respetando dos reglas: se mueve un
disco a la vez y nunca puede quedar un disco grande encima de uno chico.

## 2. La estrategia

Para mover `n` discos de `A` a `C` usando `B` como auxiliar:

1. mover los `n-1` discos de arriba de `A` a `B` (usando `C` como auxiliar);
2. mover el disco más grande de `A` a `C`;
3. mover los `n-1` discos de `B` a `C` (usando `A` como auxiliar).

El caso base es `n = 1`: el disco se mueve directo de origen a destino.

## 3. Qué hace el código

| Función | Qué hace |
|---|---|
| `Mover(n int, origen, aux, destino string) []Movimiento` | Devuelve la lista de movimientos para pasar `n` discos de `origen` a `destino`. |
| `mover(n int, origen, aux, destino string, movimientos *[]Movimiento)` | La recursión: los tres pasos de la estrategia, acumulando en la lista. |

`Movimiento` es un struct con el disco movido (`Disco`) y las varas de salida
y llegada (`Origen`, `Destino`).

Esta es la versión de estudio: no imprime ni dibuja nada, solo devuelve la
lista. La versión animada, en [`../hanoi/`](../hanoi/), resuelve el mismo
problema entregando cada movimiento a un callback que redibuja las varas.

## 4. Cómo correrlo

```bash
go run ./recursividad/hanoi-simple/
```

Imprime la lista numerada de movimientos.

## 5. Para experimentar

- Cambiá `n` en `main.go`: con 3 discos hay 7 movimientos, con 4 hay 15, con
  10 hay 1023.
- Acumulá los movimientos en un slice y contá cuántas veces se mueve cada
  disco: el más chico se mueve en todos los pasos impares.
- Seguí la secuencia a mano sobre tres pilas de monedas y comprobá que nunca
  queda un disco grande encima de uno chico.

## 6. Tests

```bash
go test ./recursividad/hanoi-simple/
```

Los tests comprueban la secuencia exacta para 3 discos y que un movimiento
simulado respete las reglas del problema.

## 7. Complejidad

`O(2^n)`: cada disco duplica la cantidad de movimientos, que es exactamente
`2^n - 1`. Con valores grandes el programa se vuelve impracticable: es el
costo de un algoritmo exponencial. El espacio también es `O(2^n)` porque la
lista guarda todos los movimientos.
