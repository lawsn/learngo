package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	total := 0

	for index, number := range numbers {
		fmt.Println(index, number)
	}

	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

func f1_5() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
