# Búsqueda binaria

## 1. El problema

Dado un slice de enteros **ordenado** de menor a mayor y un valor `x`, encontrar
la posición de `x` dentro del slice. Si `x` no está, devolver `-1`.

## 2. La estrategia

División y conquista: se compara `x` con el elemento del medio. Si es igual, se
encontró la posición; si es menor, se repite el proceso en la mitad izquierda; si
es mayor, en la mitad derecha. Cada paso descarta la mitad del problema sin
mirarla.

El caso base es un intervalo vacío (`inicio > fin`): significa que `x` no está y
se devuelve `-1`.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `BusquedaBinaria(array []int, x int) int` | La posición de `x` en `array`, o `-1` si no existe. |

Hay una función auxiliar recursiva `busquedaBinaria(array, inicio, fin, x)` que
acota el intervalo de búsqueda en cada llamada; la pública se encarga de
inicializarlo con todo el slice.

## 4. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./division_y_conquista/busqueda-binaria/
```

## 6. Para experimentar

- Cambiá los elementos a buscar en `main.go`: probá el primero, el último y un
  valor que no esté en el array.
- Modificá el array. Recordá que el algoritmo **exige** que esté ordenado: con
  un array desordenado los resultados no son confiables.
- Probá un array de un solo elemento y uno vacío.

## 7. Tests

```bash
go test ./division_y_conquista/busqueda-binaria/
```

## 8. Complejidad

`O(log n)`: en cada llamada se descarta la mitad del intervalo, así que la
cantidad de pasos crece con el logaritmo de `n` (el largo del array).
