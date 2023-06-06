package main

func main(){
	nombres:= [...]string{"pepe","juan","luis","jose"}
	for i:=0; i<len(nombres); i++{
		println(nombres[i])
	}

	//For each
	for _, elemento:= range nombres{
		println(elemento)
	}
}