package main

import (
	"fmt"
)

func main() {
	testMatrix := [][]int{
		{1, 2, 3},
		{4, 5},
		{6},
	}

	fmt.Print(mirrorMatrix(testMatrix))
}

func mirrorMatrix(matrix [][]int) [][]int {
	resultMatrix := make([][]int, len(matrix))

	for i, line := range matrix {
		mirroredLine := make([]int, len(line))

		for j, value := range line {
			mirroredLine[len(line)-1-j] = value
		}

		resultMatrix[i] = mirroredLine
	}

	return resultMatrix
}
