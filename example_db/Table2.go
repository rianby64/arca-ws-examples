package example

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/lib/pq"
	grid "github.com/rianby64/arca-grid"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

// BindTable2WithPg whatever
func BindTable2WithPg(
	s *arca.JSONRPCExtensionWS,
	connStr string,
	db *sql.DB,
	dbName string,
	mirrors *[]*sql.DB,
) *grid.Grid {

	type Table2 struct {
		ID   int64
		Num3 float64
		Num4 float64
	}

	g := grid.Grid{}

	var queryHandler grid.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify grid.NotifyCallback,
	) (interface{}, error) {
		rows, err := db.Query(`
		SELECT
			"ID",
			"Num3",
			"Num4"
		FROM "Table2"
		ORDER BY "ID"
		`)
		if err != nil {
			log.Fatal(err)
		}

		var results []Table2

		var iID interface{}
		var iNum3 interface{}
		var iNum4 interface{}

		for rows.Next() {
			err := rows.Scan(
				&iID,
				&iNum3,
				&iNum4,
			)
			if err != nil {
				log.Fatal(err)
			}

			var ID int64
			var Num3 float64
			var Num4 float64

			if iID != nil {
				ID = iID.(int64)
			}
			if iNum3 != nil {
				Num3 = iNum3.(float64)
			}
			if iNum4 != nil {
				Num4 = iNum4.(float64)
			}

			results = append(results, Table2{
				ID:   ID,
				Num3: Num3,
				Num4: Num4,
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
		setters := []string{}
		for key, value := range params {
			if key == "ID" {
				continue
			}
			if key == "Num3" || key == "Num4" {
				Value := value.(float64)
				setters = append(setters, fmt.Sprintf(`"%v"=%v`, key, Value))
			}
		}
		strSetters := strings.Join(setters, ",")
		ID := params["ID"].(float64)
		query := fmt.Sprintf(`
		UPDATE "Table2"
			SET %v
			WHERE "ID"='%v';
		`, strSetters, ID)
		_, err := db.Exec(query)
		if err != nil {
			log.Println(err)
		}
		for _, mirror := range *mirrors {
			_, err := mirror.Exec(query)
			if err != nil {
				log.Println(err)
			}
		}
		return nil, nil
	}

	var insertHandler grid.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify grid.NotifyCallback,
	) (interface{}, error) {
		params := (*requestParams).(map[string]interface{})
		fields := []string{}
		values := []string{}
		for key, value := range params {
			if key == "ID" {
				continue
			}
			fields = append(fields, fmt.Sprintf(`"%v"`, key))
			if key == "Num3" || key == "Num4" {
				Value := value.(float64)
				values = append(values, fmt.Sprintf(`%v`, Value))
			}
		}
		strValues := strings.Join(values, ",")
		strFields := strings.Join(fields, ",")
		query := fmt.Sprintf(`
		INSERT INTO "Table2"(%v)
			VALUES(%v);
		`, strFields, strValues)
		_, err := db.Exec(query)
		if err != nil {
			log.Println(err)
		}
		for _, mirror := range *mirrors {
			_, err := mirror.Exec(query)
			if err != nil {
				log.Println(err)
			}
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

		query := fmt.Sprintf(`
		DELETE FROM "Table2"
			WHERE "ID"='%v';
		`, ID)
		_, err := db.Exec(query)
		if err != nil {
			log.Println(err)
		}
		for _, mirror := range *mirrors {
			_, err := mirror.Exec(query)
			if err != nil {
				log.Println(err)
			}
		}
		return nil, nil
	}

	methods := grid.QUID{
		Query:  &queryHandler,
		Update: &updateHandler,
		Insert: &insertHandler,
		Delete: &deleteHandler,
	}

	BindArcaWithGrid(connStr, s, &g, &methods, "Table2")

	/*
		Here I have to handle the way to redirect the notifications...
		Still I'm exploring
	*/

	type pgNotifyJSONRPC struct {
		Method string
		Source string
		Db     string
		Result interface{}
	}

	reportProblem := func(_ pq.ListenerEventType, err error) {
		if err != nil {
			log.Fatalln(err)
		}
	}

	minReconn := 10 * time.Second
	maxReconn := time.Minute
	listener := pq.NewListener(connStr, minReconn, maxReconn, reportProblem)
	err := listener.Listen("jsonrpc")
	if err != nil {
		panic(err)
	}

	go (func() {
		for {
			msg, ok := <-listener.Notify
			if !ok {
				return
			}
			var notification pgNotifyJSONRPC
			payload := []byte(msg.Extra)

			err := json.Unmarshal(payload, &notification)
			if err != nil {
				log.Println(err, ":: Notification ERROR")
			}

			var context interface{} = map[string]string{
				"Source": notification.Source,
				"Db":     notification.Db,
			}
			var response arca.JSONRPCresponse

			response.Method = notification.Method
			response.Context = context
			response.Result = notification.Result

			log.Println(dbName, response, ":: Notification")

			s.Broadcast(&response)
		}
	})()
	return &g
}
