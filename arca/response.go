package arca

import "github.com/gorilla/websocket"

func response(request *JSONRPCrequest, conn *websocket.Conn,
	result interface{}) error {
	var response JSONRPCresponse
	response.Jsonrpc = "2.0"
	response.Context = request.Context
	response.Method = request.Method
	response.Result = result

	// response
	if &request.ID != nil {
		response.ID = request.ID
		return conn.WriteJSON(response)
	}

	// broadcast
	for _, connection := range conns {
		connection.WriteJSON(response)
	}
	return nil
}
