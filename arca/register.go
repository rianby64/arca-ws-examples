package arca

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

var dummy requestHandler = func(requestParams *interface{}, context *interface{}) (interface{}, error) {
	return nil, fmt.Errorf("dummy executed with source %s",
		(*context).(map[string]interface{})["source"].(string))
}

// RegisterSource whatever
func RegisterSource(name string, methods *DIRUD) {
	handlers[name] = map[string]*requestHandler{
		"describe": &dummy,
		"insert":   &dummy,
		"read":     &dummy,
		"update":   &dummy,
		"delete":   &dummy,
	}
	log.Println(handlers[name])

	handler := handlers[name]
	if methods.Insert != nil {
		handler["insert"] = &methods.Insert
	}
	if methods.Read != nil {
		handler["read"] = &methods.Read
	}
	if methods.Update != nil {
		handler["update"] = &methods.Update
	}

	if methods.Delete != nil {
		handler["delete"] = &methods.Delete
	}
}

func appendConnection(conn *websocket.Conn) {
	connsTmp := make([]*websocket.Conn, len(connections)+1)
	for i, c := range connections {
		if c == conn {
			return
		}
		connsTmp[i] = c
	}
	connsTmp[len(connections)] = conn
	connections = connsTmp
}

func removeConnection(conn *websocket.Conn, closeConn bool) {
	for i, c := range connections {
		if c == conn {
			connections = append(connections[:i], connections[i+1:]...)
			break
		}
	}
	if _, ok := subscriptions[conn]; ok {
		delete(subscriptions, conn)
	}
	if closeConn {
		conn.Close()
	}
}

func subscribe(conn *websocket.Conn, source string) {
	found := false
	if list, ok := subscriptions[conn]; ok {
		for _, value := range list {
			if value == source {
				found = true
				break
			}
		}
	}
	if !found {
		subscriptions[conn] = append(subscriptions[conn], source)
	}
}

func unsubscribe(conn *websocket.Conn, source string) {
}
