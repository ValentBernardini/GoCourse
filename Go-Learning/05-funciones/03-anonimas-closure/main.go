package main
import "strings"
import "fmt"
//Closure

func repeat(n int) func(cadena string)string{
	return func(cadena string)string {
		return strings.Repeat(cadena, n)
	}
}

func main() {

	repeat3:= repeat(3)
	fmt.Println(repeat3("Hola"))
										//funcion anonima
	/*func()  {
		fmt.Println("Hola Mundo")
	}()
	

	myfunc := func(nombre string) string {
		return fmt.Sprintf("Hola %s", nombre)

	}

	fmt.Println(myfunc("valen"))
	*/
}