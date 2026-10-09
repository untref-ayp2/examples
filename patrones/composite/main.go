package main

import "fmt"

func main() {
	readme := &Archivo{nombre: "README.md", bytes: 2048}
	licencia := &Archivo{nombre: "LICENSE", bytes: 1024}

	src := &Carpeta{nombre: "src"}
	src.Agregar(&Archivo{nombre: "main.go", bytes: 4096})
	src.Agregar(&Archivo{nombre: "utils.go", bytes: 1536})

	docs := &Carpeta{nombre: "docs"}
	docs.Agregar(&Archivo{nombre: "manual.pdf", bytes: 524288})

	proyecto := &Carpeta{nombre: "mi-proyecto"}
	proyecto.Agregar(readme)
	proyecto.Agregar(licencia)
	proyecto.Agregar(src)
	proyecto.Agregar(docs)

	fmt.Println(proyecto.Tamanio())
	// Salida: 532992
}
