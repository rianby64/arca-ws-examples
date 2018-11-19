package main

import "github.com/gorilla/websocket"

// MyParams whatever
type MyParams struct {
	Message string
	A       []string
}

type errorMyParams struct {
	why string
}

func (e errorMyParams) Error() string {
	return e.why
}

// Person whatever
type Person struct {
	ID    int
	Name  string
	Email string
}

// People whatever
type People []Person

// Object whatever
type Object map[string]interface{}

func processRequest(request *JSONRPCrequest, conn *websocket.Conn) error {

	var params MyParams

	// I'd like to use .(Object) instead of that long definition. How to proceed?
	// message, ok := request.Params.(Object)["Message"]

	// extract Message
	message, ok := request.Params.(map[string]interface{})["Message"]
	if !ok {
		return errorMyParams{"Error while converting Message"}
	}
	params.Message = message.(string)

	// Extract A
	preA, ok := request.Params.(map[string]interface{})["A"].([]interface{})
	if !ok {
		return errorMyParams{"Error while converting A"}
	}
	params.A = make([]string, len(preA))
	for key, value := range preA {
		conv, ok := value.(string)
		if !ok {
			return errorMyParams{"Error while converting A member"}
		}

		params.A[key] = conv
	}

	var myResponse JSONRPCresponse
	myResponse.ID = "responseID"
	myResponse.Jsonrpc = "2.0"
	myResponse.Result = People{
		Person{1, "Bob", "bob@mail.com"},
		Person{2, "Jeff", "jeff@mail.com"},
		Person{3, "Alice", "alice@mail.com"},
	}

	// Echo
	return conn.WriteJSON(myResponse)
}
