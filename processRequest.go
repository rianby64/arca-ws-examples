package main

import (
	"github.com/gorilla/websocket"
)

// JSONRPCContainer whatever
type JSONRPCContainer struct {
	Method string
	ID     string
}

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
	if request.Method == "deleteUser" {
		return deleteUser(request, conn)
	}
	return nil
}
