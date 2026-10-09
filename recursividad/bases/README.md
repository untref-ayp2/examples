# Bases

## 1. El problema

Dado un número entero no negativo `n`, queremos escribir su representación en
base 10 y en una base arbitraria `base` (por ejemplo, 2, 8 o 16) como texto.

## 2. La estrategia

La representación se arma dígito a dígito, del más significativo al menos
significativo. El último dígito es el resto de dividir por la base (`n % base`);
los que van antes salen de la división entera (`n / base`), que recorta el
número y permite repetir el mismo razonamiento: es una recursión natural. El
caso base aparece solo: cuando `n < base`, la representación es el dígito
correspondiente a `n`.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `ImprimirEnDecimal(n int) string` | `n` como texto en base 10. |
| `ImprimirEnBase(n int, base int) string` | `n` como texto en la `base` dada (hasta base 16; los dígitos van de `0` a `f`). |

Ninguna de las dos imprime: devuelven un `string` y el programa principal lo
muestra en pantalla.

## 4. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./recursividad/bases/
```

## 6. Para experimentar

- Cambiá la base: probá `ImprimirEnBase` con 2, 8 y 16 y compará los resultados.
- Usá números grandes: con `n = 2026` o `n = 65535` se ve bien la descomposición.
- Probá `ImprimirEnBase(0, 2)`: ¿qué devuelve y por qué?

## 7. Tests

```bash
go test ./recursividad/bases/
```

## 8. Complejidad

`O(log n)`: cada llamada divide `n` por la base, así que la profundidad de la
recursión crece con la cantidad de dígitos de `n` en esa base.
