package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	rollDice(12)
}

func rollDice(necessarySum int) {
	var dice1, dice2, diceSum, count int

	for diceSum != necessarySum {
		count++

		dice1 = rand.IntN(6) + 1
		dice2 = rand.IntN(6) + 1
		diceSum = dice1 + dice2

		if count >= 1000 {
			fmt.Println("Подозрение на бесконечный цикл, выходим")
			break
		}

		if diceSum != necessarySum {
			fmt.Printf("Выпало %d и %d, в сумме %d, бросаем еще раз.\n", dice1, dice2, diceSum)
		} else {
			switch {
			case count%10 == 1 && count%100/10 != 1:
				fmt.Printf("Выпало %d и %d, в сумме %d, на это потребовался %d бросок.\n", dice1, dice2, diceSum, count)
			case count%10 <= 4 && count%100/10 != 1:
				fmt.Printf("Выпало %d и %d, в сумме %d, на это потребовалось %d броска.\n", dice1, dice2, diceSum, count)
			default:
				fmt.Printf("Выпало %d и %d, в сумме %d, на это потребовалось %d бросков.\n", dice1, dice2, diceSum, count)
			}

		}
	}
}
