package main
import "fmt"


//variable global
var mensaje string
func function1(){
	mensaje = "Hola desde la funcion 1"
	fmt.Println(mensaje)
}

func function2(){
	mensaje = "Hola desde la funcion 2"
	fmt.Println(mensaje)
}

func main(){
	mensaje = "Hola desde la funcion main"
	fmt.Println(mensaje)
	defer function1() // defer es una funcion que se ejecuta al final de la funcion main
	function2()
	fmt.Println(mensaje)
}