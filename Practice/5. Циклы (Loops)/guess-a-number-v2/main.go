package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Привет, я - Нумератор.\nЯ загадаю целое число, а вы попробуйте его угадать!")
	fmt.Println("Чуть не забыл правила:\n" +
		"1. Я загадываю числа от 1 до 100\n" +
		"2. Я буду играть с вами пока вам не надоест\n" +
		"3. Если вы сдаетесь, напишите сдаюсь или surrender\n" +
		"4. Если захотите выйти из игры, то просто напишите выход или exit")
	for {
		parseUserInput()
	}
}

var guesses, random, gamesCount int
var gameInProgress bool = false

func parseUserInput() {
	userInput := ""
	fmt.Scanln(&userInput)
	userInput = strings.ToLower(userInput)

	if userInput == "выход" || userInput == "exit" {
		fmt.Println("До свидания!")
		if gamesCount > 0 {
			switch {
			case gamesCount%10 == 1 && gamesCount%100/10 != 1:
				fmt.Printf("Мы с вами провели %d замечательную игру.\nПриходите еще!", gamesCount)
			case gamesCount%10 >= 2 && gamesCount%10 <= 4 && gamesCount%100/10 != 1:
				fmt.Printf("Мы с вами поиграли целых %d раза. Было замечательно!\n", gamesCount)
			default:
				fmt.Printf("Мы с вами замечательно провели время. Сыграли %d раз!\n", gamesCount)
			}
		}
		os.Exit(0)
	}

	if userInput == "сдаюсь" || userInput == "surrender" {
		if gameInProgress == false {
			fmt.Println("Подождите сдаваться, мы даже еще не начали. Вводите ваше число")
			return
		}
		fmt.Printf("Давайте подведем итоги. Я загадал %d. Количество ваших попыток: %d.\nДавайте попробуем еще раз, я загадал новое число, полегче", random, guesses)
		gameInProgress = false
		return
	}
	num, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Что-то мне подсказывает что вы ввели не целое число. Попробуйте еще раз, я не буду вам засчитывать попытку")
		return
	}

	if num < 1 || num > 100 {
		fmt.Println("Число должно быть от 1 до 100! Попытка не засчитана.")
		return
	}

	gameLogic(num)
}

func gameLogic(num int) {
	if gameInProgress == false {
		guesses = 0
		gamesCount++
		gameInProgress = true
		random = rand.IntN(100) + 1
	}

	guesses++

	switch {
	case num > random:
		fmt.Println("Ваше число больше числа что я загадал")
	case num < random:
		fmt.Println("Ваше число меньше числа что я загадал")
	case num == random:
		fmt.Printf("Поздравляю, вы угадали число. Число ваших попыток: %d!\nИгра окончена!\nДавайте продолжим, я загадал новое число!\n", guesses)
		gameInProgress = false
	}
}
