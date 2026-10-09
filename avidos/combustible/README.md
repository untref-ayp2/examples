# Autito a combustible (recargas mínimas)

## 1. El problema

Un auto debe recorrer `distancia` kilómetros con un tanque que le da
`capacidad` kilómetros de autonomía. En el camino hay estaciones de servicio
en kilómetros conocidos, ordenadas de menor a mayor. Hay que calcular la
**cantidad mínima de recargas** necesarias para llegar a destino, sabiendo que
el tanque arranca lleno.

## 2. La estrategia

Algoritmo ávido: en cada tramo se avanza hasta la **estación más lejana que se
puede alcanzar con el tanque actual** y se recarga ahí. Elegir la más lejana
maximiza el avance por recarga y deja el resto del camino más corto; se puede
demostrar que esta regla usa la menor cantidad de recargas posible.

El caso de corte es cuando el tanque actual alcanza para llegar al destino:
ahí no hace falta recargar más.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `MinimasRecargas(distancia int, capacidad int, estaciones []int) int` | La cantidad mínima de recargas para recorrer la distancia. |

Internamente, `paradasDeRecarga(distancia, capacidad, estaciones) []int`
aplica la estrategia ávida y devuelve las **estaciones donde se recarga**, en
orden; `MinimasRecargas` es el largo de esa lista. `estaciones` no se modifica.

## 4. Visualización

La capa `combustible_visual.go` dibuja el **recorrido de las estaciones** en
ASCII: una fila de camino de `0` a `distancia` (escalada para no estirarse en
terminales angostas) con `S` en la partida, `E` en cada estación de servicio,
`*` en las estaciones donde se recarga y `D` en el destino. Debajo se rotulan
los kilómetros de cada estación.

| Función | Qué hace |
|---|---|
| `dibujarRecorrido(w io.Writer, distancia int, estaciones []int, paradas []int, delay time.Duration)` | Escribe el recorrido en `w` y espera `delay` (no imprime por su cuenta). |

## 5. Cómo correrlo

```bash
go run ./avidos/combustible/
```

## 6. Para experimentar

- Cambiá la `capacidad` del tanque en `main.go`: con la carga justa para
  llegar (`950`) no hace falta recargar; con un tanque chico (`200`) hay que
  parar en todas las estaciones.
- Probá con estaciones **salteadas** (que falte una en el medio) y fijate si
  el algoritmo las evita cuando no hacen falta.
- Agregá una estación cerca del destino y otra lejos: la ávida siempre elige
  la más lejana alcanzable.
- Bajá el `delay` a 0 para ver el recorrido de una sola pasada.

## 7. Tests

```bash
go test ./avidos/combustible/
```

## 8. Complejidad

`O(n)` con `n` la cantidad de estaciones: el algoritmo recorre la lista una
sola vez y cada estación se descarta apenas deja de ser útil, así que el
trabajo total es lineal en el largo de `estaciones`.
