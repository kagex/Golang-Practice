package main

import (
	"fmt"
)

func main() {
	fmt.Print(sumOfDigits(456))
}

func sumOfDigits(number int) int {
	if number < 0 {
		number = -number
	}
	if number < 10 {
		return number
	}
	return number%10 + sumOfDigits(number/10)
}
