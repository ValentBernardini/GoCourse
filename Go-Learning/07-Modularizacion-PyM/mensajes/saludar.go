package mensajes

import "fmt"

func Saludar() {
	fmt.Println("Hola, ¿cómo estás?")
}

const mensaje = "Hola, ¿cómo estás?"  //Private 

func privateFunction() {					//Private
	fmt.Println("Hola, ¿cómo estás?")
}

func Imprimir() {
	fmt.Println(mensaje)
	privateFunction()
}