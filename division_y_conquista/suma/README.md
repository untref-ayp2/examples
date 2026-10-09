# Suma de un array

## 1. El problema

Dado un slice de enteros, devolver la suma de todos sus elementos.

## 2. La estrategia

División y conquista: se parte el slice por la mitad y se suman recursivamente
cada mitad; la suma total es la suma de ambas. Los casos base son un slice
**vacío** (suma `0`) y un slice de **un elemento** (suma ese elemento).

Cada llamada divide el problema en dos mitades del mismo tamaño, hasta llegar a
slices de a lo sumo un elemento.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `SumaArray(arreglo []int) int` | La suma de los elementos de `arreglo`. Un slice vacío suma 0. |

## 4. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./division_y_conquista/suma/
```

## 6. Para experimentar

- Cambiá los slices en `main.go`: probá el caso vacío y el de un solo elemento.
- Agregá slices con números negativos y mirá cómo se combinan las mitades.
- Pensá por qué los casos base son el vacío **y** el de un elemento: sin el
  segundo, la recursión no tendría por dónde cortar y terminaría en el vacío.

## 7. Tests

```bash
go test ./division_y_conquista/suma/
```

## 8. Complejidad

`O(n)`: cada elemento participa en una única hoja del árbol de recursión y el
trabajo por nivel es proporcional a `n`, así que el total crece linealmente con
el largo del slice.
