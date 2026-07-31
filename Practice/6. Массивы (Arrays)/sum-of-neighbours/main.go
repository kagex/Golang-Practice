package main

import "fmt"

const arraySize = 10

func main() {
	testArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := SumNeighbours(testArray)
	fmt.Println(result)
}

func SumNeighbours(arr [arraySize]int) (result [arraySize]int) {
	arrLength := len(arr) - 1
	for i := 0; i <= arrLength; i++ {
		switch {
		case i == 0:
			result[i] = arr[i+1]
		case i == arrLength:
			result[i] = arr[i-1]
		default:
			result[i] = arr[i-1] + arr[i+1]
		}
	}
	return result
}
