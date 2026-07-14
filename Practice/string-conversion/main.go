// Пакет main - точка входа в программу
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Выводит результат конвертации числа 62.231413 в строку
	fmt.Println(conversionToString(62.231413))
}

// conversionToString - функция конвертации числа с плавающей точкой в строку
// Принимает: price float64 - число для конвертации
// Возвращает: string - строковое представление числа с 3 знаками после запятой
// Использует strconv.FormatFloat с форматом 'f' (десятичный без экспоненты)
func conversionToString(price float64) string {
	stringPrice := strconv.FormatFloat(price, 'f', 3, 64)

	return stringPrice
}
