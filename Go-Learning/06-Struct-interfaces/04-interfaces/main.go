package main
import "fmt"
//Interfaces

type Animal interface{
	Move()	string
}

type Dog struct{}
type Fish struct{}
type Parrot struct{}

func (*Dog)Move()string{
	return "I'm walking like a dog"
}

func (*Fish)Move()string{
	return "I'm swimming like a fish"
}

func (*Parrot)Move()string{
	return "I'm flying like a parrot"
}

func moveAnimal(animal Animal){
	fmt.Println(animal.Move())
}


func main(){
	dog := Dog{}
	moveAnimal(&dog)

	fish := Fish{}
	moveAnimal(&fish)

	parrot := Parrot{}
	moveAnimal(&parrot)
	
}