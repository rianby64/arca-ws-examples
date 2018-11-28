package examples

import (
	"errors"

	"../arca"
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
var lastUsersID = len(people)

var usersCRUD = arca.DIRUD{
	Read: func(requestParams *interface{}, context *interface{},
		response arca.ResponseHandler) []error {
		return response(&people)
	},
	Update: func(requestParams *interface{}, context *interface{},
		response arca.ResponseHandler) []error {
		params := (*requestParams).(map[string]interface{})
		preid, ok := params["ID"]
		if !ok {
			return []error{errors.New("params in request doesn't contain ID")}
		}
		preid2, ok := preid.(float64)
		if !ok {
			return []error{errors.New("ID in params isn't int")}
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
				return response(&people[index])
			}
		}
		return []error{errors.New("nothing")}
	},
	Insert: func(requestParams *interface{}, context *interface{},
		response arca.ResponseHandler) []error {
		params := (*requestParams).(map[string]interface{})
		lastUsersID++
		newPerson := Person{ID: lastUsersID}
		if email, ok := params["Email"]; ok {
			newPerson.Email = email.(string)
		}
		if name, ok := params["Name"]; ok {
			newPerson.Name = name.(string)
		}
		people = append(people, newPerson)
		return response(&newPerson)
	},
	Delete: func(requestParams *interface{}, context *interface{},
		response arca.ResponseHandler) []error {
		params := (*requestParams).(map[string]interface{})
		preid, ok := params["ID"]
		if !ok {
			return []error{errors.New("params in request doesn't contain ID")}
		}
		preid2, ok := preid.(float64)
		if !ok {
			return []error{errors.New("ID in params isn't int")}
		}

		id := int(preid2)
		deletedPerson := Person{ID: id}
		for i, person := range people {
			if person.ID == id {
				people = append(people[:i], people[i+1:]...)
				return response(&deletedPerson)
			}
		}
		return []error{errors.New("nothing")}
	},
}
