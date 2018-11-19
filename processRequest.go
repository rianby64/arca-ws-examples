package main

import (
	"errors"

	"github.com/gorilla/websocket"
)

// Person whatever
type Person struct {
	ID    int
	Name  string
	Email string
}

// People whatever
type People []Person

// This is my Data-Base!
var people = People{
	Person{1, "Bob", "bob@mail.com"},
	Person{2, "Jeff", "jeff@mail.com"},
	Person{3, "Alice", "alice@mail.com"},
}

func response(request *JSONRPCrequest, conn *websocket.Conn,
	result interface{}, id *string) error {
	var myResponse JSONRPCresponse
	if id != nil {
		myResponse.ID = *id
	}
	myResponse.Jsonrpc = "2.0"
	myResponse.Result = result

	// Echo
	return conn.WriteJSON(myResponse)
}

func processRequest(request *JSONRPCrequest, conn *websocket.Conn) error {
	if request.Method == "getUsers" {
		id := request.ID
		return response(request, conn, people, &id)
	}
	if request.Method == "updateUser" {
		params := request.Params.(map[string]interface{})
		preid, ok := params["ID"]
		if !ok {
			return errors.New("params in request doesn't contain ID")
		}

		id := int(preid.(float64))
		for index, person := range people {
			if person.ID == id {
				if email, ok := params["Email"]; ok {
					people[index].Email = email.(string)
				}
				if name, ok := params["Name"]; ok {
					people[index].Name = name.(string)
				}
				return response(request, conn, people[index], nil)
			}
		}
	}
	return nil
}
