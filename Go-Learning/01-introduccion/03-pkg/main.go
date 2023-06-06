package main
import "fmt"

func main() {
	hola := "hola"
	mundo := "mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

nombre:= "valen"
edad:= 23

fmt.Printf("Hola, %s y si edad es %d \n", nombre, edad)

mensaje:= fmt.Sprintf("hola, %s y si edad es %d ", nombre, edad)
fmt.Println(mensaje)

fmt.Printf("nombre, %T \n", nombre)

fmt.Printf("ingrese un nombre")
fmt.Scanln(&nombre)

fmt.Printf("ingrese una edad")
fmt.Scanln(&edad)

fmt.Println(nombre, edad)
}