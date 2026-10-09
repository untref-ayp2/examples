# Palíndromo

## 1. El problema

Decidir si un texto es un palíndromo: una palabra o frase que se lee igual de
izquierda a derecha que de derecha a izquierda (por ejemplo, «reconocer»).

## 2. La estrategia

Un texto es palíndromo si su primer carácter es igual al último **y** el texto
interior (todo lo que queda entre el primero y el último) también lo es. Ese
«también» es la recursión: el problema se reduce a un texto cada vez más
chico. El caso base aparece cuando el texto tiene 0 o 1 carácter: siempre es
palíndromo.

## 3. Qué hace el código

| Función | Qué devuelve |
|---|---|
| `EsPalindromo(texto string) bool` | `true` si `texto` es palíndromo, `false` en caso contrario. |

Devuelve `bool`; el programa principal prueba varias palabras.

## 4. Visualización

No tiene capa visual: la salida es texto plano en la terminal.

## 5. Cómo correrlo

```bash
go run ./recursividad/palindromo/
```

## 6. Para experimentar

- El ejemplo compara los caracteres exactos: distingue mayúsculas y no ignora
  espacios. Por eso `ANITA LAVA LA TINA` da `false` aunque, normalizado, sea
  palíndromo.
- Mayúsculas: probá `DABALEARROZALAZORRAELABAD` en minúsculas.
- Acentos: probá palabras como «aviva» y variantes acentuadas.
- Espacios: ¿cómo habría que modificar el código para que ignore los espacios
  y las mayúsculas?

## 7. Tests

```bash
go test ./recursividad/palindromo/
```

## 8. Complejidad

`O(n)`: cada llamada compara el primer y el último carácter y recorta dos del
texto, así que hay una llamada por cada par de caracteres.
