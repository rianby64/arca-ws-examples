package arca

import (
	"log"

	"github.com/gorilla/websocket"
)

func processJSONRPCrequest(request *JSONRPCrequest, conn *websocket.Conn) {
	if request.Context == nil {
		log.Println("Context must be present in request", request)
		return
	}
	contextRequest, ok := request.Context.(map[string]interface{})
	if !ok {
		log.Println("Context must be an Object", request)
		return
	}
	if contextRequest["source"] == nil {
		log.Println("Context must define a source", request)
		return
	}
	sourceRequest, ok := contextRequest["source"].(string)
	if !ok {
		log.Println("Context has an incorrect source",
			contextRequest, "expecting an string")
		return
	}

	source, ok := handlers[sourceRequest]
	if !ok {
		log.Printf("source not found for '%s'", sourceRequest)
		return
	}

	handler, ok := source[request.Method]
	if !ok {
		log.Printf("handler not found for '%s' in source '%s'",
			request.Method, sourceRequest)
		return
	}
	result, err := handler(&request.Params, &request.Context)
	if err != nil {
		log.Println("handler error", err)
		return
	}
	for _, err := range response(request, conn, result) {
		if err != nil {
			log.Println("response error", err)
		}
	}
}
