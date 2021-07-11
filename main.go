package main

import (
	"fmt"

	"github.com/lawsn/learngo/something"
)

func main() {
	fmt.Println("Hello world")
	something.SayHello()
	something.SayBye()

	var name string = "vincoh"
	name = "Lynn"
	fmt.Println(name)
}
