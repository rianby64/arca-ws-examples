package arca

import (
	"github.com/gorilla/websocket"
)

var connections []*websocket.Conn
var subscriptions map[*websocket.Conn][]string
var handlers requestHandlers
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
