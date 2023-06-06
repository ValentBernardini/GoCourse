package main

import (
	"fmt"
	"os"
)


func main() {

	//Funcion
	defer func(){
		if error := recover(); error != nil{  		//recover es una funcion que se ejecuta cuando hay un panic
			fmt.Println("Se ha cerrado de manera inseperada el programa")
		}
	}()

	if file, error := os.Open("hello.txt") ;error != nil{ 	// abre el archivo hello.txt
		panic("Error abriendo el archivo")		// si hay error, muestra el mensaje
	}else{		// si no hay error
		defer func() {  	//
			fmt.Println("Cerrando el archivo") 
			file.Close()	// cierra el archivo
		}()
		contenido := make([]byte, 254) // crea un slice de bytes de 254 posiciones
		long,_ := file.Read(contenido)		// lee el archivo y lo guarda en el slice de bytes
		contenidoArchivo:= string(contenido[:long])		// convierte el slice de bytes en un string
		fmt.Println(contenidoArchivo)		// imprime el contenido del archivo
	}
	
}