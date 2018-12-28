package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	example "./example_db"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

func main() {
	ws := arca.JSONRPCExtensionWS{}
	dbs := map[string]*sql.DB{}

	dbName1 := "arca-1"
	connStr1 := fmt.Sprintf(
		"user=arca password=arca dbname=%v sslmode=disable",
		dbName1)
	db1, err := sql.Open("postgres", connStr1)
	if err != nil {
		log.Fatal(err)
	}
	dbs[dbName1] = db1
	example.ConnectNotifyWithArca(connStr1, dbName1, dbName1, &ws, &dbs)

	dbName2 := "arca-2"
	connStr2 := fmt.Sprintf(
		"user=arca password=arca dbname=%v sslmode=disable",
		dbName2)
	db2, err := sql.Open("postgres", connStr2)
	if err != nil {
		log.Fatal(err)
	}
	dbs[dbName2] = db2
	example.ConnectNotifyWithArca(connStr2, dbName2, dbName1, &ws, &dbs)

	dbName3 := "arca-3"
	connStr3 := fmt.Sprintf(
		"user=arca password=arca dbname=%v sslmode=disable",
		dbName3)
	db3, err := sql.Open("postgres", connStr3)
	if err != nil {
		log.Fatal(err)
	}
	dbs[dbName3] = db3
	example.ConnectNotifyWithArca(connStr3, dbName3, dbName1, &ws, &dbs)

	example.BindTable1WithPg(&ws, connStr1, dbName1, &dbs)
	example.BindTable2WithPg(&ws, connStr1, dbName1, &dbs)
	example.BindViewSum1WithPg(&ws, connStr2, dbName2, &dbs)
	example.BindViewSum2WithPg(&ws, connStr3, dbName3, &dbs)

	http.HandleFunc("/ws", ws.Handle)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Serving")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
