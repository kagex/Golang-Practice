package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Здравствуйте, я Анализатор текста. Напечатайте или вставьте текст, и я проанализирую.")
	userInput, err := GetInput()

	if err != nil {
		fmt.Println("Возникла ошибка с введенным вами текстом, попробуйте еще раз!")
		GetInput()
		return
	}

	letters, digits, spaces, punctuation := CountCharacters(userInput)

	DisplayResults(letters, digits, spaces, punctuation)
}

func GetInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	userInput, err := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)

	return userInput, err
}

func CountCharacters(text string) (letters, digits, spaces, punctuation int) {
	for _, v := range text {
		switch {
		case unicode.IsLetter(v):
			letters++
		case unicode.IsDigit(v):
			digits++
		case unicode.IsSpace(v):
			spaces++
		case unicode.IsPunct(v):
			punctuation++
		}
	}
	return
}

func DisplayResults(letters, digits, spaces, punctuation int) {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Количество букв: %d\n", letters)
	fmt.Fprintf(&sb, "Количество цифр: %d\n", digits)
	fmt.Fprintf(&sb, "Количество пробелов: %d\n", spaces)
	fmt.Fprintf(&sb, "Количество знаков препинания: %d\n", punctuation)

	fmt.Println(sb.String())
}
