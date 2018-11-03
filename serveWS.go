package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// JMessage is the structure of the message that comes from there to here
type JMessage struct {
	Message string
}

// is for serving WS
func serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	defer (func() {
		conn.Close()
	})()

	if err != nil {
		log.Println("connecting", err)
		return
	}

	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			log.Println("reading", err)
			return
		}

		var action JMessage
		if err := json.Unmarshal(data, &action); err != nil {
			log.Println(err)
		} else if err := conn.WriteMessage(messageType, data); err != nil {
			log.Println("sending", err)
		}

		log.Println("end loop")
	}
}
