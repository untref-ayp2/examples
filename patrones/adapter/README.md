# Adapter

## 1. El problema

El patrón *Adapter* permite que dos componentes con interfaces incompatibles
trabajen juntos: un adaptador envuelve al componente existente (el
*Adaptado*) y traduce su interfaz a la que espera el cliente.

En esta carpeta hay **dos implementaciones independientes** del mismo patrón,
cada una en su propia subcarpeta:

| Subcarpeta | Interfaz del cliente | Adaptado (legacy) | Adaptador |
|---|---|---|---|
| `bicicleta` | `Bicicleta.Frenar()` | `Playera.PedalearHaciaAtras()` | `PlayeraAdapter.Frenar()` |
| `robot` | `Medible.Medir() float64` | `Robot.Medir() (int, int)` | `RobotAdapter.Medir() float64` |

## 2. La estrategia

En `bicicleta`, la `Playera` es código *legacy*: frena pedaleando hacia atrás.
`PlayeraAdapter` expone `Frenar()` y delega en `PedalearHaciaAtras()`, de modo
que la playera se puede usar en cualquier lugar donde se espere una
`Bicicleta`.

En `robot`, el `Robot` mide en metros y centímetros. `RobotAdapter.Medir()`
convierte la medición a pulgadas para satisfacer la interfaz `Medible`, sin
tocar al `Robot`.

En ambos casos el cliente programa contra la interfaz (`Bicicleta` o
`Medible`) y no conoce la existencia del componente *legacy*.

## 3. Qué hace el código

| Firma | Rol |
|---|---|
| `Bicicleta.Frenar()` | Interfaz que espera el cliente: frena la bicicleta. |
| `Playera.PedalearHaciaAtras()` | Implementación legacy de la playera: frena pedaleando hacia atrás. |
| `BicicletaDeCarreras.Frenar()` | Implementación directa de la interfaz. |
| `PlayeraAdapter.Frenar()` | Traduce `Frenar()` a `PedalearHaciaAtras()`. |
| `Medible.Medir() float64` | Interfaz que espera el cliente: una distancia en pulgadas. |
| `Robot.Medir() (int, int)` | Medición legacy del robot: metros y centímetros. |
| `RobotAdapter.Medir() float64` | Traduce metros/centímetros a pulgadas. |

## 4. Cómo correrlos

```bash
go run ./patrones/adapter/bicicleta/
go run ./patrones/adapter/robot/
```

## 5. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 6. Para experimentar

- Cambiá la implementación legacy de la playera (`PedalearHaciaAtras`) y mirá
  cómo el resto del programa sigue funcionando sin cambios: el adaptador
  aísla al cliente de lo que cambia adentro.
- Cambiá las medidas del `Robot` en `main.go` y verificá la conversión a
  pulgadas. ¿Qué pasa si `centimetros` supera los 100?
- Agregá otro adaptador (por ejemplo, una `Moto` legacy que frena con el pie)
  y usala donde se espera una `Bicicleta`.

## 7. Complejidad

—: el patrón agrega una capa de indirección constante; no cambia el orden de
crecimiento del algoritmo subyacente.
