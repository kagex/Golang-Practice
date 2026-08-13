package main

import "fmt"

func main() {
	testSlice := []int{1, 3, 2, 3, 4, 3, 2, 1}
	fmt.Println(countMaxFrequency(testSlice))
}

func countMaxFrequency(slice []int) (maxFrequency int) {
	numbers := make(map[int]int)
	for _, value := range slice {
		numbers[value]++
	}
	for _, count := range numbers {
		if count > maxFrequency {
			maxFrequency = count
		}
	}
	return maxFrequency
}
