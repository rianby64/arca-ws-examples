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
var lastID = len(people)

func getUsers(request *JSONRPCrequest, conn *websocket.Conn) error {
	id := request.ID
	return response(request, conn, people, &id)
}

func updateUser(request *JSONRPCrequest, conn *websocket.Conn) error {
	params := request.Params.(map[string]interface{})
	preid, ok := params["ID"]
	if !ok {
		return errors.New("params in request doesn't contain ID")
	}
	preid2, ok := preid.(float64)
	if !ok {
		return errors.New("ID in params isn't int")
	}

	id := int(preid2)
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
	return nil
}

func insertUser(request *JSONRPCrequest, conn *websocket.Conn) error {
	params := request.Params.(map[string]interface{})
	lastID++
	newPerson := Person{ID: lastID}
	if email, ok := params["Email"]; ok {
		newPerson.Email = email.(string)
	}
	if name, ok := params["Name"]; ok {
		newPerson.Name = name.(string)
	}
	people = append(people, newPerson)
	return response(request, conn, newPerson, nil)
}
