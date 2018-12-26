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

	mirrors := []*sql.DB{}

	connStr1 := "user=arca password=arca dbname=arca-1 sslmode=disable"
	db1, err := sql.Open("postgres", connStr1)
	if err != nil {
		log.Fatal(err)
	}
	example.BindTable1WithPg(&ws, connStr1, db1, &mirrors)
	example.BindTable2WithPg(&ws, connStr1, db1, &mirrors)

	connStr2 := "user=arca password=arca dbname=arca-2 sslmode=disable"
	db2, err := sql.Open("postgres", connStr2)
	if err != nil {
		log.Fatal(err)
	}
	example.BindViewSum1WithPg(&ws, connStr2, db2)
	mirrors = append(mirrors, db2)

	connStr3 := "user=arca password=arca dbname=arca-3 sslmode=disable"
	db3, err := sql.Open("postgres", connStr3)
	if err != nil {
		log.Fatal(err)
	}
	example.BindViewSum2WithPg(&ws, connStr3, db3)
	mirrors = append(mirrors, db3)

	http.HandleFunc("/ws", ws.Handle)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Serving")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
