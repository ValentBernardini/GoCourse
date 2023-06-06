package main

import "fmt"


func main(){
	var consumo, descuento float64
	var datosDescuento string

fmt.Print("ingrese un numerp: ")
fmt.Scanln(&consumo)


	if consumo >= 0{
		if consumo <= 100{
			datosDescuento = "10% de descuento"
			descuento = consumo * 0.10
		}else if consumo > 100 && consumo <= 200{
			datosDescuento = "20% de descuento"
			descuento = consumo * 0.20
		} else if consumo > 200 {
			datosDescuento = "30% de descuento"
			descuento = consumo * 0.30
		}
		
	}else{
		fmt.Println("ingrese un numero positivo")
	}

	montoDescuento:= consumo - descuento
	iva:= montoDescuento * 1.19
	totalAPagar:= iva + montoDescuento
	fmt.Println(datosDescuento)
	fmt.Println("consumo: " ,consumo)
	fmt.Println("Descuento : ", descuento)
	fmt.Println("monto con descuento: ", montoDescuento)
	fmt.Println("iva: ", iva)
	fmt.Println("total a pagar: ", totalAPagar)


}