package main

import (
	"log"
	"net/http"

	"./arca"
	"./examples"
)

func main() {

	examples.Start()
	http.HandleFunc("/ws", arca.Handle)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Serving")
}
