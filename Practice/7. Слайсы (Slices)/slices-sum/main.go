package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	fmt.Println(SumSlices(slice1, slice2))
}

func SumSlices(slice1, slice2 []int) []int {
	minLength := min(len(slice1), len(slice2))
	result := make([]int, minLength)

	for i := range minLength {
		result[i] = slice1[i] + slice2[i]
	}
	return result
}
