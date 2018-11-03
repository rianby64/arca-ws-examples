package main

import (
	"testing"
	"time"
)

func runServer() {

	println("begin")
	go main()
	println("wait")
	time.Sleep(time.Second * 5)
	println("done")
}

func Test_cli(t *testing.T) {

	defer (func() {
		println("defered")
	})()

	go runServer()
	println("here we go")

	time.Sleep(time.Second * 6)
	println("now its over")
}
