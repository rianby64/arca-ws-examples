package example

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	grid "github.com/rianby64/arca-grid"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

// BindViewSum1WithPg whatever
func BindViewSum1WithPg(
	s *arca.JSONRPCExtensionWS,
	connStr string,
	db *sql.DB,
) *grid.Grid {

	type ViewSum1 struct {
		ID         string
		Table1ID   int64
		Table2ID   int64
		Table1Num1 float64
		Table2Num3 float64
		Sum13      float64
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
			"Table1ID",
			"Table2ID",
			"Table1Num1",
			"Table2Num3",
			"Sum13"
		FROM "ViewSum1"
		`)
		if err != nil {
			log.Fatal(err)
		}

		var results []ViewSum1

		var iID interface{}
		var iTable1ID interface{}
		var iTable2ID interface{}
		var iTable1Num1 interface{}
		var iTable2Num3 interface{}
		var iSum13 interface{}

		for rows.Next() {
			err := rows.Scan(
				&iID,
				&iTable1ID,
				&iTable2ID,
				&iTable1Num1,
				&iTable2Num3,
				&iSum13,
			)
			if err != nil {
				log.Fatal(err)
			}

			var ID string
			var Table1ID int64
			var Table2ID int64
			var Table1Num1 float64
			var Table2Num3 float64
			var Sum13 float64

			if iID != nil {
				ID = iID.(string)
			}
			if iTable1ID != nil {
				Table1ID = iTable1ID.(int64)
			}
			if iTable2ID != nil {
				Table2ID = iTable2ID.(int64)
			}
			if iTable1Num1 != nil {
				Table1Num1 = iTable1Num1.(float64)
			}
			if iTable2Num3 != nil {
				Table2Num3 = iTable2Num3.(float64)
			}
			if iSum13 != nil {
				Sum13 = iSum13.(float64)
			}

			results = append(results, ViewSum1{
				ID:         ID,
				Table1ID:   Table1ID,
				Table2ID:   Table2ID,
				Table1Num1: Table1Num1,
				Table2Num3: Table2Num3,
				Sum13:      Sum13,
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
			if key == "ID" ||
				key == "Table1ID" ||
				key == "Table2ID" ||
				key == "Sum13" {
				continue
			}
			if key == "Table1Num1" ||
				key == "Table2Num3" {
				Value := value.(float64)
				setters = append(setters, fmt.Sprintf(`"%v"=%v`, key, Value))
			}
		}
		strSetters := strings.Join(setters, ",")
		ID := params["ID"].(string)
		query := fmt.Sprintf(`
		UPDATE "ViewSum1"
			SET %v
			WHERE "ID"='%v';
		`, strSetters, ID)
		_, err := db.Exec(query)
		if err != nil {
			log.Println(err)
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
			if key == "ID" ||
				key == "Table1ID" ||
				key == "Table2ID" ||
				key == "Sum13" {
				continue
			}
			fields = append(fields, fmt.Sprintf(`"%v"`, key))
			if key == "Table1Num1" || key == "Table2Num3" {
				Value := value.(float64)
				values = append(values, fmt.Sprintf(`%v`, Value))
			}
		}
		strValues := strings.Join(values, ",")
		strFields := strings.Join(fields, ",")
		query := fmt.Sprintf(`
		INSERT INTO "ViewSum1"(%v)
			VALUES(%v);
		`, strFields, strValues)
		_, err := db.Exec(query)
		if err != nil {
			log.Println(err)
		}
		return nil, nil
	}

	var deleteHandler grid.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify grid.NotifyCallback,
	) (interface{}, error) {
		params := (*requestParams).(map[string]interface{})
		ID := params["ID"].(string)

		query := fmt.Sprintf(`
		DELETE FROM "ViewSum1"
			WHERE "ID"='%v';
		`, ID)
		_, err := db.Exec(query)
		if err != nil {
			log.Println(err)
		}
		return nil, nil
	}

	methods := grid.QUID{
		Query:  &queryHandler,
		Update: &updateHandler,
		Insert: &insertHandler,
		Delete: &deleteHandler,
	}

	BindArcaWithGrid(connStr, s, &g, &methods, "ViewSum1")
	return &g
}
