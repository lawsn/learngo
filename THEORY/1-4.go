package main

import (
	"fmt"
	"strings"

	"github.com/lawsn/learngo/something"
)

func lenAndUpper(name string) (length int, uppercase string) {
	// return len(name), strings.ToUpper(name)

	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	fmt.Println("Hello world")
	something.SayHello()
	something.SayBye()

	var name string = "vincoh"
	name = "Lynn"
	fmt.Println(name)
}
