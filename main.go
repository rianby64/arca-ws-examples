package main

import (
	"database/sql"
	"log"
	"net/http"

	example "./example_db"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

func main() {
	ws := arca.JSONRPCExtensionWS{}
	connStr := "user=arca password=arca dbname=arca-1 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	example.BindServerWithPg(&ws, connStr, db)

	http.HandleFunc("/ws", ws.Handle)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Serving")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
