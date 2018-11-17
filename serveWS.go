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

// MyParams whatever
type MyParams struct {
	message string
	a       []string
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

		var request JSONRPCrequest
		if err := json.Unmarshal(data, &request); err != nil {
			log.Println(err)
		} else if err := conn.WriteMessage(messageType, data); err != nil {
			log.Println("sending", err)
		}

		log.Println(request.ID, request.Jsonrpc, request.Method, request.Params)
		var params MyParams

		preA := request.Params["a"].([]interface{})

		params.message = request.Params["message"].(string)
		params.a = make([]string, len(preA))

		for key, value := range preA {
			params.a[key] = value.(string)
		}

		log.Println(params)

		log.Println("end loop")
	}
}
