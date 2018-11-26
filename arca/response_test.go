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
		t.Log("Reflect method from response")
	} else {
		t.Log(response.Method, request.Method)
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
	t.Log("Response to the sender if ID is given")
	request := JSONRPCrequest{}
	request.ID = "ID"
	request.Method = "method"
	request.Context = map[string]interface{}{"source": "context"}

	conn := websocket.Conn{}
	result := reflect.ValueOf(map[string]string{"result": "result"}).Interface()
	counter := 0

	writeJSON = func(_ *websocket.Conn, response *JSONRPCresponse) error {
		checkRequestResponse(t, &result, &request, response)
		if response.ID == request.ID {
			t.Log("Reflect ID from response")
		} else {
			t.Log(response.ID, request.ID)
			t.Error("ID must be reflected")
		}
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

func Test_response_one_request_broadcast(t *testing.T) {
	t.Log("Response via broadcast if ID is not given")
	source := "context"
	conn1 := &websocket.Conn{}
	conn2 := &websocket.Conn{}

	appendConnection(conn1)
	appendConnection(conn2)
	subscribe(conn1, source)
	subscribe(conn2, source)

	request := JSONRPCrequest{}
	request.ID = ""
	request.Method = "method"
	request.Context = map[string]interface{}{"source": source}

	result := reflect.ValueOf(map[string]string{"result": "result"}).Interface()
	counter := 0

	writeJSON = func(_ *websocket.Conn, response *JSONRPCresponse) error {
		checkRequestResponse(t, &result, &request, response)
		counter++
		return nil
	}

	errs := response(&request, &websocket.Conn{}, &result)
	if counter == 2 {
		t.Log("Two responses were send")
	} else {
		t.Error("No response was send", errs)
	}
	if len(errs) == 2 {
		if errs[0] == nil {
			t.Log("The response was send correctly")
		} else {
			t.Error(errs)
		}
	}
	setupGlobals()
}
