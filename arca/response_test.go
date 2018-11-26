package arca

import (
	"reflect"
	"testing"

	"github.com/gorilla/websocket"
)

func Test_response_one_request_with_ID(t *testing.T) {
	request := JSONRPCrequest{}
	request.ID = "ID"
	request.Method = "method"
	request.Context = map[string]string{"source": "context"}

	conn := websocket.Conn{}
	result := reflect.ValueOf(map[string]string{"result": "result"}).Interface()
	counter := 0

	writeJSON = func(conn *websocket.Conn, response *JSONRPCresponse) error {
		if response.Context == &request.Context {
			t.Log("Reflect context from response")
		} else {
			t.Error("Context must be reflected")
		}
		if response.Method == request.Method {
			t.Log("Reflect method from response")
		} else {
			t.Error("Method must be reflected")
		}
		if response.Result == &result {
			t.Log("Reflect result from response")
		} else {
			t.Error("Result must be reflected")
		}
		counter++
		return nil
	}

	response(&request, &conn, &result)
	if counter == 1 {
		t.Log("One response was send")
	} else {
		t.Error("No response was send")
	}
	setupGlobals()
}
