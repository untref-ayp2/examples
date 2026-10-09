# Ejemplos — Algoritmos y Programación II

Repositorio público con los ejemplos de código del bloque de Diseño de
Algoritmos de la materia **Algoritmos y Programación II** (UNTREF y CUDI). Reúne
en un solo lugar los ejemplos que se ven y se resuelven en clase, con una
estructura uniforme y sin dependencias externas.

Cada ejemplo vive en su propia carpeta `tema/ejemplo/`, con la lógica del
algoritmo, sus tests, un `main.go` ejecutable y un `README.md` didáctico.

## Requisitos

- Go 1.22 o superior.
- No requiere dependencias externas: solo se usa la biblioteca estándar de Go.

## Estructura

```
examples/
├── recursividad/
├── division_y_conquista/
├── patrones/
├── avidos/
├── backtracking/
├── programacion_dinamica/
└── ordenamientos/
```

Cada tema tiene un `README.md` que lista sus ejemplos. Cada ejemplo se puede
ejecutar y testear de forma independiente.

## Cómo correr un ejemplo

```bash
go run ./tema/ejemplo/
```

Por ejemplo:

```bash
go run ./recursividad/factorial/
go run ./backtracking/reinas/
go run ./programacion_dinamica/fibonacci/bottom-up/
```

Los parámetros que se pueden tocar están comentados en el `main.go` de cada
ejemplo y documentados en su `README.md`.

## Cómo correr los tests

Todos los tests del repositorio:

```bash
go test ./...
```

Los de un tema o un ejemplo puntual:

```bash
go test ./recursividad/...
go test ./programacion_dinamica/fibonacci/bottom-up/
```

## Temas

| Tema | Contenido | Ejemplos |
|---|---|---|
| [recursividad](recursividad/) | Algoritmos recursivos básicos | bases, factorial, hanoi, palindromo, par-impar |
| [division_y_conquista](division_y_conquista/) | Estrategia de división y conquista | busqueda-binaria, potencia, suma |
| [patrones](patrones/) | Patrones de diseño | adapter (bicicleta, robot), composite, iterator (matriz, lista-enlazada) |
| [avidos](avidos/) | Algoritmos ávidos (*greedy*) | actividades, cambio, combustible |
| [backtracking](backtracking/) | Exploración con retroceso | cambio, ratmaze, reinas, sudoku |
| [programacion_dinamica](programacion_dinamica/) | Programación dinámica | cambio, fibonacci, ladron, mochila |
| [ordenamientos](ordenamientos/) | Algoritmos de ordenamiento | burbujeo, insercion, seleccion, mergesort, quicksort |

## Makefile

```bash
make build    # compila todos los paquetes
make test     # corre todos los tests
make test-v   # corre todos los tests en modo verboso
make fmt      # formatea el código
make lint     # corre golangci-lint
make clean    # limpia la caché de tests
```

## Licencia

Ver [LICENSE](LICENSE).
