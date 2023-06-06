package figuras
import "fmt"

type figura interface {
	Area() float64
	Perimetro() float64
}


func Medidas(f figura){
	fmt.Println(f)
	fmt.Println("Area: ", f.Area())
	fmt.Println("Perimetro: ", f.Perimetro())
}

