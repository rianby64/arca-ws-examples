package arca

import (
	"log"

	"github.com/gorilla/websocket"
)

type requestHandler func(request *JSONRPCrequest, conn *websocket.Conn) error
type requestHandlers map[string]requestHandler

var handlers = requestHandlers{}

func processRequest(request *JSONRPCrequest, conn *websocket.Conn) {
	handler, ok := handlers[request.Method]
	if !ok {
		log.Println("There's no handler for", request.Method)
	}
	err := handler(request, conn)
	if err != nil {
		log.Println(err)
	}
}
