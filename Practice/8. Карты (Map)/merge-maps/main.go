package main

import "fmt"

func main() {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"b": 3, "c": 4, "d": 5}

	fmt.Println(mergeMaps(m1, m2))
}

func mergeMaps(m1, m2 map[string]int) map[string]int {
	result := make(map[string]int)

	for index, value := range m1 {
		result[index] += value
	}

	for index, value := range m2 {
		result[index] += value
	}

	return result
}
