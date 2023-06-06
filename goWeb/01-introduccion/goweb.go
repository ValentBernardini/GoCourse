package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("el metodo es: ", r.Method)
	fmt.Fprintln(rw, "Hola mundo de Goweb")
}

func PageNF(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}
func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Error",	http.StatusInternalServerError)
}

func Saludar(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.Query())		//Query convierte los parametros en un map
	name:= r.URL.Query().Get("name")
	age:= r.URL.Query().Get("age")

	fmt.Fprintf(rw, "Hola %s tienes %s años", name, age)
}


func main() {
	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/page", PageNF)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	//Routes
	// http.HandleFunc("/", Hola)
	// http.HandleFunc("/page", PageNF)
	// http.HandleFunc("/error", Error)
	// http.HandleFunc("/saludar", Saludar)

	//Create server
	server:= &http.Server{
		Addr: "localhost:3000",
		Handler: mux,
	}
	fmt.Println("Server running on port 3000")
	fmt.Println("Run Server: http://localhost:3000")
	//log.Fatal(http.ListenAndServe("localhost:3000", mux))
	log.Fatal(server.ListenAndServe())

	//Ejecutar en terminal:  go run github.com/pilu/fresh
}
