package main

import (
	"fmt"
	"slices"
)

func main() {
	testSlice := []int{15, -20}
	fmt.Print(PlayWithSlice(testSlice))
}

func PlayWithSlice(slice []int) []int {
	clone := slices.Clone(slice)

	position := -1

	for i := len(clone) - 1; i >= 0; i-- {
		if clone[i] > 10 {
			position = i + 1
			break
		}
	}

	if position != -1 {
		clone = append(clone[:position], append([]int{100}, clone[position:]...)...)
	}

	elementSum := 0
	for _, value := range clone {
		elementSum += value
	}

	if elementSum > 100 {
		clone = append(clone, 500)
	}

	evenCount, oddCount := 0, 0
	for _, value := range slice {

		if value%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	if evenCount > oddCount {
		clone = append([]int{1000}, clone...)
	}
	return clone
}
