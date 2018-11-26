package arca

import (
	"log"
	"net/http"
)

// Handle is for serving WS
func Handle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	defer removeConnection(conn)

	if err != nil {
		log.Println("connecting", err)
		return
	}

	appendConnection(conn)
	for {
		var request JSONRPCrequest
		if err := conn.ReadJSON(&request); err != nil {
			log.Println(err)
			return
		}
		processJSONRPCrequest(&request, conn)
	}
}
