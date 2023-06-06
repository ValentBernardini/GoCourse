


package figuras


type Cuadrado struct {
	Lado float64
}


func (c *Cuadrado) Area() float64 {
	return c.Lado * c.Lado
}

func (c *Cuadrado) Perimetro() float64 {
	return c.Lado * 4
}



