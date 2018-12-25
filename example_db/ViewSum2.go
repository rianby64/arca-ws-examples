package example

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	grid "github.com/rianby64/arca-grid"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

// BindViewSum2WithPg whatever
func BindViewSum2WithPg(
	s *arca.JSONRPCExtensionWS,
	connStr string,
	db *sql.DB,
) *grid.Grid {

	type ViewSum2 struct {
		ID         string
		Table1ID   int64
		Table2ID   int64
		Table1Num2 float64
		Table2Num4 float64
		Sum24      float64
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
			"Table1Num2",
			"Table2Num4",
			"Sum24"
		FROM "ViewSum2"
		`)
		if err != nil {
			log.Fatal(err)
		}

		var results []ViewSum2

		var iID interface{}
		var iTable1ID interface{}
		var iTable2ID interface{}
		var iTable1Num2 interface{}
		var iTable2Num4 interface{}
		var iSum24 interface{}

		for rows.Next() {
			err := rows.Scan(
				&iID,
				&iTable1ID,
				&iTable2ID,
				&iTable1Num2,
				&iTable2Num4,
				&iSum24,
			)
			if err != nil {
				log.Fatal(err)
			}

			var ID string
			var Table1ID int64
			var Table2ID int64
			var Table1Num2 float64
			var Table2Num4 float64
			var Sum24 float64

			if iID != nil {
				ID = iID.(string)
			}
			if iTable1ID != nil {
				Table1ID = iTable1ID.(int64)
			}
			if iTable2ID != nil {
				Table2ID = iTable2ID.(int64)
			}
			if iTable1Num2 != nil {
				Table1Num2 = iTable1Num2.(float64)
			}
			if iTable2Num4 != nil {
				Table2Num4 = iTable2Num4.(float64)
			}
			if iSum24 != nil {
				Sum24 = iSum24.(float64)
			}

			results = append(results, ViewSum2{
				ID:         ID,
				Table1ID:   Table1ID,
				Table2ID:   Table2ID,
				Table1Num2: Table1Num2,
				Table2Num4: Table2Num4,
				Sum24:      Sum24,
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
				key == "Sum24" {
				continue
			}
			if key == "Table1Num2" ||
				key == "Table2Num4" {
				Value := value.(float64)
				setters = append(setters, fmt.Sprintf(`"%v"=%v`, key, Value))
			}
		}
		strSetters := strings.Join(setters, ",")
		ID := params["ID"].(string)
		query := fmt.Sprintf(`
		UPDATE "ViewSum2"
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
			if key == "ID" ||
				key == "Table1ID" ||
				key == "Table2ID" ||
				key == "Sum24" {
				continue
			}
			fields = append(fields, fmt.Sprintf(`"%v"`, key))
			if key == "Table1Num2" || key == "Table2Num4" {
				Value := value.(float64)
				values = append(values, fmt.Sprintf(`%v`, Value))
			}
		}
		strValues := strings.Join(values, ",")
		strFields := strings.Join(fields, ",")
		query := fmt.Sprintf(`
		INSERT INTO "ViewSum2"(%v)
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
		ID := params["ID"].(string)

		query := fmt.Sprintf(`
		DELETE FROM "ViewSum2"
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

	BindArcaWithGrid(connStr, s, &g, &methods, "ViewSum2")
	return &g
}
