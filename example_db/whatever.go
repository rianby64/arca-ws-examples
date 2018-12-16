package example

import (
	"encoding/json"
	"log"
	"time"

	"github.com/lib/pq"
	arca "github.com/rianby64/arca-ws-jsonrpc"
)

type pgNotifyJSONRPC struct {
	Method string
	Source string
	Result interface{}
}

// Want to Extend the package arca-ws-jsonrpc with the function below...

func listenToPgNotifyToArca(connStr string, s *arca.JSONRPCServerWS) {
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
