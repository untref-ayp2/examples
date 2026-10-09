# Ejemplo: Problema de las N Reinas

## 1. El problema

Colocar N reinas en un tablero N × N de manera que ninguna ataque a otra:
ni en la misma fila, ni en la misma columna, ni en ninguna diagonal.

## 2. La estrategia

Backtracking sobre el esquema general del apunte: se coloca una reina por
fila y, para cada fila, se prueban las columnas. Una posición es factible si
no hay otra reina en la misma columna ni en las diagonales (usando que
`fila - columna` es constante en una diagonal y `fila + columna` en la otra).
Si una columna funciona se registra la reina y se pasa a la fila siguiente;
si ninguna columna de esa fila funciona, se quita la reina anterior (vuelta
atrás) y se prueba la siguiente columna.

## 3. Qué hace el código

La solución parcial es un `[]int`: el **índice** es la fila y el **valor** es
la columna donde va la reina de esa fila. Por ejemplo, `[2, 0, 3, 1]` para
N = 4.

| Función | Qué devuelve |
|---|---|
| `NReinas(n int) []int` | La primera solución: un slice con la columna de la reina de cada fila, o `nil` si no existe. |
| `NReinasConPaso(n int, obs ObservadorReinas) []int` | Igual, pero notifica cada evento a `obs`. |
| `NReinasTodas(n int) [][]int` | Todas las soluciones. Por ejemplo, 2 para N = 4, 10 para N = 5 y 92 para N = 8. |
| `ObservadorReinas` | Tipo con dos callbacks: `Intento(fila, columna, factible, parcial)` en **cada** intento de colocación y `Solucion(solucion) bool` al completar una solución (devolver `true` obliga a seguir buscando; `false` detiene). |

El núcleo del algoritmo son las funciones del esquema:
`esSolucion(n, solucionParcial)`, `esFactible(fila, columna, solucionParcial)`,
`registrar(solucionParcial, columna)` y `borrar(solucionParcial)`, con la
recursión general en `backtracking(n, fila, solucionParcial, obs)`.

## 4. Visualización

`reinas_visual.go` es la capa de dibujo, separada del backtracking:

- `TableroNReinas(n, reinas, actualFila, actualColumna, factible)` devuelve
  el tablero con las reinas ya ubicadas (♛), resaltando el intento actual
  en verde si es factible o en rojo (`×`) si está atacado.
- `LeerCantidadReinas(entrada, salida)` pregunta cuántas reinas se quieren
  (entre 4 y 10; con Enter usa 6).
- `AnimarReinas(n, entrada, salida, demora)` recorre los intentos, dibuja
  cada paso con una pausa y, al encontrar una solución, la deja apilada
  arriba mientras anima debajo la búsqueda de la próxima; tras cada
  solución pregunta si seguir.

## 5. Cómo correrlo

```bash
go run ./backtracking/reinas/
```

Primero pregunta cuántas reinas querés (entre 4 y 10; por defecto 6) y
después anima la búsqueda: se ve cómo cada reina prueba cada columna hasta
encontrar su lugar. Al hallar una solución el programa pregunta `¿Buscar
otra solución? [s/n]` y, si se responde que sí, fuerza al backtracking a
continuar hasta la próxima. Cuando se agotan, avisa cuántas encontró.

## 6. Para experimentar

El programa ya te deja elegir al arrancar: pregunta cuántas reinas querés
(entre 4 y 10; con Enter usa 6). Además, desde el código:

- `NReinasTodas(n)` te dice **cuántas** soluciones hay: 2 para 4 reinas, 10
  para 5, 4 para 6 y 92 para 8.
- `NReinasConPaso(n, obs)` te deja engancharte en cada intento y en cada
  solución; el observador decide si seguir buscando o parar.
- La pausa entre pasos es el último argumento de `AnimarReinas` (en
  `main.go`): bajala para 8 o 10 reinas, que tienen muchísimas más ramas.

## 7. Tests

```bash
go test ./backtracking/reinas/
```

Verifican que las soluciones sean válidas, cuántas hay para cada N y que la
animación se detenga o continúe según la respuesta del usuario. Agregá tus
propios valores de N y probá.

## 8. Complejidad

`O(n!)` en el peor caso (la cantidad de formas de ubicar reinas sin
restricciones), con poda fuerte: la factibilidad descarta columnas y
diagonales atacadas, y la primera solución se alcanza mucho antes de
recorrer todo el árbol. La memoria es `O(n)` (una solución por fila).
