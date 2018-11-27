package arca

import (
	"testing"

	"github.com/gorilla/websocket"
)

func Test_matchHandler_request_without_Method(t *testing.T) {
	t.Log("Match a handler fails if no method defined in request")
	conn := &websocket.Conn{}
	request := &JSONRPCrequest{}

	handler, err := matchHandler(request, conn)
	t.Log(err)
	if handler == nil {
		if err == nil {
			t.Error("nil handler must lead to an error")
		}
	} else {
		t.Error("a Method must be defined at Handler")
	}
}

func Test_matchHandler_request_without_Context(t *testing.T) {
	t.Log("Match a handler fails if no context defined in request")
	conn := &websocket.Conn{}
	request := &JSONRPCrequest{}
	request.Method = "method"

	handler, err := matchHandler(request, conn)
	t.Log(err)
	if handler == nil {
		if err == nil {
			t.Error("nil handler must lead to an error")
		}
	} else {
		t.Error("a Context must be defined at Handler")
	}
}

func Test_matchHandler_request_with_incorrect_Context(t *testing.T) {
	t.Log("Match a handler fails if context defined in request is not an object")
	conn := &websocket.Conn{}
	request := &JSONRPCrequest{}
	request.Method = "method"
	request.Context = []string{}

	handler, err := matchHandler(request, conn)
	t.Log(err)
	if handler == nil {
		if err == nil {
			t.Error("nil handler must lead to an error")
		}
	} else {
		t.Error("a Context must be defined at Handler")
	}
}

func Test_matchHandler_request_with_Context_without_source(t *testing.T) {
	t.Log("Match a handler fails if context defined in request doesn't contain a source")
	conn := &websocket.Conn{}
	request := &JSONRPCrequest{}
	request.Method = "method"
	request.Context = map[string]interface{}{"whatever": ""}

	handler, err := matchHandler(request, conn)
	t.Log(err)
	if handler == nil {
		if err == nil {
			t.Error("nil handler must lead to an error")
		}
	} else {
		t.Error("The Context must contain a source")
	}
}
