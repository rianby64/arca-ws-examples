package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		defer (func() {
			conn.Close()
		})()

		if err != nil {
			log.Println("connecting", err)
			return
		}

		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println("reading", err)
				return
			}
			msgTr := []byte(string(p) + " from server") // just if don't believe
			if err := conn.WriteMessage(messageType, msgTr); err != nil {
				log.Println("sending", err)
				return
			}
			log.Println("end loop")
		}
	})

	http.Handle("/", http.FileServer(http.Dir("./static")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
