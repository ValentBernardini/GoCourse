package main
import "fmt"


func main(){
	if nombre, edad:= "klen",18;nombre == "valen"{
		fmt.Println("Hola",nombre) 
	}else{
		fmt.Printf("Nombre: %s, Edad: %d",nombre,edad)
	}


	//Obetener valor de un mapa
	users:= make(map[string]string)
	users["klen"]="klen@gmail.com" 
	users["valen"]="valen@gmail.com"
	
	
	if correo, ok:= users["klen"];ok{
		fmt.Println(correo,ok)
	}else{
		fmt.Println("No se encontro el usuario")
	}


	edades:= make(map[string]int)
	edades["klen"]=18
	edades["valen"]=20

	if edad,ok:= edades["klen"];ok{
		fmt.Println(edad,ok)
	}else{
		fmt.Println("No se encontro la edad")
	}
	 
}