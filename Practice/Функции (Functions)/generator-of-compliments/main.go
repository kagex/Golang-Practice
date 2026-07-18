package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	name := "Александр"
	fmt.Print(generateCompliment(name))
}

func generateCompliment(name string) (compliment string) {
	switch random := rand.IntN(3); {
	case random == 0:
		compliment = fmt.Sprintf("Ты великолепен, %s!", name)
	case random == 1:
		compliment = fmt.Sprintf("У тебя потрясающая улыбка, %s!", name)
	default:
		compliment = fmt.Sprintf("Ты вдохновляешь, %s!", name)
	}
	return compliment
}
