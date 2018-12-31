package example

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	grid "github.com/rianby64/arca-grid"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

// BindViewTable1WithPg whatever
func BindViewTable1WithPg(
	s *arca.JSONRPCExtensionWS,
	db *sql.DB,
) *grid.Grid {

	type ViewTable1 struct {
		ID   int64
		I    int64
		Num1 float64
		Num2 float64
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
			"I",
			"Num1",
			"Num2"
		FROM "ViewTable1"
		ORDER BY "ID"
		`)
		if err != nil {
			log.Fatal(err)
		}

		var results []ViewTable1

		var iID interface{}
		var iI interface{}
		var iNum1 interface{}
		var iNum2 interface{}

		for rows.Next() {
			err := rows.Scan(
				&iID,
				&iI,
				&iNum1,
				&iNum2,
			)
			if err != nil {
				log.Fatal(err)
			}

			var ID int64
			var I int64
			var Num1 float64
			var Num2 float64

			if iID != nil {
				ID = iID.(int64)
			}
			if iI != nil {
				I = iI.(int64)
			}
			if iNum1 != nil {
				Num1 = iNum1.(float64)
			}
			if iNum2 != nil {
				Num2 = iNum2.(float64)
			}

			results = append(results, ViewTable1{
				ID:   ID,
				I:    I,
				Num1: Num1,
				Num2: Num2,
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
			if key == "Num1" ||
				key == "Num2" {
				Value := value.(float64)
				setters = append(setters, fmt.Sprintf(`"%v"=%v`, key, Value))
			}
			if key == "I" {
				Value := int64(value.(float64))
				setters = append(setters, fmt.Sprintf(`"%v"=%v`, key, Value))
			}
		}
		strSetters := strings.Join(setters, ",")
		ID := int64(params["ID"].(float64))
		query := fmt.Sprintf(`
		UPDATE "ViewTable1"
			SET %v
			WHERE "ID"=%v;
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
			if key == "ID" {
				continue
			}
			fields = append(fields, fmt.Sprintf(`"%v"`, key))
			if key == "Num1" || key == "Num2" {
				Value := value.(float64)
				values = append(values, fmt.Sprintf(`%v`, Value))
			}
			if key == "I" {
				Value := int64(value.(float64))
				values = append(values, fmt.Sprintf(`%v`, Value))
			}
		}
		strValues := strings.Join(values, ",")
		strFields := strings.Join(fields, ",")
		query := fmt.Sprintf(`
		INSERT INTO "ViewTable1"(%v)
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
		ID := int64(params["ID"].(float64))

		query := fmt.Sprintf(`
		DELETE FROM "ViewTable1"
			WHERE "ID"=%v;
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

	BindArcaWithGrid(s, &g, &methods, "ViewTable1")
	return &g
}
