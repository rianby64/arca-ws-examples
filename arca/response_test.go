package arca

import (
	"reflect"
	"testing"

	"github.com/gorilla/websocket"
)

func checkRequestResponse(t *testing.T, result *interface{},
	request *JSONRPCrequest, response *JSONRPCresponse) {
	if response.Context == &request.Context {
		t.Log("Reflect context from response")
	} else {
		t.Log(response.Context, request.Context)
		t.Error("Context must be reflected")
	}
	if response.Method == request.Method {
		t.Log(response.Method, request.Method)
		t.Log("Reflect method from response")
	} else {
		t.Error("Method must be reflected")
	}
	if response.Result == result {
		t.Log("Reflect result from response")
	} else {
		t.Log(response.Result, result)
		t.Error("Result must be reflected")
	}
}

func Test_response_one_request_with_ID(t *testing.T) {
	request := JSONRPCrequest{}
	request.ID = "ID"
	request.Method = "method"
	request.Context = map[string]string{"source": "context"}

	conn := websocket.Conn{}
	result := reflect.ValueOf(map[string]string{"result": "result"}).Interface()
	counter := 0

	writeJSON = func(conn *websocket.Conn, response *JSONRPCresponse) error {
		checkRequestResponse(t, &result, &request, response)
		counter++
		return nil
	}

	errs := response(&request, &conn, &result)
	if counter == 1 {
		t.Log("One response was send")
	} else {
		t.Error("No response was send")
	}
	if len(errs) == 1 {
		if errs[0] == nil {
			t.Log("The response was send correctly")
		} else {
			t.Error(errs)
		}
	}
	setupGlobals()
}
