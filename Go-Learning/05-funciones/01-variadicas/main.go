package main

import "fmt"

func sumar(nombre string, numeros ...int) (string, int) {
	mensaje := fmt.Sprintf("Suma de %s es: ", nombre)
	var total int
	for _, num := range numeros {
		total += num
	}
	return mensaje, total
}

func main() {
	mensaje, result := sumar("valen", 4, 5, 8, 6)
	fmt.Println(mensaje, result)
}
