package main
import "fmt"
import "errors"

func division(dividendo,divisor int)(int,error){
	if divisor == 0 {
		return 0, errors.New("no se puede dividir por cero")
	}else{
		return dividendo/divisor, nil // nil es el valor por defecto de error entonces al no tener errores, no muestra nada en este caso
	}
}


func main(){
	result, error:=division(4,0)
	if error == nil {
		fmt.Println("El resultado es:",result)
	}else{
		fmt.Println(error)
	}
}