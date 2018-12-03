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
		response chan interface{}) error {
		response <- &people
		return nil
	},
	Update: func(requestParams *interface{}, context *interface{},
		response chan interface{}) error {
		params := (*requestParams).(map[string]interface{})
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
				response <- &people[index]
				return nil
			}
		}
		return errors.New("nothing")
	},
	Insert: func(requestParams *interface{}, context *interface{},
		response chan interface{}) error {
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
		response <- &newPerson
		return nil
	},
	Delete: func(requestParams *interface{}, context *interface{},
		response chan interface{}) error {
		params := (*requestParams).(map[string]interface{})
		preid, ok := params["ID"]
		if !ok {
			return errors.New("params in request doesn't contain ID")
		}
		preid2, ok := preid.(float64)
		if !ok {
			return errors.New("ID in params isn't int")
		}

		id := int(preid2)
		deletedPerson := Person{ID: id}
		for i, person := range people {
			if person.ID == id {
				people = append(people[:i], people[i+1:]...)
				response <- &deletedPerson
				return nil
			}
		}
		return errors.New("nothing")
	},
}
