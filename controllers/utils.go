package example

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/lib/pq"
	grid "github.com/rianby64/arca-grid"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

// BindArcaWithGrid whatever
func BindArcaWithGrid(
	s *arca.JSONRPCExtensionWS,
	g *grid.Grid,
	methods *grid.QUID,
	source string,
) {

	var queryMethod arca.JSONRequestHandler = g.Query
	var updateMethod arca.JSONRequestHandler = g.Update
	var insertMethod arca.JSONRequestHandler = g.Insert
	var deleteMethod arca.JSONRequestHandler = g.Delete

	s.RegisterMethod(source, "read", &queryMethod)
	s.RegisterMethod(source, "update", &updateMethod)
	s.RegisterMethod(source, "insert", &insertMethod)
	s.RegisterMethod(source, "delete", &deleteMethod)

	g.Register(methods)
}

// ConnectNotifyWithArca whatever
func ConnectNotifyWithArca(
	connStr string,
	dbName string,
	s *arca.JSONRPCExtensionWS,
	dbs *map[string]*sql.DB,
) {

	type pgNotifyJSONRPC struct {
		Method  string
		Source  string
		Primary bool
		View    bool
		Result  interface{}
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
				log.Println("Disconnected", dbName)
				return
			}
			var notification pgNotifyJSONRPC
			payload := []byte(msg.Extra)

			err := json.Unmarshal(payload, &notification)
			if err != nil {
				log.Fatalln(payload, "Cant unmarshal it")
			}

			var context interface{} = map[string]interface{}{
				"Source":  notification.Source,
				"Primary": notification.Primary,
				"View":    notification.View,
			}
			var response arca.JSONRPCresponse

			response.Method = notification.Method
			response.Context = context
			response.Result = notification.Result

			if notification.Primary {
				log.Println("primary ::", notification)
				for dbNameContext := range *dbs {
					request := arca.JSONRPCrequest{}
					request.Method = notification.Method
					request.Context = map[string]interface{}{
						"Source": notification.Source,
						"Db":     dbNameContext,
					}
					request.Params = notification.Result
					log.Println("request proc ::", request)
					s.ProcessRequest(&request)
				}
				log.Println("")
				continue
			}

			if notification.View {
				log.Println("view ::", notification)
				request := arca.JSONRPCrequest{}
				request.Method = notification.Method
				request.Context = map[string]interface{}{
					"Source": notification.Source,
				}
				request.Params = notification.Result
				log.Println("request proc ::", request)
				s.ProcessRequest(&request)
				log.Println("")
				continue
			}

			go s.Broadcast(&response)
		}
	})()
}
