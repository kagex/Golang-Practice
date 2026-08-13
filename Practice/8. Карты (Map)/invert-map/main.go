package main

import (
	"fmt"
	"slices"
)

type products map[int][]string

func main() {
	m := map[string]int{
		"banana":     2,
		"apple":      1,
		"grapefruit": 3,
		"cherry":     1,
	}
	invertedMap := invertMap(m)
	printMap(invertedMap)
}

func invertMap(m map[string]int) products {
	p := make(products)
	for key, value := range m {
		p[value] = append(p[value], key)
	}

	return p
}

func printMap(p products) {
	keys := make([]int, 0, len(p))

	for key := range p {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	fmt.Println("{")

	for _, key := range keys {
		values := p[key]
		sortedValues := make([]string, len(values))
		copy(sortedValues, values)

		// Сортируем строки в алфавитном порядке
		slices.Sort(sortedValues)

		fmt.Printf("  %d: [", key)

		for i, value := range sortedValues {
			if i > 0 {
				fmt.Print(", ")
			}

			fmt.Printf("%q", value)
		}

		fmt.Println("],")
	}

	fmt.Println("}")
}
