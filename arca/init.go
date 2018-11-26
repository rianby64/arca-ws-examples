package arca

import (
	"github.com/gorilla/websocket"
)

var conns []*websocket.Conn
var subscriptions map[*websocket.Conn][]string
var handlers requestHandlers

var upgrader = websocket.Upgrader{
	ReadBufferSize:  64,
	WriteBufferSize: 64,
}

func setupGlobals() {
	conns = []*websocket.Conn{}
	subscriptions = map[*websocket.Conn][]string{}
	handlers = requestHandlers{}
}

func init() {
	setupGlobals()
}
