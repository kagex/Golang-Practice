package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Print(printMagic(numbers))
}

func printMagic(numbers []int) string {
	var sb strings.Builder

	sb.WriteString("[")
	for i := 0; i < len(numbers); i++ {
		multiplication := 1
		for j, value := range numbers {
			if j == i {
				continue
			}
			multiplication *= value
		}
		sb.WriteString(strconv.Itoa(multiplication))
		if i != len(numbers)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
