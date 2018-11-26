package arca

import (
	"github.com/gorilla/websocket"
)

var conns []*websocket.Conn
var subscriptions map[*websocket.Conn][]string
var handlers requestHandlers
var writeJSON func(conn *websocket.Conn, response *JSONRPCresponse) error

var upgrader = websocket.Upgrader{
	ReadBufferSize:  64,
	WriteBufferSize: 64,
}

func setupGlobals() {
	conns = []*websocket.Conn{}
	subscriptions = map[*websocket.Conn][]string{}
	handlers = requestHandlers{}
	writeJSON = func(conn *websocket.Conn, response *JSONRPCresponse) error {
		return conn.WriteJSON(response)
	}
}

func init() {
	setupGlobals()
}
