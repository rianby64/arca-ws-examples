package arca

import (
	"testing"

	"github.com/gorilla/websocket"
)

func Test_setupGlobals_initial(t *testing.T) {
	t.Log("Check initial state")
	if len(conns) == 0 {
		t.Log("conns has 0 item")
	} else {
		t.Error("expected conns to be nil")
	}
	if len(subscriptions) == 0 {
		t.Log("subscriptions has 0 item")
	} else {
		t.Error("expected subscriptions to be nil")
	}
	if len(handlers) == 0 {
		t.Log("handlers has 0 item")
	} else {
		t.Error("expected handlers to be nil")
	}
}

func Test_setupGlobals_executed(t *testing.T) {
	t.Log("Check result of execution setupGlobals")
	conn := &websocket.Conn{}
	conns = append(conns, conn)
	subscriptions[conn] = nil
	handlers[""] = nil

	if len(conns) == 1 {
		t.Log("conns has 1 item")
	} else {
		t.Error("expected conns to have one element")
	}
	if len(subscriptions) == 1 {
		t.Log("subscriptions has 1 item")
	} else {
		t.Error("expected subscriptions to have one element")
	}
	if len(handlers) == 1 {
		t.Log("handlers has 1 item")
	} else {
		t.Error("expected handlers to have one element")
	}
	setupGlobals()
}

func Test_appendConnection_once(t *testing.T) {
	t.Log("Check appendConnection if call it once")
	conn := &websocket.Conn{}
	conn2 := &websocket.Conn{}
	appendConnection(conn)

	if len(conns) == 1 {
		t.Log("append one element")
		appendConnection(conn2)
		if len(conns) == 2 {
			t.Log("append zero elements")
		} else {
			t.Error("append repeating element")
		}
	} else {
		t.Error("append is not happening")
	}
	setupGlobals()
}

func Test_appendConnection_twice(t *testing.T) {
	t.Log("Check appendConnection if call it once")
	conn := &websocket.Conn{}
	conn2 := &websocket.Conn{}
	appendConnection(conn)
	appendConnection(conn)

	if len(conns) == 1 {
		t.Log("append one element")
		appendConnection(conn2)
		appendConnection(conn2)
		if len(conns) == 2 {
			t.Log("append zero elements")
		} else {
			t.Error("append repeating element")
		}
	} else {
		t.Error("append is not happening")
	}
	setupGlobals()
}

func Test_removeConnection(t *testing.T) {
	t.Log("Check removeConnection")
	conn := &websocket.Conn{}
	appendConnection(conn)
	removeConnection(conn, false)

	if len(conns) == 0 {
		t.Log("append one element")
	} else {
		t.Error("remove is not happening")
	}
	setupGlobals()
}

func Test_subscribe_once(t *testing.T) {
	t.Log("Check subscribe if call it once")
	conn := &websocket.Conn{}
	sourceMock := "source_mock"
	appendConnection(conn)
	subscribe(conn, sourceMock)

	if len(subscriptions) == 1 {
		t.Log("subscriptions got one element")
		if subscriptions[conn][0] == sourceMock {
			t.Log("subscriptions got one element")
		} else {
			t.Error(sourceMock, "mock is not present")
		}
	} else {
		t.Error("subscriptions array is dirty")
	}
	setupGlobals()
}

func Test_subscribe_twice(t *testing.T) {
	t.Log("Check subscribe if call it once")
	conn := &websocket.Conn{}
	sourceMock := "source_mock"
	appendConnection(conn)
	subscribe(conn, sourceMock)
	subscribe(conn, sourceMock)

	if len(subscriptions) == 1 {
		t.Log("subscriptions got one element")
		if subscriptions[conn][0] == sourceMock {
			t.Log("subscriptions got one element")
		} else {
			t.Error(sourceMock, "mock is not present")
		}
	} else {
		t.Error("subscriptions array is dirty")
	}
	setupGlobals()
}

func Test_RegisterSource(t *testing.T) {
	sourceMock := "source_mock"
	RegisterSource(sourceMock, DIRUD{})

	if len(handlers) == 1 {
		t.Log("handlers got one element")
	} else {
		t.Error("handlers array is dirty")
	}
}
