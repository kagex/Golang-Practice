package main

import (
	"fmt"
	"maps"
)

func main() {
	inputMap := map[string][]int{
		"a": {1, 2, 4},
		"b": {1, 1, 1, 1},
		"c": {20, 30, -100},
	}
	fmt.Println("Первоначальный map:", inputMap)

	RemoveSlicesBySum(inputMap)

	fmt.Println("\nMap после удаления слайсов с суммой элементов больше 6:", inputMap)
}

func RemoveSlicesBySum(input map[string][]int) {
	maps.DeleteFunc(input, func(key string, value []int) bool {
		var elementSum int
		for _, v := range value {
			elementSum += v
		}

		return elementSum > 6
	})
}
