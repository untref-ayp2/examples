package main

// Este archivo muestra el patrón Composite: Archivo (hoja) y Carpeta
// (compuesto) implementan la misma interfaz Componente.Tamanio(). Una carpeta
// calcula su tamaño sumando el de sus componentes, de modo que el cliente
// trata igual a un archivo suelto que a un árbol entero.

// Componente define la operación común a archivos y carpetas.
type Componente interface {
	Tamanio() int64
}

// Archivo representa un elemento simple (hoja) del sistema.
type Archivo struct {
	nombre string
	bytes  int64
}

// Tamanio devuelve el tamaño del archivo en bytes.
func (a *Archivo) Tamanio() int64 {
	return a.bytes
}

// Carpeta representa un elemento compuesto que contiene otros componentes.
type Carpeta struct {
	nombre      string
	componentes []Componente
}

// Tamanio devuelve el tamaño total de la carpeta: la suma del tamaño de cada
// componente (que puede ser un archivo u otra carpeta).
func (c *Carpeta) Tamanio() int64 {
	var total int64
	for _, comp := range c.componentes {
		total += comp.Tamanio()
	}
	return total
}

// Agregar suma un componente (archivo o subcarpeta) a la carpeta.
func (c *Carpeta) Agregar(comp Componente) {
	c.componentes = append(c.componentes, comp)
}
