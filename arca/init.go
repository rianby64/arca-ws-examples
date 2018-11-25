package arca

import (
	"github.com/gorilla/websocket"
)

var conns = []*websocket.Conn{}
var subscriptions = map[*websocket.Conn][]string{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var handlers = requestHandlers{}
