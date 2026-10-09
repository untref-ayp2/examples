# Cambio: variante top-down

## 1. El problema

Dado un monto y denominaciones de monedas, encontrar la **menor cantidad de
monedas** que suman exactamente el monto.

## 2. La estrategia

Como la variante recursiva, pero con **memoización**: un mapa guarda el
resultado de cada monto ya resuelto, así un subproblema repetido se recupera en
O(1) en lugar de recalcularse. Se resuelve «de arriba hacia abajo»: desde el
monto pedido hacia los más chicos.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `Cambio(monto int, coins []int) int` | Mínima cantidad de monedas para `monto`. |

## 4. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./programacion_dinamica/cambio/top-down/
```

## 6. Para experimentar

- Subí el monto en `main.go`: sigue siendo instantáneo (O(m·n)).
- Agregá una moneda de 0: la recursión nunca termina (el monto no baja).

## 7. Tests

```bash
go test ./programacion_dinamica/cambio/top-down/
```

## 8. Complejidad

O(m·n) en tiempo (m = monto, n = cantidad de monedas) y O(m) en espacio por el
mapa de memoización más la pila de recursión.
