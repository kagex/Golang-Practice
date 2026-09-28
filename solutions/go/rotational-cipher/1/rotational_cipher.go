package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	cipher := make([]rune, 0, len(plain))
	for _, r := range plain {
		switch {
		case unicode.IsUpper(r):
			r = rune('A' + (int(r-'A')+shiftKey)%26)
		case unicode.IsLower(r):
			r = rune('a' + (int(r-'a')+shiftKey)%26)
		}
		cipher = append(cipher, r)
	}
	return string(cipher)
}
