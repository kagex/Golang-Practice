package main

import (
	"fmt"
	"strings"
)

func main() {

	message, formattedString, containsPrefix := stringFormater("Go - это не просто язык, это СТИЛЬ ЖИЗНИ!")

	fmt.Println(message)
	fmt.Println(formattedString)
	fmt.Println(containsPrefix)
}

func stringFormater(message string) (string, string, bool) {
	message = strings.TrimSpace(message)                       // Обрезка пробелов в начале и конце строки
	formattedString := strings.ToLower(message)                // Перевод строки в нижний регистр
	containsPrefix := strings.HasPrefix(formattedString, "go") // Проверка что строка в нижнем регистре начинается с подстроки

	return message, formattedString, containsPrefix
}
