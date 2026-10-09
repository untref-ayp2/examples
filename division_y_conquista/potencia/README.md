# Potencia

## 1. El problema

Calcular `n^x` (base `n` elevada al exponente entero no negativo `x`) sin usar la
función de potenciación del lenguaje.

## 2. La estrategia

División y conquista sobre el exponente:

- `n^0 = 1` y `n^1 = n` son los casos base.
- Si `x` es par: `n^x = (n^(x/2))^2`. Se calcula una sola vez `n^(x/2)` y se
  multiplica por sí misma.
- Si `x` es impar: `n^x = n * (n^((x-1)/2))^2`.

Al dividir el exponente a la mitad en cada paso, la cantidad de llamadas
recursivas crece mucho más lento que con la definición ingenua por producto
repetido.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `Potencia(n int, x int) int` | El valor de `n^x` para `x >= 0`. |

## 4. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./division_y_conquista/potencia/
```

## 6. Para experimentar

- Cambiá los casos en `main.go`: probá exponentes pares e impares, y el caso
  `x = 0`.
- Probá `Potencia(2, 16)`: notá la diferencia con hacer 16 multiplicaciones.
- Cuidado con exponentes grandes: el resultado puede desbordar el `int`.

## 7. Tests

```bash
go test ./division_y_conquista/potencia/
```

## 8. Complejidad

`O(log x)`: cada llamada divide el exponente a la mitad, así que la cantidad de
pasos crece con el logaritmo de `x`.
