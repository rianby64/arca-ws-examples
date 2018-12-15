package main

import (
	"log"
	"net/http"

	example "./example_db"
)

func main() {

	example.Start()
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Serving")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
