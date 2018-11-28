package arca

import (
	"log"
	"net/http"
)

// Handle is for serving WS
func Handle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgradeConnection(w, r)
	defer removeConnection(conn)

	if err != nil {
		log.Println("connecting", err)
		return
	}

	appendConnection(conn)
	for {
		var request JSONRPCrequest
		if err := readJSON(conn, &request); err != nil {
			log.Println(err)
			return
		}
		for _, err := range processRequest(&request, conn) {
			if err != nil {
				log.Println(err)
			}
		}
	}
}
