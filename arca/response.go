package arca

import (
	"github.com/gorilla/websocket"
)

func response(request *JSONRPCrequest, conn *websocket.Conn,
	result *interface{}) []error {

	var response JSONRPCresponse
	response.Context = &request.Context
	response.Method = request.Method
	response.Result = result

	// response
	if len(request.ID) > 0 {
		response.ID = request.ID
		return []error{writeJSON(conn, &response)}
	}

	// broadcast
	var errs []error
	source := request.Context.(map[string]interface{})["source"].(string)
	for _, conn := range conns {
		subscribers, ok := subscriptions[conn]
		if !ok {
			continue
		}
		for _, subscribed := range subscribers {
			if source == subscribed {
				errs = append(errs, writeJSON(conn, &response))
				break
			}
		}
	}
	return errs
}
