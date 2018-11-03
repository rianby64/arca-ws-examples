package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(w, r)
	})

	http.Handle("/", http.FileServer(http.Dir("./static")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
