package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Гласные буквы русского алфавита, разделённые запятыми
const vowels = "а, е, ё, и, о, у, ы, э, ю, я"

func main() {
	testArray := []string{"АЕИОУ", "кхм", ""}
	CountVowelsInArray(testArray)
}

// CountVowelsInArray подсчитывает количество гласных букв в каждом элементе массива
func CountVowelsInArray(arr []string) {
	// Фиксированный массив из 3 элементов для хранения результатов
	result := ""

	var counter int

	// Проходим по каждому элементу массива с индексом
	for _, word := range arr {
		counter = 0
		// Для каждой буквы слова (в нижнем регистре) проверяем, является ли она гласной
		for _, letter := range strings.ToLower(word) {
			if strings.ContainsRune(vowels, letter) {
				counter++
			}
		}
		// Сохраняем результат для текущего элемента
		result += strconv.Itoa(counter) + " "
	}
	fmt.Println(result)
}
