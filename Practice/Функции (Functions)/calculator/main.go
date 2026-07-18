package main

import (
	"errors"
	"fmt"
)

func main() {
	var firstArgument, secondArgument float64 = 10, 5
	var operation string = "*"

	result, err := calculate(firstArgument, secondArgument, operation)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%.2f %s %.2f = %.2f", firstArgument, operation, secondArgument, result)
	}
}

func calculate(firstArgument, secondArgument float64, operation string) (result float64, err error) {

	switch operation {
	case "+", "add":
		return firstArgument + secondArgument, nil
	case "-", "subtract":
		return firstArgument - secondArgument, nil
	case "*", "multiply":
		return firstArgument * secondArgument, nil
	case "/", "divide":
		if secondArgument == 0 {
			return 0, errors.New("division by zero")
		}
		return firstArgument / secondArgument, nil
	default:
		return 0, errors.New("unknown operation")
	}
}
