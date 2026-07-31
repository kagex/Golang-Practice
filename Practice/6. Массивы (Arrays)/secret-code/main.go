package main

import (
	"fmt"
	"strconv"
)

func main() {
	testArray := [5]int{3, 8, 1, 8, 1}
	fmt.Print(generateCode(testArray))
}

func generateCode(input [5]int) (code string) {
	minValue := input[0]
	maxValue := input[0]

	for _, value := range input {
		if value < minValue {
			minValue = value
		}
		if value > maxValue {
			maxValue = value
		}
	}

	code += strconv.Itoa(minValue)

	for _, value := range input {
		switch {
		case value%2 == 0:
			code += "E" + strconv.Itoa(value)
		case value%2 != 0:
			code += "O" + strconv.Itoa(value)
		}
	}

	code += strconv.Itoa(maxValue)
	return code
}
