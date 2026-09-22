package luhn

import "strings"

func Valid(id string) bool {
	length := len(id) - 1
	if len(strings.TrimSpace(id)) < 1 {
		return false
	}
	sum := 0
	digits := 0
	needDouble := false
	for i := length; i >= 0; i-- {
		c := id[i]
		if c == ' ' {
			continue
		}
		if c < '0' || c > '9' {
			return false
		}
		value := int(c - '0')
		if needDouble {
			value *= 2

			if value > 9 {
				value -= 9
			}
		}
		sum += value
		digits++
		needDouble = !needDouble
	}
	return digits > 1 && sum%10 == 0
}