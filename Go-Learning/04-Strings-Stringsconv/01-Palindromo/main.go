package main

import (
	"fmt"
	"strings"
)



//Detectar una palabra palindroma
func Espalindromo(palabra string)bool{
	fmt.Printf("Ingresa una palabra: ")
	fmt.Scanln(&palabra)
	palabra = strings.ToLower(palabra)
	
	palabra = strings.ReplaceAll(palabra," ","")//reemplazamos los espacios por nada

	palabraInvertida:= reverse(palabra)
	return palabra == palabraInvertida
	
}



func reverse(cadena string)string{
	
	arrayCadena:= strings.Split(cadena,"")			//split recibe un string y lo convierte en un array
	arrayInvertido:= make([]string,0)	//creamos un array vacio
	for i := len(arrayCadena)-1; i>=0; i--{  //recorremos el array de atras hacia adelante
		arrayInvertido = append(arrayInvertido,arrayCadena[i]) //agregamos al array vacio el array invertido
	}
	return strings.Join(arrayInvertido,"") //unimos el array invertido
	
	
}



func main(){
	if Espalindromo("Luz azul"){
		fmt.Println("Es palindromo")
	}else{
		fmt.Println("No es palindromo")
	}
	


}