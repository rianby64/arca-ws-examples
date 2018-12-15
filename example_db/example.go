package example

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // for db
)

// Start whatever
func Start() {
	log.Println("here we go with the db")
	connStr := "user=arca password=arca dbname=arca sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT id, name, email FROM test`)
	if err != nil {
		log.Fatal(err)
	}

	var id interface{}
	var name interface{}
	var email interface{}
	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, email)
	}

	rows.Close()
}
