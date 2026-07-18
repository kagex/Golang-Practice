package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var id string = "D1-12541" //если id = "error", вызовется ошибка с неизвестным пользователем
	result, err := userProfile(id)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(result)
	}
}

func userProfile(id string) (info string, err error) {

	balance, err := fetchUserInfo(id)

	if err != nil {
		return "", fmt.Errorf("fetch error: %w", err)
	}

	return fmt.Sprintf("Пользователь с id %s имеет на счету %0.2f руб.", id, float64(balance)/100), nil
}

func fetchUserInfo(id string) (balance int, err error) {

	if id == "error" {
		return 0, fmt.Errorf("Пользователь с идентификатором %s не найден.", id)
	}

	balance = generateBalance()

	if balance < 0 {
		return 0, fmt.Errorf("Извините, данный пользователь заблокирован. \nЕго баланс равен: %d копеек", balance)
	}
	return
}

func generateBalance() (balance int) { // Вспомогательная функция для генерации случайного значения баланса в копейках
	leftBorder := rand.IntN(2000) * -1
	rightBorder := rand.IntN(2001) + 1000
	modificator := rand.IntN(9) + 2
	balance = (rand.IntN(rightBorder-leftBorder) + leftBorder) * modificator
	return
}
