package main

import "fmt"

func main(){
	var vocal string

	fmt.Print("Ingresa una Letra: ")
	fmt.Scanf("%s",&vocal)
	/*
	switch {
		case vocal == "a" || vocal == "e" || vocal == "i" || vocal == "o" || vocal == "u"|| vocal == "A" || vocal == "E" || vocal == "I" || vocal == "O" || vocal == "U":
			fmt.Println(vocal,"Es una vocal")
		default:
			fmt.Println(vocal,"No es una vocal")
	}
		*/
	switch vocal {
		case "a","e","i","o","u","A","E","I","O","U":	
		fmt.Println(vocal,"Es una vocal")
		default:
			fmt.Println(vocal,"No es una vocal")
	}
			



}