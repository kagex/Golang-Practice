package main

import (
	"fmt"
	"strings"
)

func main() {
	printTable(5)
}

func printTable(num int) {
	var sb strings.Builder

	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			fmt.Fprintf(&sb, "%d x %d = %d", i, j, i*j)
			if j != num {
				sb.WriteString("\t")
			} else {
				sb.WriteString("\n")
			}
		}
	}
	fmt.Print(sb.String())
}
