package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	age     int
	favFood []string
}

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
	favFood := []string{"kimchi", "ramen"}
	nico := person{"nico", 18, favFood}
	fmt.Println(nico.age)
}
