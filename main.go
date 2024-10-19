package main

import (
	"fmt"
)

func FiveSteps(array []int, x int) []int {
	for i := 0; i < len(array); i++ {
		if array[i] != x {
			array = append(array, array[i])
		}
	}
	return array[len(array):]
}

func main() {
	var a []int = []int{1, 2, 3, 4, 5}
	fmt.Println(FiveSteps(a, 2))
}
