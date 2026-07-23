package main

import "fmt"

func main() {
	// Данные пользователя
	age := 18
	var role string = "user"
	var status string = "active"

	// Доступ разрешен, если:
	// 1. Роль admin или moderator — доступ в любом случае
	// 2. Роль user, возраст >= 18 и статус активный
	access := (role == "admin" || role == "moderator") || (age >= 18 && status == "active" && role == "user")

	fmt.Printf("Пользователь доступ: %t", access)
}
