package main

import "fmt"

func main() {
	fmt.Print(filterEven(1, 2, 3, 4, 5, 6))
}

func filterEven(values ...int) []int {
	result := []int{}

	for _, v := range values {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}
