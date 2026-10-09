# Ladrón: variante bottom-up optimizada

## 1. El problema

Una calle con `houses` casas, cada una con un botín. No se pueden asaltar dos
casas contiguas. Maximizar el botín total.

## 2. La estrategia

Para el máximo hasta la posición `i` solo hacen falta los dos valores
anteriores. En vez de guardar toda la tabla, dos variables (`prevPrev` y
`prev`) se van desplazando en cada paso. Mismo razonamiento bottom-up, O(1) de
espacio.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `Rob(houses []int) int` | Botín máximo sin robar dos casas contiguas. |

## 4. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./programacion_dinamica/ladron/bottom-up-optimizado/
```

## 6. Para experimentar

- Cambiá los valores o agregá casas en `main.go`: O(n) de tiempo y espacio
  constante.
- Ojo: la versión asume al menos **dos casas**; las calles vacías o de una sola
  casa están cubiertas por `recursivo/` y `top-down/`.

## 7. Tests

```bash
go test ./programacion_dinamica/ladron/bottom-up-optimizado/
```

## 8. Complejidad

O(n) en tiempo y O(1) en espacio: solo se conservan las dos últimas
posiciones.
