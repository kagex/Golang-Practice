package main

import (
	"fmt"
	"log"
	"strings"
)

const (
	keyboard, keyboardPrice     = "Клавиатура JZ9", 19200
	headphones, headphonesPrice = "Наушники N45", 9600
	smartphone, smartphonePrice = "Смартфон S10", 55000
)

func main() {

	var userInput string

	fmt.Println("Добрый день.\nКакой товар вас интересует?")

	if _, err := fmt.Scan(&userInput); err != nil {
		log.Fatal(err)
	}

	switch {
	case strings.Contains(strings.ToLower(keyboard), strings.ToLower(userInput)):
		fmt.Printf("%s: %d", keyboard, keyboardPrice)

	case strings.Contains(strings.ToLower(headphones), strings.ToLower(userInput)):
		fmt.Printf("%s: %d", headphones, headphonesPrice)

	case strings.Contains(strings.ToLower(smartphone), strings.ToLower(userInput)):
		fmt.Printf("%s: %d", smartphone, smartphonePrice)

	default:
		fmt.Printf("Товар %s не найден.", userInput)
	}
}
