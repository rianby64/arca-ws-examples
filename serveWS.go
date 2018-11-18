package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// MyParams whatever
type MyParams struct {
	Message string
	A       []string
}

type errorMyParams struct {
	why string
}

func (e errorMyParams) Error() string {
	return e.why
}

// is for serving WS
func serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	defer (func() {
		conn.Close()
	})()

	if err != nil {
		log.Println("connecting", err)
		return
	}

	for {
		var request JSONRPCrequest
		if err := conn.ReadJSON(&request); err != nil {
			log.Println(err)
			return
		}

		if err := processRequest(&request, conn); err != nil {
			log.Println(err)
		}
	}
}

func processRequest(request *JSONRPCrequest, conn *websocket.Conn) error {

	var params MyParams

	// extract Message
	message, ok := request.Params["Message"]
	if !ok {
		return errorMyParams{"Error while converting Message"}
	}
	params.Message = message.(string)

	// Extract A
	preA, ok := request.Params["A"].([]interface{})
	if !ok {
		return errorMyParams{"Error while converting A"}
	}
	params.A = make([]string, len(preA))
	for key, value := range preA {
		conv, ok := value.(string)
		if !ok {
			return errorMyParams{"Error while converting A member"}
		}

		params.A[key] = conv
	}

	// Echo
	return conn.WriteJSON(request)
}
