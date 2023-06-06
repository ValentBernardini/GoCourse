package main
import "fmt"

func main() {
var numeros[5]int 
fmt.Println(numeros)

numeros[0] = 10
numeros[1] = 20
numeros[2] = 30
fmt.Println(numeros)
fmt.Println(numeros[2])



nombres := [3]string{"juan","pedro","maria"}
fmt.Println(nombres)

colores:=[...]string{"rojo","verde","azul"}
fmt.Println(colores, len(colores))
}