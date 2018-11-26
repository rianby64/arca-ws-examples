package arca

import (
	"errors"

	"github.com/gorilla/websocket"
)

func response(request *JSONRPCrequest, conn *websocket.Conn,
	result interface{}) []error {

	var response JSONRPCresponse
	response.Context = request.Context
	response.Method = request.Method
	response.Result = result

	// response
	if len(request.ID) > 0 {
		response.ID = request.ID
		return []error{conn.WriteJSON(response)}
	}

	// broadcast
	var errs []error
	subscribeds, ok := subscriptions[conn]
	if !ok {
		errs = append(errs, errors.New("Incredible"))
		return errs
	}
	for _, conn := range conns {
		source := response.Context.(map[string]interface{})["source"].(string)
		for _, subscribed := range subscribeds {
			if source == subscribed {
				errs = append(errs, conn.WriteJSON(response))
				break
			}
		}
	}
	return errs
}
