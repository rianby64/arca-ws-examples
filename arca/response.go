package arca

import "github.com/gorilla/websocket"

func response(request *JSONRPCrequest, conn *websocket.Conn,
	result interface{}) []error {
	var response JSONRPCresponse
	response.Jsonrpc = "2.0"
	response.Context = request.Context
	response.Method = request.Method
	response.Result = result

	// response
	if len(request.ID) > 0 {
		response.ID = request.ID
		return []error{conn.WriteJSON(response)}
	}

	// broadcast
	var errors []error
	for _, connection := range conns {
		errors = append(errors, connection.WriteJSON(response))
	}
	return errors
}
