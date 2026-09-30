package affinecipher

import (
    "errors"
	"strings"
    "unicode"
)
func IsCoprime(a, b int) bool {
	for b != 0 {
		a, b = b, a%b
	}

	return a == 1
}

func Encode(text string, a, b int) (string, error) {
    if !IsCoprime(a, 26) {
        return "", errors.New("a not coprime to m")
    }

    text = strings.ToLower(text)
	var sb strings.Builder
    count := 0
    
    for _, v := range text {
		if unicode.IsLetter(v) {
            if count == 5 {
				sb.WriteRune(' ')
				count = 0
			}
			encodedRune := rune((a*int(v-'a')+b)%26) + 'a'
            sb.WriteRune(encodedRune)
            count++
        }

        if unicode.IsDigit(v) {
            if count == 5 {
				sb.WriteRune(' ')
				count = 0
			}
            sb.WriteRune(v)
            count++
        }
    }
    return sb.String(), nil
}

func modInverse(a, m int) int {
	a = ((a % m) + m) % m
	for i := 1; i < m; i++ {
		if (a*i)%m == 1 {
			return i
		}
	}
	return -1
}

func Decode(text string, a, b int) (string, error) {
	if !IsCoprime(a, 26) {
		return "", errors.New("a not coprime to m")
	}

	aInv := modInverse(a, 26)
	var sb strings.Builder

	for _, v := range text {
		if unicode.IsLetter(v) {
			n := (aInv * (int(v-'a') - b)) % 26
			if n < 0 {
				n += 26
			}
			sb.WriteRune(rune(n) + 'a')
		} 
        if unicode.IsDigit(v) {
			sb.WriteRune(v)
		}
	}

	return sb.String(), nil
}
