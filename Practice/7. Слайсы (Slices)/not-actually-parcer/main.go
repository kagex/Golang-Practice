package main

import (
	"fmt"
	"strings"
)

func main() {
	testString := "https://choose.your/destiny?user=&user=second"
	fmt.Print(GetUserParam(testString))
}

func GetUserParam(url string) string {
	urlLength := len(url) // Ваша переменная
	searchStart := 0      // Откуда начинать очередной поиск

	for {
		index := strings.Index(url[searchStart:], "user=") // Ваша переменная index

		if index == -1 {
			break
		}

		actualIndex := searchStart + index

		isValidParam := false
		if actualIndex == 0 {
			isValidParam = true
		} else {
			prevChar := url[actualIndex-1]
			if prevChar == '?' || prevChar == '&' {
				isValidParam = true
			}
		}

		if isValidParam {
			var sb strings.Builder

			for i := actualIndex + 5; i < urlLength; i++ {
				if url[i] == '&' || url[i] == '#' {
					break
				}
				sb.WriteByte(url[i])
			}

			value := sb.String()

			if value != "" {
				return value
			}
		}

		searchStart = actualIndex + 5
	}

	return "not found"
}
