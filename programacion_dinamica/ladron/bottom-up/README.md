# Ladrón: variante bottom-up

## 1. El problema

Una calle con `houses` casas, cada una con un botín. No se pueden asaltar dos
casas contiguas. Maximizar el botín total.

## 2. La estrategia

**Tabulación**: `table[i]` guarda el máximo botín considerando las casas hasta
la posición `i`, con `table[i] = max(houses[i] + table[i-2], table[i-1])`. Se
llena de izquierda a derecha y la respuesta es la última celda.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `Rob(houses []int) int` | Botín máximo sin robar dos casas contiguas. |

## 4. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./programacion_dinamica/ladron/bottom-up/
```

## 6. Para experimentar

- Cambiá los valores o agregá casas en `main.go` (O(n)).
- Ojo: la versión asume al menos **dos casas**; las calles vacías o de una sola
  casa están cubiertas por `recursivo/` y `top-down/`.

## 7. Tests

```bash
go test ./programacion_dinamica/ladron/bottom-up/
```

## 8. Complejidad

O(n) en tiempo y O(n) en espacio por la tabla.