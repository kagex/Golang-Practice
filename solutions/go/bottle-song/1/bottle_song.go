package bottlesong

import (
	"fmt"
	"strings"
)

var numbers = []string{
	"No", "One", "Two", "Three", "Four", "Five",
	"Six", "Seven", "Eight", "Nine", "Ten",
}

func Recite(startBottles, takeDown int) []string {
	var result []string

	for i := 0; i < takeDown; i++ {
		if i > 0 {
			result = append(result, "")
		}
		result = append(result, verse(startBottles-i)...)
	}
	return result
}

func verse(n int) []string {
	currentBottle := "bottles"
	if n == 1 {
		currentBottle = "bottle"
	}
	remaining := n - 1
	remainingBottle := "bottles"
	if remaining == 1 {
		remainingBottle = "bottle"
	}
	firstLine := fmt.Sprintf("%s green %s hanging on the wall,",
		numbers[n], currentBottle)
	thirdLine := "And if one green bottle should accidentally fall,"
	fourthLine := fmt.Sprintf("There'll be %s green %s hanging on the wall.",
		strings.ToLower(numbers[remaining]), remainingBottle)
	return []string{firstLine, firstLine, thirdLine, fourthLine}
}