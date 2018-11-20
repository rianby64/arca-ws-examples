package main

import (
	"github.com/gorilla/websocket"
)

func response(request *JSONRPCrequest, conn *websocket.Conn,
	result interface{}, id *string) error {
	var myResponse JSONRPCresponse
	myResponse.Jsonrpc = "2.0"
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

func processRequest(request *JSONRPCrequest, conn *websocket.Conn) error {
	if request.Method == "getUsers" {
		return getUsers(request, conn)
	}
	if request.Method == "updateUser" {
		return updateUser(request, conn)
	}
	if request.Method == "insertUser" {
		return insertUser(request, conn)
	}
	return nil
}
