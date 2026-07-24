package main

import (
	"fmt"
	"strings"
)

func main() {
	PrintReplaced("Утку, я люблю.")
}

// PrintReplaced заменяет все вхождения буквы 'у' на букву 'а' в переданной строке
// strings.Builder для эффективного посимвольного построения результата
// Цикл проходит по каждой руне входной строки, заменяет 'у' на 'а'
func PrintReplaced(input string) {
	// Создает пустой strings.Builder для накопления результирующей строки
	var sb strings.Builder

	for _, ch := range input {
		// Проверяет, равен ли текущий символ букве 'у'
		if ch == 'у' {
			ch = 'а'
		}
		// Добавляет текущий символ в буфер
		sb.WriteRune(ch)
	}

	// Вывод накопленной в буфере строки
	fmt.Print(sb.String())
}
