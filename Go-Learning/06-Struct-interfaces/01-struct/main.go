package main

import "fmt"

//struct persona

type Persona struct{
	nombre string
	edad int
}

//methods 
func (p *Persona) imprimir(){
	fmt.Printf("Nombre: %s, Edad: %d\n", p.nombre, p.edad)
}

func (p *Persona) cambiarNombre(nuevoNombre string){
	p.nombre = nuevoNombre
}

//herencia
type Empleado struct{
	Persona
	sueldo float64
}

func main(){
	p1 := Persona{"Juan", 20}
	p2 := Persona{"Pedro", 30}
	p3:= Persona{
		nombre:"klen",
		edad: 40,
	}
	p1.cambiarNombre("MIGUEL") //se puede modificar los valores de los campos
	p1.imprimir()
	p2.cambiarNombre("Pablo")	//se puede modificar los valores de los campos
	p2.imprimir()
	p3.imprimir()

	em1:= Empleado{
		sueldo: 1000,
	}
	em1.nombre = "Empleado1"
	em1.edad = 20
	em1.imprimir()
}