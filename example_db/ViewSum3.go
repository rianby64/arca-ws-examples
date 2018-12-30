package example

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	grid "github.com/rianby64/arca-grid"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

// BindViewSum3WithPg whatever
func BindViewSum3WithPg(
	s *arca.JSONRPCExtensionWS,
	db *sql.DB,
) *grid.Grid {

	type ViewSum3 struct {
		ID         string
		Table1ID   int64
		Table2ID   int64
		Table1I    int64
		Table2I    int64
		Table1Num1 float64
		Table1Num2 float64
		Table2Num3 float64
		Table2Num4 float64
		Sum1234    float64
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
			"Table1I",
			"Table2I",
			"Table1Num1",
			"Table1Num2",
			"Table2Num3",
			"Table2Num4",
			"Sum1234"
		FROM "ViewSum3"
		ORDER BY "Table1ID", "Table2ID"
		`)
		if err != nil {
			log.Fatal(err)
		}

		var results []ViewSum3

		var iID interface{}
		var iTable1ID interface{}
		var iTable2ID interface{}
		var iTable1I interface{}
		var iTable2I interface{}
		var iTable1Num1 interface{}
		var iTable1Num2 interface{}
		var iTable2Num3 interface{}
		var iTable2Num4 interface{}
		var iSum1234 interface{}

		for rows.Next() {
			err := rows.Scan(
				&iID,
				&iTable1ID,
				&iTable2ID,
				&iTable1I,
				&iTable2I,
				&iTable1Num1,
				&iTable1Num2,
				&iTable2Num3,
				&iTable2Num4,
				&iSum1234,
			)
			if err != nil {
				log.Fatal(err)
			}

			var ID string
			var Table1ID int64
			var Table2ID int64
			var Table1I int64
			var Table2I int64
			var Table1Num1 float64
			var Table1Num2 float64
			var Table2Num3 float64
			var Table2Num4 float64
			var Sum1234 float64

			if iID != nil {
				ID = iID.(string)
			}
			if iTable1ID != nil {
				Table1ID = iTable1ID.(int64)
			}
			if iTable2ID != nil {
				Table2ID = iTable2ID.(int64)
			}
			if iTable1I != nil {
				Table1I = iTable1I.(int64)
			}
			if iTable2I != nil {
				Table2I = iTable2I.(int64)
			}
			if iTable1Num1 != nil {
				Table1Num1 = iTable1Num1.(float64)
			}
			if iTable1Num2 != nil {
				Table1Num2 = iTable1Num2.(float64)
			}
			if iTable2Num3 != nil {
				Table2Num3 = iTable2Num3.(float64)
			}
			if iTable2Num4 != nil {
				Table2Num4 = iTable2Num4.(float64)
			}
			if iSum1234 != nil {
				Sum1234 = iSum1234.(float64)
			}

			results = append(results, ViewSum3{
				ID:         ID,
				Table1ID:   Table1ID,
				Table2ID:   Table2ID,
				Table1I:    Table1I,
				Table2I:    Table2I,
				Table1Num1: Table1Num1,
				Table1Num2: Table1Num2,
				Table2Num3: Table2Num3,
				Table2Num4: Table2Num4,
				Sum1234:    Sum1234,
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
				key == "Sum1234" || key == "CreatedAt" {
				continue
			}
			if key == "Table1Num1" ||
				key == "Table1Num2" ||
				key == "Table2Num3" ||
				key == "Table2Num4" {
				Value := value.(float64)
				setters = append(setters, fmt.Sprintf(`"%v"=%v`, key, Value))
			}
			if key == "Table1I" ||
				key == "Table2I" {
				Value := int64(value.(float64))
				setters = append(setters, fmt.Sprintf(`"%v"=%v`, key, Value))
			}
		}
		strSetters := strings.Join(setters, ",")
		ID := params["ID"].(string)
		query := fmt.Sprintf(`
		UPDATE "ViewSum3"
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
				key == "Sum1234" || key == "CreatedAt" {
				continue
			}
			fields = append(fields, fmt.Sprintf(`"%v"`, key))
			if key == "Table1Num1" ||
				key == "Table1Num2" ||
				key == "Table2Num3" ||
				key == "Table2Num4" {
				Value := value.(float64)
				values = append(values, fmt.Sprintf(`%v`, Value))
			}
			if key == "Table1I" ||
				key == "Table2I" {
				Value := int64(value.(float64))
				values = append(values, fmt.Sprintf(`%v`, Value))
			}
		}
		strValues := strings.Join(values, ",")
		strFields := strings.Join(fields, ",")
		query := fmt.Sprintf(`
		INSERT INTO "ViewSum3"(%v)
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
		DELETE FROM "ViewSum3"
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

	BindArcaWithGrid(s, &g, &methods, "ViewSum3")
	return &g
}
