package main

import "fmt"

func main() {
	a:= 4
	b:= 5

	var r bool 

	//igualdad
	r = a == b
	fmt.Printf("%d es igual a %d? %t \n", a, b, r) 
//Desigualdad
	r = a != b
fmt.Printf("%d es DISTINTO a %d? %t \n", a, b, r) 
//Mayor que
r = a > b
fmt.Printf("%d es mayor  a %d? %t \n", a, b, r) 
//Menor que 
r = a < b
fmt.Printf("%d es menor a %d? %t \n", a, b, r) 
}