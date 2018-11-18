package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// MyTestParams whatever
type MyTestParams struct {
	Message string
	A       []string
}

// MyRequest whatever
type MyRequest struct {
	Jsonrpc string
	Method  string
	Params  interface{}
	ID      string
}

func Test_cli(t *testing.T) {
	println("let's begin")
	go main()

	time.Sleep(time.Second * 2)
	println("here we go")

	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		println("Something went wrong", err)
		return
	}
	println("Hello")

	req := MyRequest{
		"2.0",
		"mymethod",
		MyTestParams{
			"thats my message",
			[]string{"a", "xx", "b", "yy", "c", "zz"},
		},
		"myId",
	}

	r, err := json.Marshal(req)
	if err != nil {
		println("Error", err)
	}
	println(string(r))
	c.WriteMessage(websocket.TextMessage, r)

	_, data, err := c.ReadMessage()
	if err != nil {
		println("Error at reading", err)
	}
	println(string(data))

	println("now its over")
}
