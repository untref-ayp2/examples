package main

// Componente define la operación común a archivos y carpetas.
type Componente interface {
	Tamanio() int64
}

// Archivo representa un elemento simple (hoja) del sistema.
type Archivo struct {
	nombre string
	bytes  int64
}

func (a *Archivo) Tamanio() int64 {
	return a.bytes
}

// Carpeta representa un elemento compuesto que contiene otros componentes.
type Carpeta struct {
	nombre      string
	componentes []Componente
}

func (c *Carpeta) Tamanio() int64 {
	var total int64
	for _, comp := range c.componentes {
		total += comp.Tamanio()
	}
	return total
}

func (c *Carpeta) Agregar(comp Componente) {
	c.componentes = append(c.componentes, comp)
}
