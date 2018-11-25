package arca

import (
	"log"

	"github.com/gorilla/websocket"
)

func processJSONRPCrequest(request *JSONRPCrequest, conn *websocket.Conn) {
	handler, ok := handlers["Users"][request.Method]
	if !ok {
		log.Println("There's no handler for", request.Method)
		return
	}
	result, err := handler(&request.Params, &request.Context)
	if err != nil {
		log.Println("handler error", err)
		return
	}
	for _, err := range response(request, conn, result) {
		if err != nil {
			log.Println("response error", err)
		}
	}
}
