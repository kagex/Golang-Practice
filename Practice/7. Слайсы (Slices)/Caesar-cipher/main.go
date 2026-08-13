package main

import (
	"fmt"
)

func main() {
	str := "Зашифруй меня!"
	fmt.Println("Первоначальное сообщение:", str)
	encodedStr := CaesarCode(str, 5, true)
	fmt.Println("Зашифрованное сообщение: ", encodedStr)

	decodedStr := CaesarCode(encodedStr, 5, false)
	fmt.Println("\nРасшифрованное сообщение: ", decodedStr)
}

func CaesarCode(text string, shift int, encode bool) string {
	var result []rune
	switch encode {
	case true:
		for _, v := range text {
			result = append(result, v+rune(shift))
		}
	case false:
		for _, v := range text {
			result = append(result, v-rune(shift))
		}
	}
	return string(result)
}
