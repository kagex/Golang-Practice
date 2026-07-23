package main

import "fmt"

func main() {
	// Объявление переменной для хранения текущего времени
	var currentTime int

	// Вывод приветственного сообщения и запрос ввода у пользователя
	fmt.Println("- Привет, я тут чтобы помочь вам узнать какое сейчас время суток! Сколько сейчас часов?")
	fmt.Print("\n- Сейчас на часах: ")
	_, err := fmt.Scanln(&currentTime)

	// Проверка на корректность введенных данных (24 часа считаем валидным)
	if err != nil || currentTime < 0 || currentTime > 24 {
		fmt.Println("Неверно задано время")
		return
	}

	// Определение части суток на основе введенного времени
	var partOfTime string
	switch {
	case currentTime >= 6 && currentTime < 12:
		partOfTime = "утро"
	case currentTime >= 12 && currentTime < 18:
		partOfTime = "день"
	case currentTime >= 18 && currentTime < 23:
		partOfTime = "вечер"
	default:
		partOfTime = "ночь"
	}

	// Вывод результата пользователю
	fmt.Printf("Сейчас %dч. - %s.", currentTime, partOfTime)
}
