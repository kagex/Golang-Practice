package main

import (
	"fmt"
	"maps"
	"math"
)

func main() {
	m1 := map[string][]int{
		"a": {0, 2, 3},
		"b": {4, 5, 6},
	}
	m2 := map[string][]int{
		"a": {3, 1, 2},
		"b": {6, 5, 4},
	}
	fmt.Println(compareMaxValues(m1, m2))
}

func compareMaxValues(map1, map2 map[string][]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	result := maps.EqualFunc(map1, map2, func(value1, value2 []int) bool {
		return getMaxValue(value1) == getMaxValue(value2)
	})
	return result
}

func getMaxValue(slice []int) int {
	if len(slice) == 0 {
		return math.MinInt
	}
	maxValue := slice[0]
	for _, value := range slice {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}
