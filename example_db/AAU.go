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

// BindServerWithPg whatever
func BindServerWithPg(
	s *arca.JSONRPCExtensionWS,
	connStr string,
	db *sql.DB,
) *grid.Grid {

	type AAURow struct {
		ID          string
		Parent      string
		Expand      bool
		Description string
		Information string
		Unit        string
		Qop         float64
		Estimated   float64
		CreatedAt   time.Time
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
			"Parent",
			"Expand",
			"Description",
			"Information",
			"Unit",
			"Qop",
			"Estimated",
			"CreatedAt"
		FROM "AAU"
		ORDER BY normalize_keynote("ID")
		`)
		if err != nil {
			log.Fatal(err)
		}

		var results []AAURow

		var iID interface{}
		var iParent interface{}
		var iExpand interface{}
		var iDescription interface{}
		var iInformation interface{}
		var iUnit interface{}
		var iQop interface{}
		var iEstimated interface{}
		var iCreatedAt interface{}

		for rows.Next() {
			err := rows.Scan(
				&iID,
				&iParent,
				&iExpand,
				&iDescription,
				&iInformation,
				&iUnit,
				&iQop,
				&iEstimated,
				&iCreatedAt,
			)
			if err != nil {
				log.Fatal(err)
			}

			var ID string
			var Parent string
			var Expand bool
			var Description string
			var Information string
			var Unit string
			var Qop float64
			var Estimated float64
			var CreatedAt time.Time

			if iID != nil {
				ID = iID.(string)
			}
			if iParent != nil {
				Parent = iParent.(string)
			}
			if iExpand != nil {
				Expand = iExpand.(bool)
			}
			if iDescription != nil {
				Description = iDescription.(string)
			}
			if iInformation != nil {
				Information = iInformation.(string)
			}
			if iUnit != nil {
				Unit = iUnit.(string)
			}
			if iQop != nil {
				Qop = iQop.(float64)
			}
			if iEstimated != nil {
				Estimated = iEstimated.(float64)
			}
			if iCreatedAt != nil {
				CreatedAt = iCreatedAt.(time.Time)
			}

			results = append(results, AAURow{
				ID:          ID,
				Parent:      Parent,
				Expand:      Expand,
				Description: Description,
				Information: Information,
				Unit:        Unit,
				Qop:         Qop,
				Estimated:   Estimated,
				CreatedAt:   CreatedAt,
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
			if key == "Parent" ||
				key == "Description" {
				Value := value.(string)
				setters = append(setters, fmt.Sprintf(`"%v"='%v'`, key, Value))
			}
			if key == "Expand" {
				Value := value.(bool)
				setters = append(setters, fmt.Sprintf(`"%v"=%v`, key, Value))
			}
			if key == "Qop" {
				Value := value.(float64)
				setters = append(setters, fmt.Sprintf(`"%v"=%v`, key, Value))
			}
		}
		strSetters := strings.Join(setters, ",")
		ID := params["ID"].(string)
		query := fmt.Sprintf(`
		UPDATE "AAU"
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
		/*
			INSERT INTO test("Name", "Email")
			VALUES ($1, $2);
		*/
		params := (*requestParams).(map[string]interface{})
		fields := []string{}
		values := []string{}
		for key, value := range params {
			if key == "ID" {
				continue
			}
			fields = append(fields, fmt.Sprintf(`"%v"`, key))
			if key == "Parent" ||
				key == "Description" {
				Value := value.(string)
				values = append(values, fmt.Sprintf(`'%v'`, Value))
			}
			if key == "Expand" {
				Value := value.(bool)
				values = append(values, fmt.Sprintf(`%v`, Value))
			}
			if key == "Qop" {
				Value := value.(float64)
				values = append(values, fmt.Sprintf(`%v`, Value))
			}
		}
		strValues := strings.Join(values, ",")
		strFields := strings.Join(fields, ",")
		query := fmt.Sprintf(`
		INSERT INTO "AAU"(%v)
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
		DELETE FROM "AAU"
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

	go BindArcaWithGrid(connStr, s, &g, &methods, "AAU")
	return &g
}
