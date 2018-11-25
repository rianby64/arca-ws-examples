package main

import (
	"log"
	"net/http"

	"./arca"
)

func main() {
	arca.RegisterSource("Users", usersCRUD)

	http.HandleFunc("/ws", arca.Handler)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Serving")
}
