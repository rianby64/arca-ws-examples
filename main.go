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

	/*
		connStr1 := "user=arca password=arca dbname=arca-1 sslmode=disable"
		db1, err := sql.Open("postgres", connStr1)
		if err != nil {
			log.Fatal(err)
		}
	*/

	connStr2 := "user=arca password=arca dbname=arca-2 sslmode=disable"
	db2, err := sql.Open("postgres", connStr2)
	if err != nil {
		log.Fatal(err)
	}

	example.BindTable1WithPg(&ws, connStr2, db2)
	example.BindTable2WithPg(&ws, connStr2, db2)
	example.BindViewSum1WithPg(&ws, connStr2, db2)

	http.HandleFunc("/ws", ws.Handle)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Serving")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
