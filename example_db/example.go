package example

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // for db
	grid "github.com/rianby64/arca-grid"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

// GridTest whatever
func GridTest(s *arca.JSONRPCServerWS) *grid.Grid {

	type testRow struct {
		ID    int64
		Name  string
		Email string
	}

	g := grid.Grid{}

	connStr := "user=arca password=arca dbname=arca sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	var queryHandler grid.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify grid.NotifyCallback,
	) (interface{}, error) {
		rows, err := db.Query(`SELECT id, name, email FROM test`)
		if err != nil {
			log.Fatal(err)
		}

		var results []testRow

		var id interface{}
		var name interface{}
		var email interface{}
		for rows.Next() {
			err := rows.Scan(&id, &name, &email)
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, testRow{
				ID:    id.(int64),
				Name:  name.(string),
				Email: email.(string),
			})
		}

		rows.Close()
		return results, nil
	}
	g.RegisterMethod("query", &queryHandler)

	var queryMethod arca.JSONRequestHandler = g.Query
	s.RegisterMethod("test", "read", &queryMethod)

	return &g
}
