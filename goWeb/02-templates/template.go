package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)


//Estructura
type User struct {
	UserName string
	Email string
	Age int
	Active bool
	Admin bool
	Cursos []Curso
}

type Curso struct {
	Name string
}

func Index(rw http.ResponseWriter, r *http.Request) {
	c1:= Curso{"Go desde cero"}
	c2:= Curso{"Go avanzado"}
	c3:= Curso{"Go intermedio"}
	c4:= Curso{"Go profesional"}
	template, err:=	template.ParseFiles("index.html")

	cursos := []Curso {c1, c2, c3, c4}	//Slice de cursos
	user := User{"Valen", "valentinbernardini@gmail.com", 25, true, true, cursos}

		if err != nil {
			panic(err)
		}else{
			template.Execute(rw, user)
		}
}

func main() {
	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index) 


	//Create server
	server:= &http.Server{
		Addr: "localhost:3000",
		Handler: mux,
	}
	fmt.Println("Server running on port 3000")
	fmt.Println("Run Server: http://localhost:3000")
	//log.Fatal(http.ListenAndServe("localhost:3000", mux))
	log.Fatal(server.ListenAndServe())
}