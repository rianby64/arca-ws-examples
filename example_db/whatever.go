package example

import (
	"encoding/json"
	"log"
	"time"

	"github.com/lib/pq"
	grid "github.com/rianby64/arca-grid"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

// BindArcaWithGrid whatever
func BindArcaWithGrid(
	connStr string,
	s *arca.JSONRPCServerWS,
	g *grid.Grid,
	methods *grid.QUID) {

	var queryMethod arca.JSONRequestHandler = g.Query
	var updateMethod arca.JSONRequestHandler = g.Update
	var insertMethod arca.JSONRequestHandler = g.Insert
	var deleteMethod arca.JSONRequestHandler = g.Delete

	s.RegisterMethod("test", "read", &queryMethod)
	s.RegisterMethod("test", "update", &updateMethod)
	s.RegisterMethod("test", "insert", &insertMethod)
	s.RegisterMethod("test", "delete", &deleteMethod)

	g.Register(methods)

	type pgNotifyJSONRPC struct {
		Method string
		Source string
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

	for {
		msg, ok := <-listener.Notify
		if !ok {
			return
		}
		var notification pgNotifyJSONRPC
		payload := []byte(msg.Extra)
		json.Unmarshal(payload, &notification)

		var context interface{} = map[string]string{
			"source": notification.Source,
		}
		var response arca.JSONRPCresponse

		response.Method = notification.Method
		response.Context = context
		response.Result = notification.Result

		s.Broadcast(&response)
	}
}
