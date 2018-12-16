package example

import (
	"database/sql"
	"log"

	grid "github.com/rianby64/arca-grid"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

// BindServerWithPg whatever
func BindServerWithPg(
	s *arca.JSONRPCServerWS,
	connStr string,
	db *sql.DB,
) *grid.Grid {

	type testRow struct {
		ID    int64
		Name  string
		Email string
	}

	g := grid.Grid{}

	var queryHandler grid.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify grid.NotifyCallback,
	) (interface{}, error) {
		rows, err := db.Query(`
		SELECT "ID", "Name", "Email" FROM test
		`)
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
			var Name string
			var Email string
			if name != nil {
				Name = name.(string)
			}
			if email != nil {
				Email = email.(string)
			}
			results = append(results, testRow{
				ID:    id.(int64),
				Name:  Name,
				Email: Email,
			})
		}

		rows.Close()
		return results, nil
	}

	var updateHandler grid.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify grid.NotifyCallback,
	) (interface{}, error) {
		params := (*requestParams).(map[string]interface{})
		name, okName := params["Name"]
		email, okEmail := params["Email"]
		ID := params["ID"].(float64)

		if okName && okEmail {
			db.Exec(`
			UPDATE test
			SET "Name"=$1, "Email"=$2
			WHERE "ID"=$3;
			`, name, email, ID)
		} else if okName {
			db.Exec(`
			UPDATE test
			SET "Name"=$1
			WHERE "ID"=$2;
			`, name, ID)
		} else if okEmail {
			db.Exec(`
			UPDATE test
			SET "Email"=$1
			WHERE "ID"=$2;
			`, email, ID)
		}
		return nil, nil
	}

	var insertHandler grid.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify grid.NotifyCallback,
	) (interface{}, error) {
		params := (*requestParams).(map[string]interface{})
		name, okName := params["Name"]
		email, okEmail := params["Email"]

		if okName && okEmail {
			db.Exec(`
			INSERT INTO test("Name", "Email")
			VALUES ($1, $2);
			`, name, email)
		} else if okName {
			db.Exec(`
			INSERT INTO test("Name")
			VALUES ($1);
			`, name)
		} else if okEmail {
			db.Exec(`
			INSERT INTO test("Email")
			VALUES ($1);
			`, email)
		}
		return nil, nil
	}

	var deleteHandler grid.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify grid.NotifyCallback,
	) (interface{}, error) {
		params := (*requestParams).(map[string]interface{})
		ID := params["ID"].(float64)

		db.Exec(`
		DELETE FROM test
		WHERE "ID"=$1;
		`, ID)
		return nil, nil
	}

	methods := grid.QUID{
		Query:  &queryHandler,
		Update: &updateHandler,
		Insert: &insertHandler,
		Delete: &deleteHandler,
	}

	go BindArcaWithGrid(connStr, s, &g, &methods)
	return &g
}
