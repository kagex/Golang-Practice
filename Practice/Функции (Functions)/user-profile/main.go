package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	name, age := "Александр", 29

	message, err := UserProfileToString(name, age)

	if err != nil {
		fmt.Printf("Получена ошибка от функции: %v", err)
	} else {
		fmt.Print(message)
	}
}

func UserProfileToString(name string, age int) (string, error) {

	if age < 0 {
		return "", errors.New("negative age")
	}

	if len(name) == 0 {
		return "", errors.New("empty name")
	}

	name = strings.TrimSpace(name)

	if len(name) == 0 {
		return "", errors.New("name cannot contain only spaces")
	}

	return fmt.Sprintf("Имя человека: %s, возраст: %d.", name, age), nil
}
