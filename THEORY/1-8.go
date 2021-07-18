package main

import (
	"fmt"
)

func f1_8() {
	a := 2
	b := &a
	a = 10
	fmt.Println(a, b, *b)
	*b = 20
	fmt.Println(a)
}
