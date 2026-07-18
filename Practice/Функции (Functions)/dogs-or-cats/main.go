package main

import "fmt"

func main() {
	fmt.Print(PetBattle(3, 2))
}

func PetBattle(countCats, countDogs int) (result string) {
	switch {
	case countCats > countDogs:
		result = fmt.Sprintf("Котики победили со счетом %d:%d!", countCats, countDogs)
	case countCats > countDogs:
		result = fmt.Sprintf("Собачки победили со счетом %d:%d!", countDogs, countCats)
	default:
		result = "Ничья! Все дружат!"
	}

	return result
}
