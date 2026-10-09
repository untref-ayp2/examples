# Ladrón: variante top-down

## 1. El problema

Una calle con `houses` casas, cada una con un botín. No se pueden asaltar dos
casas contiguas. Maximizar el botín total.

## 2. La estrategia

Como la variante recursiva, pero con **memoización**: un mapa guarda el máximo
botín desde cada posición `current`. Se resuelve «de arriba hacia abajo» y cada
subproblema se calcula una sola vez.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `Rob(houses []int) int` | Botín máximo sin robar dos casas contiguas. |

## 4. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./programacion_dinamica/ladron/top-down/
```

## 6. Para experimentar

- Probá calles vacías y de una sola casa: devuelve 0 y el valor de la casa.
- Agregá casas en `main.go`: sigue siendo instantáneo (O(n)).

## 7. Tests

```bash
go test ./programacion_dinamica/ladron/top-down/
```

## 8. Complejidad

O(n) en tiempo y O(n) en espacio por el mapa de memoización más la pila de
recursión.
