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
				_ *interface{}, response chan interface{}) error {
				subscribe(conn, sourceRequest)
				close(response)
				return nil
			}
			handler = &subscribeHandler
		} else if request.Method == "unsubscribe" {
			var unsubscribeHandler requestHandler = func(_ *interface{},
				_ *interface{}, response chan interface{}) error {
				unsubscribe(conn, sourceRequest)
				close(response)
				return nil
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

func processRequest(request *JSONRPCrequest, conn *websocket.Conn) []error {
	var errHandler []error
	handler, err := matchHandler(request, conn)
	if err != nil {
		return []error{fmt.Errorf("context error %s", err)}
	}
	resultChan := make(chan interface{})
	go (func() {
		if err := (*handler)(&request.Params, &request.Context, resultChan); err != nil {
			log.Println("error", err)
			errHandler = []error{fmt.Errorf("handler error %s", err)}
		}
	})()
	result, ok := <-resultChan
	if !ok {
		if len(errHandler) > 0 {
			return errHandler
		}
		return []error{nil}
	}
	if len(errHandler) > 0 {
		return errHandler
	}
	defer close(resultChan)
	return response(request, conn, &result)
}
