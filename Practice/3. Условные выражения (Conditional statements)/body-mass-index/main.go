package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var weight, height float64

	// Считываем вес пользователя
	fmt.Print("Введите ваш вес (кг): ")

	// Проверка корректности ввода веса
	if _, err := fmt.Scan(&weight); err != nil || weight <= 0 {
		log.Fatal("Неверно задан вес")
	}

	// Считываем рост пользователя
	fmt.Print("Введите ваш рост (см): ")

	// Проверка корректности ввода роста
	if _, err := fmt.Scan(&height); err != nil || height <= 0 {
		log.Fatal("Неверно задан рост")
	}

	// Вычисляем ИМТ и получаем категорию
	bodyMassIndex, bodyMassCategory := calculateBMI(weight, height)

	// Выводим результат
	fmt.Printf("\nВаш ИМТ: %.2f\nКатегория: %s", bodyMassIndex, bodyMassCategory)
}

// calculateBMI рассчитывает индекс массы тела и определяет категорию веса
// weight: вес в килограммах
// height: рост в сантиметрах
// Возвращает: значение ИМТ и категорию веса
func calculateBMI(weight, height float64) (bodyMassIndex float64, bodyMassCategory string) {
	// Формула: ИМТ = вес / (рост в метрах)²
	bodyMassIndex = weight / (math.Pow(height/100, 2))

	// Определяем категорию веса по ИМТ
	switch {
	case bodyMassIndex < 18.5:
		bodyMassCategory = "Недостаточный вес"
	case bodyMassIndex >= 18.5 && bodyMassIndex < 25:
		bodyMassCategory = "Нормальный вес"
	case bodyMassIndex >= 25 && bodyMassIndex < 30:
		bodyMassCategory = "Избыточный вес"
	default:
		bodyMassCategory = "Ожирение"
	}
	return bodyMassIndex, bodyMassCategory
}
