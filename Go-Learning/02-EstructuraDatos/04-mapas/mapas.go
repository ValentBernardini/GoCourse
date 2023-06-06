package main

import "fmt"

func main() {
dias:= make(map[int]string)
fmt.Println(dias)

//agregar datos

dias[1] = "domingo"
dias[2] ="Lunes"
dias[3] ="martes"
dias[4] ="miercoles"
dias[5] ="jueves"
dias[6] ="viernes"
dias[7] ="sabado"
fmt.Println(dias)

delete(dias,1)

//nuevo mapa
estudiantes:= make(map[string][]int)
estudiantes["juan"] = []int{1,2,3}
estudiantes["miguel"] = []int{10,4,5}
fmt.Println(estudiantes)
}