package arca

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var connections []*websocket.Conn
var subscriptions map[*websocket.Conn][]string
var handlers requestHandlers
var upgradeConnection func(w http.ResponseWriter,
	r *http.Request) (*websocket.Conn, error)
var readJSON func(conn *websocket.Conn, request *JSONRPCrequest) error
var writeJSON func(conn *websocket.Conn, response *JSONRPCresponse) error
var closeConnection func(conn *websocket.Conn) error

var upgrader = websocket.Upgrader{
	ReadBufferSize:  64,
	WriteBufferSize: 64,
}

func setupGlobals() {
	connections = []*websocket.Conn{}
	subscriptions = map[*websocket.Conn][]string{}
	handlers = requestHandlers{}
	upgradeConnection = func(w http.ResponseWriter,
		r *http.Request) (*websocket.Conn, error) {
		return upgrader.Upgrade(w, r, nil)
	}
	readJSON = func(conn *websocket.Conn, request *JSONRPCrequest) error {
		return conn.ReadJSON(&request)
	}
	writeJSON = func(conn *websocket.Conn, response *JSONRPCresponse) error {
		return conn.WriteJSON(response)
	}
	closeConnection = func(conn *websocket.Conn) error {
		return conn.Close()
	}
}

func init() {
	setupGlobals()
}
