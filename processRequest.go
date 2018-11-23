package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func response(request *JSONRPCrequest, conn *websocket.Conn,
	result interface{}, id *string) error {
	var myResponse JSONRPCresponse
	myResponse.Jsonrpc = "2.0"
	myResponse.Context = request.Context
	myResponse.Method = request.Method
	myResponse.Result = result

	// response
	if id != nil {
		myResponse.ID = *id
		return conn.WriteJSON(myResponse)
	}

	// broadcast
	for _, connection := range conns {
		connection.WriteJSON(myResponse)
	}
	return nil
}

type requestHandler func(request *JSONRPCrequest, conn *websocket.Conn) error
type requestHandlers map[string]requestHandler

var handlers = requestHandlers{
	"getUsers":   getUsers,
	"updateUser": updateUser,
	"insertUser": insertUser,
	"deleteUser": deleteUser,
}

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
