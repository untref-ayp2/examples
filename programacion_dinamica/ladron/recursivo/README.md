# Ladrón: variante recursiva

## 1. El problema

Una calle con `houses` casas, cada una con un botín. No se pueden asaltar dos
casas contiguas. Maximizar el botín total.

## 2. La estrategia

Decisión binaria por casa: se prueba **robar la casa actual y saltar a la
siguiente no contigua** (`current + 2`) o **no robarla** (`current + 1`). Se
toma el máximo. Sin guardar nada, el árbol de decisiones duplica trabajo y
crece exponencialmente.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `Rob(houses []int) int` | Botín máximo sin robar dos casas contiguas. |

## 4. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./programacion_dinamica/ladron/recursivo/
```

## 6. Para experimentar

- Probá calles vacías y de una sola casa: devuelve 0 y el valor de la casa.
- Agregá casas en `main.go`: los tiempos explotan (exponencial).

## 7. Tests

```bash
go test ./programacion_dinamica/ladron/recursivo/
```

## 8. Complejidad

O(2^n) en tiempo por el árbol de decisiones duplicado; O(n) de espacio por la
pila de recursión.
