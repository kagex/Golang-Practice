package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	testMap := map[string]int{
		"Mitchel Resnick":   5,
		"Linus Torvalds":    5,
		"Donald Knuth":      3,
		"Tim Berners-Lee":   5,
		"Bjarne Stroustrup": 5,
	}
	fmt.Print(countVotes(testMap))
}

func countVotes(votes map[string]int) string {
	if len(votes) == 0 {
		return "Кандидаты потерялись."
	}

	var maxVotes int
	var winners []string

	for name, count := range votes {
		if count > maxVotes {
			maxVotes = count
			winners = []string{name}
		} else if count == maxVotes {
			winners = append(winners, name)
		}
	}

	if maxVotes <= 0 {
		return "Все голоса похищены!"
	}

	sort.Strings(winners)
	return strings.Join(winners, ", ")
}
