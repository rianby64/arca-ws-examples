package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", serveWS)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Serving")
}
