package main

import "fmt"

func f1_10() {
	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)

	for key, value := range nico {
		fmt.Println(key, value)
	}
}
