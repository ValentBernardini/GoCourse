package main

import (
	"fmt"
	"math"
)

type Figura interface {
	Area() float64
	Perimetro() float64
}


type Cuadrado struct {
	lado float64
}

type circulo struct {
	radio float64
}

func (c *Cuadrado) Area() float64 {
	return c.lado * c.lado
}

func (c *Cuadrado) Perimetro() float64 {
	return c.lado * 4
}

func (c *circulo) Area() float64 {
return math.Pi * c.radio * c.radio	
}

func (c *circulo) Perimetro() float64 {
	return 2 * math.Pi * c.radio
}

func medidas(f Figura){
	fmt.Println(f)
	fmt.Println("Area: ", f.Area())
	fmt.Println("Perimetro: ", f.Perimetro())
}

func main() {
	cuadrado := Cuadrado{lado: 4}
	circulo := circulo{radio: 5}

	medidas(&cuadrado)
	medidas(&circulo)

}