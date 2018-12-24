package example

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	grid "github.com/rianby64/arca-grid"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

// BindTable1WithPg whatever
func BindTable1WithPg(
	s *arca.JSONRPCExtensionWS,
	connStr string,
	db *sql.DB,
) *grid.Grid {

	type AAURow struct {
		ID        int64
		Num1      float64
		Num2      float64
		CreatedAt time.Time
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
			"Num1",
			"Num2",
			"CreatedAt"
		FROM "Table1"
		`)
		if err != nil {
			log.Fatal(err)
		}

		var results []AAURow

		var iID interface{}
		var iNum1 interface{}
		var iNum2 interface{}
		var iCreatedAt interface{}

		for rows.Next() {
			err := rows.Scan(
				&iID,
				&iNum1,
				&iNum2,
				&iCreatedAt,
			)
			if err != nil {
				log.Fatal(err)
			}

			var ID int64
			var Num1 float64
			var Num2 float64
			var CreatedAt time.Time

			if iID != nil {
				ID = iID.(int64)
			}
			if iNum1 != nil {
				Num1 = iNum1.(float64)
			}
			if iNum2 != nil {
				Num2 = iNum2.(float64)
			}
			if iCreatedAt != nil {
				CreatedAt = iCreatedAt.(time.Time)
			}

			results = append(results, AAURow{
				ID:        ID,
				Num1:      Num1,
				Num2:      Num2,
				CreatedAt: CreatedAt,
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
			if key == "Num1" || key == "Num2" {
				Value := value.(float64)
				setters = append(setters, fmt.Sprintf(`"%v"=%v`, key, Value))
			}
		}
		strSetters := strings.Join(setters, ",")
		ID := params["ID"].(float64)
		query := fmt.Sprintf(`
		UPDATE "Table1"
			SET %v
			WHERE "ID"='%v';
		`, strSetters, ID)
		db.Exec(query)
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
			if key == "Num1" || key == "Num2" {
				Value := value.(float64)
				values = append(values, fmt.Sprintf(`%v`, Value))
			}
		}
		strValues := strings.Join(values, ",")
		strFields := strings.Join(fields, ",")
		query := fmt.Sprintf(`
		INSERT INTO "Table1"(%v)
			VALUES(%v);
		`, strFields, strValues)
		db.Exec(query)
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
		DELETE FROM "Table1"
			WHERE "ID"='%v';
		`, ID)
		db.Exec(query)
		return nil, nil
	}

	methods := grid.QUID{
		Query:  &queryHandler,
		Update: &updateHandler,
		Insert: &insertHandler,
		Delete: &deleteHandler,
	}

	go BindArcaWithGrid(connStr, s, &g, &methods, "Table1")
	return &g
}
