package examples

import (
	"log"

	"../arca"
)

// Start begin
func Start() {
	log.Println("Start")
	arca.RegisterSource("Users", &usersCRUD)
	arca.RegisterSource("Goods", &goodsCRUD)
}
