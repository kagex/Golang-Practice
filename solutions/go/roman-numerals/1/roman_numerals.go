package romannumerals

import (
	"errors"
	"strings"
)

// romanMapping хранит пары (значение → римская цифра)
// ВАЖНО: массив должен быть отсортирован по убыванию значений!
var romanMapping = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("input must be between 1 and 3999")
	}

	var result strings.Builder

	for _, mapping := range romanMapping {
		for input >= mapping.value {
			result.WriteString(mapping.symbol)
			input -= mapping.value
		}
	}

	return result.String(), nil
}