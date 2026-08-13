package main

import (
	"fmt"
)

func main() {
	testMatrix := [][]int{
		{1, 2, 3},
		{},
		{7, 8, 9, 22, 48},
		{10, 11},
	}

	fmt.Print(replaceEvenOnEvenIndices(testMatrix))
}

func replaceEvenOnEvenIndices(matrix [][]int) [][]int {
	copyMatrix := make([][]int, len(matrix))

	for i, row := range matrix {
		copyMatrix[i] = make([]int, len(row))
		for j, value := range row {
			if j%2 == 0 && value%2 == 0 {
				copyMatrix[i][j] = 0
			} else {
				copyMatrix[i][j] = value
			}
		}
	}
	return copyMatrix
}
