# Cambio de moneda (versión simple)

## 1. El problema

Dado un monto y una lista de denominaciones de monedas, armar el monto con la
**menor cantidad de monedas posible**. Por ejemplo, con `{1, 2, 9, 10, 20, 50}`
el monto `27` se arma con tres piezas (`9 + 9 + 9`), y no existe ninguna
combinación de dos monedas que sume 27.

## 2. La estrategia

Backtracking: se prueban todas las combinaciones de monedas que no se pasen del
monto y, cuando una suma exactamente, se compara con la mejor encontrada hasta
el momento. Dos podas aceleran la búsqueda: no probar monedas más grandes que lo
que queda por cambiar y abandonar las ramas que ya tienen más piezas que la
mejor solución conocida.

## 3. Qué hace el código

| Función | Qué hace |
|---|---|
| `CambioDeMoneda(cantidad int, denominaciones []int) []int` | Devuelve la mejor combinación de monedas, o `nil` si el monto no se puede armar. |
| `backtracking(cantidad int, denominaciones []int, solucion []int, mejorSolucion *[]int)` | La recursión del esquema: caso base, poda, extender y registrar. |

El esquema (caso base, poda, extender, registrar y borrar) es el mismo de la
versión animada en [`../cambio/`](../cambio/): acá no hay observador ni
callbacks, la función simplemente devuelve el resultado.

## 4. Cómo correrlo

```bash
go run ./backtracking/cambio-simple/
```

## 5. Para experimentar

- Cambiá `cantidad` y `denominaciones` en `main.go`.
- Probá `CambioDeMoneda(3, []int{5, 10})`: no hay solución, la función devuelve
  `nil`.
- Sacá la poda del largo (`len(solucion) >= len(*mejorSolucion)`) y observá
  que el resultado es el mismo pero se exploran muchas más ramas.

## 6. Tests

```bash
go test ./backtracking/cambio-simple/
```

## 7. Complejidad

Exponencial en el peor caso: sin poda explora todas las combinaciones de
monedas (`O(d^n)` con `d` denominaciones). Las podas reducen mucho el árbol en
la práctica, pero no cambian el orden del peor caso.
