package example

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	grid "github.com/rianby64/arca-grid"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

// BindViewTable2WithPg whatever
func BindViewTable2WithPg(
	s *arca.JSONRPCExtensionWS,
	db *sql.DB,
) *grid.Grid {

	type ViewTable2 struct {
		ID   int64
		Num3 float64
		Num4 float64
		I    int64
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
			"Num4",
			"I"
		FROM "ViewTable2"
		ORDER BY "ID"
		`)
		if err != nil {
			log.Fatal(err)
		}

		var results []ViewTable2

		var iID interface{}
		var iNum3 interface{}
		var iNum4 interface{}
		var iI interface{}

		for rows.Next() {
			err := rows.Scan(
				&iID,
				&iNum3,
				&iNum4,
				&iI,
			)
			if err != nil {
				log.Fatal(err)
			}

			var ID int64
			var Num3 float64
			var Num4 float64
			var I int64

			if iID != nil {
				ID = iID.(int64)
			}
			if iNum3 != nil {
				Num3 = iNum3.(float64)
			}
			if iNum4 != nil {
				Num4 = iNum4.(float64)
			}
			if iI != nil {
				I = iI.(int64)
			}

			results = append(results, ViewTable2{
				ID:   ID,
				Num3: Num3,
				Num4: Num4,
				I:    I,
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
			if key == "Num3" ||
				key == "Num4" {
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
		UPDATE "ViewTable2"
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
			if key == "Num3" || key == "Num4" {
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
		INSERT INTO "ViewTable2"(%v)
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
		DELETE FROM "ViewTable2"
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

	BindArcaWithGrid(s, &g, &methods, "ViewTable2")
	return &g
}
