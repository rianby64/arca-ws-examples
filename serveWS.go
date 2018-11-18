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
		log.Println("begin loop")
		var request JSONRPCrequest
		err := conn.ReadJSON(&request)
		if err != nil {
			log.Println("Error", err)
			return
		}

		processRequest(&request, conn)

		log.Println("end loop")
	}
}

func processRequest(request *JSONRPCrequest, conn *websocket.Conn) {

	log.Println(request.ID, request.Jsonrpc, request.Method, request.Params)
	var params MyParams

	preA := request.Params["A"].([]interface{})

	params.Message = request.Params["Message"].(string)
	params.A = make([]string, len(preA))

	for key, value := range preA {
		params.A[key] = value.(string)
	}

	err := conn.WriteJSON(request)
	if err != nil {
		log.Println("sending", err)
	}
}
