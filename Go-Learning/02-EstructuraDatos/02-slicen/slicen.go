package main

import "fmt"


func main(){
	//slicen
numeros:= []int{1,2,3}


//agregar datos
numeros = append(numeros,4)
numeros = append(numeros,5)
numeros =append(numeros, 6,7,8)
fmt.Println(numeros, len(numeros))


//sub slicen

subSlicen := numeros[0:2]
numeros[0]=100
fmt.Println(subSlicen )
fmt.Println(numeros )

//puntero
//longitud 
//capacidad
meses:= []string{"enero","febrero","marzo"}
fmt.Printf("Len: %v, Cap: %v, %p\n", len(meses), cap(meses), meses)
meses = append(meses, "abril")
fmt.Printf("Len: %v, Cap: %v, %p\n", len(meses), cap(meses), meses)

}