package diamond

import (
    "strings"
    "errors"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("invalid character")
	}

	n := int(char - 'A')
	rows := make([]string, 0, 2*n+1)

	for row := 0; row <= 2*n; row++ {
		i := row
		if row > n {
			i = 2*n - row
		}

		outSpaces := strings.Repeat(" ", n-i)
		letter := byte('A' + i)

		if i == 0 {
			rows = append(rows, outSpaces+string(letter)+outSpaces)
		} else {
			inSpaces := strings.Repeat(" ", 2*i-1)
			rows = append(rows, outSpaces+string(letter)+inSpaces+string(letter)+outSpaces)
		}
	}

	return strings.Join(rows, "\n"), nil
}