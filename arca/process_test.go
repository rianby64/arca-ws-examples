package arca

import (
	"reflect"
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
	setupGlobals()
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
	setupGlobals()
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
	setupGlobals()
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
	setupGlobals()
}

func Test_matchHandler_request_with_Context_with_source(t *testing.T) {
	t.Log("Match a handler if context defined in request contains a source")
	conn := websocket.Conn{}
	source := "source-defined"
	method := "read"
	methods := DIRUD{
		Read: func(requestParams *interface{},
			context *interface{}) (interface{}, error) {
			return nil, nil
		},
	}
	RegisterSource(source, &methods)
	request := JSONRPCrequest{}
	request.Method = method
	request.Context = map[string]interface{}{"source": source}

	handler, err := matchHandler(&request, &conn)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	if handler == nil {
		t.Error("The Context must match the handler [%s][%s]", source, method)
		if err == nil {
			t.Error("nil handler must lead to an error")
		}
	} else {
		var ptHandler = reflect.ValueOf(*handler).Pointer()
		var ptMethod = reflect.ValueOf(methods.Read).Pointer()
		if ptHandler == ptMethod {
			t.Log("Matched handler and given method are the same")
		} else {
			t.Error("Matched handler differs from given method")
		}
	}
	setupGlobals()
}

func Test_matchHandler_request_to_subscribe(t *testing.T) {
	t.Log("Match the subscribe's handler")
	source := "source-defined"
	method := "subscribe"
	conn := websocket.Conn{}
	methods := DIRUD{}
	RegisterSource(source, &methods)
	request := JSONRPCrequest{}
	request.Method = method
	request.Context = map[string]interface{}{"source": source}

	handler, err := matchHandler(&request, &conn)
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
	if handler == nil {
		t.Errorf("The Context must match the handler [%s][%s]", source, method)
		if err == nil {
			t.Error("nil handler must lead to an error")
		}
	} else {
		(*handler)(&request.Params, &request.Context)
		subscribed, isConnSubscribed := subscriptions[&conn]
		if isConnSubscribed {
			item := subscribed[0]
			if item == source {
				t.Logf("%s in subscriptions", source)
			} else {
				t.Errorf("expecting to see subscribed item = %s", source)
			}
		} else {
			t.Errorf("expecting to see %s", source)
		}
	}
	setupGlobals()
}

func Test_matchHandler_request_to_unsubscribe(t *testing.T) {
	t.Log("Match the unsubscribe's handler")
	source := "source-defined"
	otherSource := "whatever-else"
	method := "unsubscribe"
	conn := websocket.Conn{}
	methods := DIRUD{}
	RegisterSource(source, &methods)
	subscribe(&conn, source)
	subscribe(&conn, otherSource)
	request := JSONRPCrequest{}
	request.Method = method
	request.Context = map[string]interface{}{"source": source}

	handler, err := matchHandler(&request, &conn)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	if handler == nil {
		t.Errorf("The Context must match the handler [%s][%s]", source, method)
		if err == nil {
			t.Error("nil handler must lead to an error")
		}
	} else {
		(*handler)(&request.Params, &request.Context)
		subscribed, isConnSubscribed := subscriptions[&conn]
		if isConnSubscribed {
			item := subscribed[0]
			if len(subscribed) == 2 {
				t.Logf("Array of subscribed items still contain 2 elements")
			}
			if item == otherSource {
				t.Logf("%s left in subscriptions", otherSource)
			} else {
				t.Errorf("Unexpected behavior. Subscribed item != '%s'", otherSource)
			}
		} else {
			t.Errorf("expecting to see '%s' in subscriptions", source)
		}
	}
	setupGlobals()
}

func Test_processRequest_request_to_subscribe(t *testing.T) {
	t.Log("Match the subscribe's handler")
	source := "source-defined"
	conn := websocket.Conn{}
	methods := DIRUD{}
	RegisterSource(source, &methods)
	request := JSONRPCrequest{}
	request.Method = "subscribe"
	request.Context = map[string]interface{}{"source": source}

	errs := processRequest(&request, &conn)
	if len(errs) == 1 {
		t.Logf("'%s' subscribed correctly", source)
		if errs[0] == nil {
			t.Logf("Expected nil error")
		} else {
			t.Errorf("Unexpected error %s", errs[0])
		}
	} else {
		t.Errorf("expecting to see an empty array in errs '%s'", errs)
	}

	subscribed, isConnSubscribed := subscriptions[&conn]
	if isConnSubscribed {
		item := subscribed[0]
		if item == source {
			t.Logf("%s subscribed correctly", source)
		} else {
			t.Errorf("Unexpected behavior. Subscribed item != '%s'", source)
		}
	} else {
		t.Errorf("expecting to see '%s' in subscriptions", source)
	}
	setupGlobals()
}
