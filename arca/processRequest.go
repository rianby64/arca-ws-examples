package arca

import (
	"log"

	"github.com/gorilla/websocket"
)

type requestHandler func(requestParams *interface{},
	context *interface{}) (interface{}, error)
type requestHandlers map[string]map[string]requestHandler

var handlers = requestHandlers{}

func processRequest(request *JSONRPCrequest, conn *websocket.Conn) {
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

// RegisterSource whatever
func RegisterSource(name string, methods IRUD) {
	handlers[name] = map[string]requestHandler{
		"insert": methods.Insert,
		"read":   methods.Read,
		"update": methods.Update,
		"delete": methods.Delete,
	}
}
