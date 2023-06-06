package main
import "fmt"

func modificarValor(cadena *string){
	fmt.Printf("%p\n", cadena) // imprime la direccion de memoria
	*cadena = "hola desde la nueva funcion"
}

func main() {
 cadena := "hola mundo"
 fmt.Printf("%p\n", &cadena)
 fmt.Println("Antes de la funcion",cadena)

 modificarValor(&cadena)
 fmt.Println("Despues de la funcion",cadena) 
}