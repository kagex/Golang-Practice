package phonenumber

import (
	"errors"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	digits := make([]rune, 0, 11)
	
	for _, r := range phoneNumber {
		if unicode.IsDigit(r) {
			digits = append(digits, r)
		} else if unicode.IsLetter(r) {
			return "", errors.New("letters not permitted")
		} else if r == '+' || r == '(' || r == ')' || r == '-' || r == '.' || r == ' ' {
			continue
		} else {
			return "", errors.New("punctuations not permitted")
		}
	}

	length := len(digits)
	if length < 10 {
		return "", errors.New("must not be fewer than 10 digits")
	}
	if length > 11 {
		return "", errors.New("must not be greater than 11 digits")
	}

	offset := 0
	if length == 11 {
		if digits[0] != '1' {
			return "", errors.New("11 digits must start with 1")
		}
		offset = 1
	}

	if digits[offset] == '0' {
		return "", errors.New("area code cannot start with zero")
	}
	if digits[offset] == '1' {
		return "", errors.New("area code cannot start with one")
	}

	if digits[offset+3] == '0' {
		return "", errors.New("exchange code cannot start with zero")
	}
	if digits[offset+3] == '1' {
		return "", errors.New("exchange code cannot start with one")
	}

	return string(digits[offset:]), nil
}

func AreaCode(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return num[:3], nil
}

func Format(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return "(" + num[:3] + ") " + num[3:6] + "-" + num[6:], nil
}