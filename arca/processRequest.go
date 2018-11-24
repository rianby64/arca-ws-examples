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
		log.Println("handler error", err)
		return
	}
	for _, err := range response(request, conn, result) {
		if err != nil {
			log.Println("response error", err)
		}
	}
}
