package main

import "fmt"

func saludar(nombre string){
	fmt.Println("Hola ", nombre)
}

func sumar(a int, b int) int {
	return	a + b
}


func main(){
	saludar("valen")
	r:=sumar(3,4)
 	fmt.Println("la suma es: ",r)
}