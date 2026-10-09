# Cambio de moneda

## 1. El problema

Dada una cantidad de dinero y un conjunto de denominaciones (billetes y
monedas), hay que devolver el cambio para esa cantidad usando la **menor
cantidad de piezas** posible. Se puede repetir cada denominación las veces que
haga falta.

## 2. La estrategia

Backtracking: se exploran todas las combinaciones de denominaciones que suman
la cantidad pedida y se conserva la de menor longitud. En cada paso se elige
una denominación que no supere el resto (`es factible`), se la agrega a la
solución parcial (`registrar`) y se continúa con el resto; si la rama no
llega a sumar exacto, se descarta (vuelta atrás). La poda es simple: si la
solución parcial ya tiene igual o más piezas que la mejor encontrada, esa
rama no puede mejorar y se corta.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `CambioDeMoneda(cantidad int, denominaciones []int) []int` | Un slice con las monedas que dan el cambio con la menor cantidad de piezas, o `nil` si no hay combinación. |
| `CambioDeMonedaConPaso(cantidad int, denominaciones []int, obs ObservadorCambio) []int` | Igual, pero notifica cada evento a `obs`. |
| `ObservadorCambio` | Tipo con dos callbacks: `Intento(resto, parcial)` en **cada** moneda que se prueba y `Solucion(solucion, mejor)` al completar una combinación exacta. |

Con `denominaciones = [1, 2, 9, 10, 20, 50]` y `cantidad = 27`, devuelve
`[9, 9, 9]`: tres piezas, el mínimo posible.

## 4. Visualización

`cambio_visual.go` es la capa de dibujo, separada del backtracking:
`AnimarCambio(cantidad, denominaciones, salida, demora)` muestra en cada paso
la combinación parcial, el resto por cambiar y la mejor combinación
encontrada hasta el momento; al final imprime el cambio con menos piezas (o
avisa si no existe).

## 5. Cómo correrlo

```bash
go run ./backtracking/cambio/
```

## 6. Para experimentar

El programa arranca con el caso del ejemplo: cambiar `27` con
`[1, 2, 9, 10, 20, 50]`. En `main.go` podés:

- Cambiar la `cantidad`: probá montos redondos (`50`, `100`) y montos
  difíciles (`3`, `97`).
- Cambiar las `denominaciones`: ordenalas de mayor a menor o probá un
  conjunto sin combinación posible (por ejemplo `[5, 10]` con cantidad `3`),
  donde el programa avisa que no hay solución.
- Probar desde el código `CambioDeMonedaConPaso` con tu propio observador
  para engancharte en cada rama explorada.

## 7. Tests

```bash
go test ./backtracking/cambio/
```

Verifican que la solución sume exactamente la cantidad, que use el mínimo de
piezas, el caso sin solución y que el observador reciba los eventos (y que la
animación muestre la solución final).

## 8. Complejidad

Exponencial: en el peor caso se exploran todas las combinaciones posibles de
las denominaciones, `O(c^n)` con `c` la cantidad de monedas y `n` la cantidad
a cambiar. La poda por longitud y la poda de ramas que se pasan del monto
recortan el árbol en la práctica, pero el peor caso sigue siendo exponencial.
