package main

import (
	"fmt"
	"strings"

	"github.com/lawsn/learngo/something"
)

func repeateMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpper(name string) (int, string) {

	return len(name), strings.ToUpper(name)
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("Hello world")
	something.SayHello()
	something.SayBye()
	name := "vincoh"
	// name = false

	fmt.Println(name)

	fmt.Println(multiply(2, 2))

	totlaLengh, upperName := lenAndUpper(name)
	fmt.Println(totlaLengh, upperName)

	fmt.Println(lenAndUpper(name))

	repeateMe("vincoh", "vincent", "beomseok", "aaron", "tatiana")
}
