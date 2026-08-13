package main

import (
	"fmt"
	"strings"
)

// User представляет пользователя
type User struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Address Address
	Cart    []CartItem
}

// Address представляет адрес пользователя
type Address struct {
	Street     string
	City       string
	PostalCode string
}

// CartItem представляет элемент в корзине
type CartItem struct {
	Product  Product
	Quantity int
}

// Product представляет продукт в корзине
type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
	Category    string
	Brand       string
	Rating      float64
	Reviews     int
}

func main() {
	testUser := User{
		ID:    1,
		Name:  "Иван Петров",
		Email: "ivan.petrov@example.com",
		Phone: "+7 999 123-45-67",
		Address: Address{
			Street:     "Улица Ленина",
			City:       "Москва",
			PostalCode: "101000",
		},
		Cart: []CartItem{
			{
				Product: Product{
					ID:          1,
					Name:        "Ноутбук",
					Description: "Мощный ноутбук для работы и игр",
					Price:       59990,
					Category:    "Электроника",
					Brand:       "Brand A",
					Rating:      4.5,
					Reviews:     120,
				},
				Quantity: 1,
			},
			{
				Product: Product{
					ID:          2,
					Name:        "Смартфон",
					Description: "Современный смартфон с отличной камерой",
					Price:       29990,
					Category:    "Электроника",
					Brand:       "Brand B",
					Rating:      4.7,
					Reviews:     200,
				},
				Quantity: 2,
			},
			{
				Product: Product{
					ID:          3,
					Name:        "Наушники",
					Description: "Беспроводные наушники с шумоподавлением",
					Price:       7990,
					Category:    "Аудио",
					Brand:       "Brand C",
					Rating:      4.3,
					Reviews:     80,
				},
				Quantity: 1,
			},
		},
	}

	printInfo(testUser)
}

func printInfo(userInfo User) {
	// Информация о пользователе
	fmt.Printf("Покупатель %s. Телефон: %s. Адрес: г. %s, %s.\n", userInfo.Name, userInfo.Phone, userInfo.Address.City, userInfo.Address.Street)

	productExist := false          // флаг для определения, был ли найден товар соответствующей категории в корзине
	highCostProducts := []string{} // слайс для хранения товаров дороже 10 000
	totalCost := 0                 // переменная для хранения стоимости всей корзины пользователя

	// Перебираем корзину пользователя и проходимся по каждому продукту
	for _, product := range userInfo.Cart {
		// проверяем если в корзине "Электроника" и электроника не была обнаружена, и если обнаружили, то меняем флаг на true
		if product.Product.Category == "Электроника" && !productExist {
			productExist = true
		}

		// проверяем если в корзине товары дороже 10 000 и добавляем их в слайс
		if product.Product.Price >= 10000 {
			highCostProducts = append(highCostProducts, product.Product.Name)
		}

		totalCost += product.Quantity * product.Product.Price
	}

	if productExist {
		fmt.Println("Пользователь является покупателем электроники.")
	} else {
		fmt.Println("Пользователь не является покупателем электроники.")
	}

	if len(highCostProducts) > 0 {
		fmt.Printf("Товары в корзине, где цена 10000 и более: %s.\n", strings.Join(highCostProducts, ", "))
	} else {
		fmt.Println("Товары в корзине, где цена 10000 и более: отсутствуют.")
	}

	fmt.Printf("Общая сумма покупки: %d руб.\n", totalCost)
}
