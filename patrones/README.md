# Patrones de diseño

En esta carpeta se encuentran los ejemplos de patrones de diseño que se ven en
clase y se resuelven en el aula. Cada ejemplo vive en su propia carpeta y se
puede ejecutar y testear de forma independiente.

## Ejemplos

| Ejemplo | Descripción | Cómo correrlo |
|---|---|---|
| [adapter/bicicleta](adapter/bicicleta/) | Adaptador que traduce el `PedalearHaciaAtras` de una bicicleta playera a la interfaz `Bicicleta` (`Frenar`). | `go run ./patrones/adapter/bicicleta/` |
| [adapter/robot](adapter/robot/) | Adaptador que traduce las medidas de un robot al método `Medir` de la interfaz `Medible`. | `go run ./patrones/adapter/robot/` |
| [composite](composite/) | Composición de archivos y carpetas bajo una interfaz común `Componente`. | `go run ./patrones/composite/` |
| [iterator/matriz](iterator/matriz/) | Iterador sobre una matriz con la interfaz `Siguiente`/`Valor`. | `go run ./patrones/iterator/matriz/` |
| [iterator/lista-enlazada](iterator/lista-enlazada/) | Iterador sobre una lista enlazada con la misma interfaz `Siguiente`/`Valor`. | `go run ./patrones/iterator/lista-enlazada/` |

Los ejemplos de `adapter` y de `iterator` comparten un `README.md` propio de la
carpeta que compara las dos implementaciones del mismo patrón.

## Tests

```bash
go test ./patrones/...
```
