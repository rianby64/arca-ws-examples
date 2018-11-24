package arca

import (
	"log"

	"github.com/gorilla/websocket"
)

type requestHandler func(requestParams *interface{}) (interface{}, error)
type requestHandlers map[string]requestHandler

var handlers = requestHandlers{}

func processRequest(request *JSONRPCrequest, conn *websocket.Conn) {
	handler, ok := handlers[request.Method]
	if !ok {
		log.Println("There's no handler for", request.Method)
		return
	}
	result, err := handler(&request.Params)
	if err != nil {
		log.Println(err)
		return
	}
	if err := response(request, conn, result); err != nil {
		log.Println(err)
		return
	}
}
