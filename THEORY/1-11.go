package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func f1_11() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico.age)
}
