package example

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	grid "github.com/rianby64/arca-grid"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

// BindTable1WithPg whatever
func BindTable1WithPg(
	s *arca.JSONRPCExtensionWS,
	dbs *map[string]*sql.DB,
) *grid.Grid {

	type Table1 struct {
		ID   int64
		I    int64
		Num1 float64
		Num2 float64
	}

	g := grid.Grid{}

	var updateHandler grid.RequestHandler = func(
		requestParams *interface{},
		context *interface{},
		notify grid.NotifyCallback,
	) (interface{}, error) {
		var db *sql.DB
		dbNameContext, ok := (*context).(map[string]interface{})["Db"]
		if ok {
			db = (*dbs)[dbNameContext.(string)]
		} else {
			log.Fatal("Update handler triggered without dbName context")
		}
		params := (*requestParams).(map[string]interface{})
		setters := []string{}
		for key, value := range params {
			if key == "ID" || key == "CreatedAt" {
				continue
			}
			if key == "Num1" || key == "Num2" {
				Value := value.(float64)
				setters = append(setters, fmt.Sprintf(`"%v"=%v`, key, Value))
			}
			if key == "I" {
				Value := int64(value.(float64))
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
		var db *sql.DB
		dbNameContext, ok := (*context).(map[string]interface{})["Db"]
		if ok {
			db = (*dbs)[dbNameContext.(string)]
		} else {
			log.Fatal("Insert handler triggered without dbName context")
		}
		params := (*requestParams).(map[string]interface{})
		fields := []string{}
		values := []string{}
		for key, value := range params {
			if key == "ID" || key == "CreatedAt" {
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
		INSERT INTO "Table1"(%v)
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
		var db *sql.DB
		dbNameContext, ok := (*context).(map[string]interface{})["Db"]
		if ok {
			db = (*dbs)[dbNameContext.(string)]
		} else {
			log.Fatal("Delete handler triggered without dbName context")
		}
		params := (*requestParams).(map[string]interface{})
		ID := params["ID"].(float64)

		query := fmt.Sprintf(`
		DELETE FROM "Table1"
			WHERE "ID"='%v';
		`, ID)
		_, err := db.Exec(query)
		if err != nil {
			log.Println(err)
		}
		return nil, nil
	}

	methods := grid.QUID{
		Query:  nil,
		Update: &updateHandler,
		Insert: &insertHandler,
		Delete: &deleteHandler,
	}

	BindArcaWithGrid(s, &g, &methods, "Table1")
	return &g
}
