package main

import "fmt"

func main() {
	names := []string{"vincoh", "ms", "dal"}
	names = append(names, "mine")
	fmt.Println(names)
}
