package example

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	// for db
	"github.com/lib/pq"
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

	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Fatalln(err)
		}
	}

	minReconn := 10 * time.Second
	maxReconn := time.Minute
	listener := pq.NewListener(connStr, minReconn, maxReconn, reportProblem)
	err = listener.Listen("jsonrpc")
	if err != nil {
		panic(err)
	}

	type pgNotifyJSONRPC struct {
		Method string
		Source string
		Result interface{}
	}

	go (func() {
		for {
			msg, ok := <-listener.Notify
			if !ok {
				return
			}
			var notification pgNotifyJSONRPC
			payload := []byte(msg.Extra)
			json.Unmarshal(payload, &notification)

			var context interface{} = map[string]string{
				"source": notification.Source,
			}
			var response arca.JSONRPCresponse

			response.Method = notification.Method
			response.Context = context
			response.Result = notification.Result

			s.Broadcast(&response)
		}
	})()

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
	g.RegisterMethod("query", &queryHandler)

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
	g.RegisterMethod("update", &updateHandler)

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
	g.RegisterMethod("insert", &insertHandler)

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
	g.RegisterMethod("delete", &deleteHandler)

	var insertMethod arca.JSONRequestHandler = g.Insert
	s.RegisterMethod("test", "insert", &insertMethod)
	var queryMethod arca.JSONRequestHandler = g.Query
	s.RegisterMethod("test", "read", &queryMethod)
	var updateMethod arca.JSONRequestHandler = g.Update
	s.RegisterMethod("test", "update", &updateMethod)
	var deleteMethod arca.JSONRequestHandler = g.Delete
	s.RegisterMethod("test", "delete", &deleteMethod)

	return &g
}
