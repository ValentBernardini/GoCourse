package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkServer(server string, channel chan string) {
	_, err :=  http.Get(server)
	if err != nil {
		//fmt.Println(server, "is down!")
		channel <- server + " is down!"
	} else {
		//fmt.Println(server, "is working!")
		channel <- server + " is working!"
	}
}

func main() {
	init := time.Now()					// Tiempo actual

	channel := make(chan string)		// Creación de un canal


	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",

	}
	for _, server := range servers {
		go checkServer(server, channel)
	}

	for i := 0; i < len(servers); i++ {			// Recorremos el canal
		fmt.Println(<-channel)						// Imprimimos el canal
	}

	timeElapsed := time.Since(init)				// Tiempo que ha pasado desde que se ejecutó init
	fmt.Println("Time elapsed: ", timeElapsed)
}