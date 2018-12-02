package arca

import (
	"testing"
)

func Test_Handle(t *testing.T) {
	t.Log("Match a handler fails if no method defined in request")

}

/*
// THIS IS AN EXAMPLE THAT I WANT TO USE TO IMPROVE THE REQUEST PROCESSING
import (
	"log"
)

func sending(c chan<- interface{}) {
	for i := 0; i < 10; i++ {
		log.Println("sending", i)
		c <- i
	}
	log.Println("done sending")
}

func main() {
	chann := make(chan interface{})
	go (func() {
		sending(chann)
		close(chann)
	})()

	for {
		recevied, ok := <-chann
		log.Println(ok, recevied)
		if !ok {
			log.Println("breaking")
			break
		}
		log.Println(recevied, "expecting")
	}
}
*/
