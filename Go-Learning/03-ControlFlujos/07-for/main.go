package main

import "fmt"

func main() {
	//bucle infinito
	/*for {
		fmt.Println("Hola")
	}
*/
	//bucle witpo while
	numeros:= 12455
	c:= 0
	for numeros> 0{
		numeros/=10
		c++
	}
	fmt.Println("la cantidad de digitos es : ",c)

	//bucle for
	for i:= 0; i<=10; i++{
		if i%2 == 0{		//solo si el numero es par
		fmt.Println(i)}
	}
}