package arca

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func matchHandler(request *JSONRPCrequest,
	conn *websocket.Conn) (*requestHandler, error) {
	if request.Method == "" {
		return nil, fmt.Errorf("Method must be present in request")
	}
	if request.Context == nil {
		return nil, fmt.Errorf("Context must be present in request")
	}
	contextRequest, ok := request.Context.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Context must be an Object")
	}
	if contextRequest["source"] == nil {
		return nil, fmt.Errorf("Context must define a source")
	}
	sourceRequest, ok := contextRequest["source"].(string)
	if !ok {
		return nil, fmt.Errorf(
			"Context has an incorrect source expecting an string")
	}
	source, ok := handlers[sourceRequest]
	if !ok {
		return nil, fmt.Errorf("source '%s' not found in context '%s'",
			sourceRequest, contextRequest)
	}
	handler, ok := source[request.Method]
	if !ok {
		if request.Method == "subscribe" {
			var subscribeHandler requestHandler = func(_ *interface{},
				_ *interface{}) (interface{}, error) {
				subscribe(conn, sourceRequest)
				return nil, nil
			}
			handler = &subscribeHandler
		} else if request.Method == "unsubscribe" {
			var unsubscribeHandler requestHandler = func(_ *interface{},
				_ *interface{}) (interface{}, error) {
				unsubscribe(conn, sourceRequest)
				return nil, nil
			}
			handler = &unsubscribeHandler
		} else {
			return nil, fmt.Errorf(
				"Method '%s' not found. Source is '%s'",
				request.Method, sourceRequest)
		}
	}
	return handler, nil
}

func processJSONRPCrequest(request *JSONRPCrequest, conn *websocket.Conn) {
	handler, err := matchHandler(request, conn)
	if err != nil {
		log.Println("context error", err)
		return
	}
	result, err := (*handler)(&request.Params, &request.Context)
	if err != nil {
		log.Println("handler error", err)
		return
	}
	if result == nil && err == nil {
		log.Println("Empty result")
		return
	}
	for _, err := range response(request, conn, &result) {
		if err != nil {
			log.Println("response error", err)
		}
	}
}
