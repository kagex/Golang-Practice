package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	var acronym strings.Builder
	lastWasLetter := false

	for _, r := range s {
		switch {
		case unicode.IsLetter(r):
			if !lastWasLetter {
				acronym.WriteRune(r)
			}
			lastWasLetter = true
		case r == '\'' || r == '’':
		default:
			lastWasLetter = false
		}
	}

	return strings.ToUpper(acronym.String())
}