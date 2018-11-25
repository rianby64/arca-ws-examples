package arca

import (
	"github.com/gorilla/websocket"
)

var conns = []*websocket.Conn{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var handlers = requestHandlers{}
