# Cambio de moneda

## 1. El problema

Dada una cantidad de dinero y un conjunto de denominaciones (billetes y
monedas), hay que devolver el cambio para esa cantidad usando la **menor
cantidad de piezas** posible. En este ejemplo las denominaciones son las del
peso argentino: `1, 2, 5, 10, 20, 50, 100, 200, 500, 1000, 2000, 10000`.

## 2. La estrategia

Algoritmo ávido: se prueban las denominaciones de mayor a menor y, para cada
una, se usan todas las piezas que entren en el resto pendiente. Como las
denominaciones forman un sistema canónico (cada una divide o es múltiplo de
las otras), esta regla "usar siempre la pieza más grande posible" garantiza la
cantidad mínima de piezas.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `Cambiar(cantidad int) map[int]int` | Un mapa con la cantidad de piezas de cada denominación para dar el cambio. |

El mapa solo incluye las denominaciones que se usan: si `cantidad` es `1234`,
devuelve `{1000: 1, 200: 1, 20: 1, 10: 1, 2: 2}`, es decir cinco piezas en
total.

## 4. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./avidos/cambio/
```

## 6. Para experimentar

- Cambiá la `cantidad` en `main.go`. Probá con montos chicos (donde entran
  monedas de `1` y `2`) y con montos grandes (donde dominan los billetes de
  `10000` y `2000`).
- Probá un monto redondo (por ejemplo `5000`) y uno que no se pueda desarmar
  con piezas grandes (`3`, `7`, `99`).
- ¿Qué pasa con `0`? El mapa devuelto queda vacío.

## 7. Tests

```bash
go test ./avidos/cambio/
```

## 8. Complejidad

`O(n)` con `n` la cantidad de denominaciones, que acá es fija (`12`): el
algoritmo recorre la lista de denominaciones una sola vez, de mayor a menor, y
por cada una hace una división y un resto. En la práctica el costo es
constante.
