package isbnverifier

import "strings"

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i, r := range isbn {
		if i == 9 && r == 'X' {
			sum += 10 * 1
			continue
		}
		if r < '0' || r > '9' {
			return false
		}
		weight := 10 - i
		sum += int(r-'0') * weight
	}

	return sum%11 == 0
}
