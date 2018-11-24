package arca

import "github.com/gorilla/websocket"

// Response whatever
func Response(request *JSONRPCrequest, conn *websocket.Conn,
	result interface{}, id *string) error {
	var response JSONRPCresponse
	response.Jsonrpc = "2.0"
	response.Context = request.Context
	response.Method = request.Method
	response.Result = result

	// response
	if id != nil {
		response.ID = *id
		return conn.WriteJSON(response)
	}

	// broadcast
	for _, connection := range conns {
		connection.WriteJSON(response)
	}
	return nil
}
