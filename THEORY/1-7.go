package main

import (
	"fmt"
)

func canIDrinks(age int) bool {

	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func f1_7() {
	fmt.Println(canIDrinks(16))
}
