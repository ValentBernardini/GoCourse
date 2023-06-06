package main
import "fmt"
//suma
func suma(c, r float32) float32 {
return c+r
}


//cociente y residuo
func cociente(a, b float32) float32 {
fmt.Printf("Ingrese el numero 1: ")
fmt.Scanf("%d", &a)
fmt.Printf("Ingrese el numero 2: ")
fmt.Scanf("%d", &b)
fmt.Println("El cociente es: ", a/b)
return a/b
}


func main() {
	var a,b,c,r float32
	
fmt.Println("ingrese el numero 1: ")
fmt.Scanln(&a)
fmt.Println("ingrese el numero 2: ")
fmt.Scanln(&b)



suma(c,r)
cociente(a,b)
	}
	