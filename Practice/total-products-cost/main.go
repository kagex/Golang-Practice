package main

import (
	"fmt"
	//"math"
	"strconv"
)

func main() {
	// Объявляем переменные для цены и количества товара
	priceStr, quantityStr := "19.9", "19"

	// Вызываем функцию для вычисления полной стоимости и выводим результат
	fmt.Printf("Полная стоимость товаров:%.2f ", totalPrice(priceStr, quantityStr))
}

func totalPrice(priceStr, quantityStr string) (totalPrice float64) {
	// Преобразуем строку цены в float64
	price, _ := strconv.ParseFloat(priceStr, 64)
	// Преобразуем строку количества в float64
	quantity, _ := strconv.ParseFloat(quantityStr, 64)

	// Возвращаем результат
	return price * quantity

}
