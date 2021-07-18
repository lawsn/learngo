package main

import (
	"fmt"
)

func canIDrink(age int) bool {

	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func f1_6() {
	fmt.Println(canIDrink(16))
}
