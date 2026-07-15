package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	stringForParsing := "𝓗𝓮𝓵𝓵𝓸, мой друг."

	result, bytes, length := about(stringForParsing)
	fmt.Printf("Строка: %q\n", result)
	fmt.Printf("Байты: %d\n", bytes)
	fmt.Printf("Длина строки: %d\n", length)
}

func about(str string) (string, int, int) {
	length := utf8.RuneCountInString(str) // Получаем количество символов в строке
	bytes := len(str)                     // Получаем количество байт в строке

	return str, bytes, length // Возвращаем строку и две переменные указанные раннее
}
