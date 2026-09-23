package atbashcipher

import (
    "strings"
	"unicode"
)
func Atbash(s string) string {
	var sb strings.Builder
	count := 0 

	for _, r := range strings.ToLower(s) {
		switch {
		case r >= 'a' && r <= 'z':
			cipherRune := 'z' - (r - 'a')
			if count == 5 {
				sb.WriteRune(' ')
				count = 0
			}
			sb.WriteRune(cipherRune)
			count++
		case unicode.IsDigit(r):
			if count == 5 {
				sb.WriteRune(' ')
				count = 0
			}
			sb.WriteRune(r)
			count++
		}
	}
	return sb.String()
}
