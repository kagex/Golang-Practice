package main

import (
	"errors"
	"fmt"

	//"math"
	"math/rand/v2"
	"os"
)

func main() {
	for range 100 {
		random = rand.IntN(100) + 1
		guesses = 0
		result := play()
		if result != random {
			fmt.Printf("Неверный ответ. Было загадано число %d, а в ответе получили число %d", random, result)
			os.Exit(-1)
		}
	}
	fmt.Println("Тест пройден! Все 100 чисел угаданы за 6 или менее попыток.")
}

var guesses int
var random int

func guess(num int) (int, error) {
	if guesses >= 6 {
		return 0, errors.New("too many attempts")
	}
	guesses++
	if num > random {
		return -1, nil
	}
	if num < random {
		return 1, nil
	}
	return 0, nil
}

func play() int {
	var middleOfRange int // Переменная для хранения значения середины
	leftBorder := 1
	rightBorder := 100
	for leftBorder <= rightBorder {

		if leftBorder == rightBorder {
			return leftBorder
		}

		middleOfRange = (leftBorder + rightBorder) / 2
		res, err := guess(middleOfRange) // Передаем наше значение в функцию guess и сохраняем в переменную ответ от функции

		if res == 0 && err == nil { // Число угадали, выходим из цикла
			break
		}

		if res == -1 && err == nil { // Число меньше нашего значения, поэтому смещаем нашу правую границу
			rightBorder = middleOfRange - 1
		}

		if res == 1 && err == nil { // Число больше нашего значения, поэтому смещаем нашу левую границу
			leftBorder = middleOfRange + 1
		}
	}
	return middleOfRange
}
