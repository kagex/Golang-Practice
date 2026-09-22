package pangram

import "unicode"

func IsPangram(input string) bool {
    letters := make(map[rune]int)
    for _, v := range input {
        if !unicode.IsLetter(v) {continue}
        letters[unicode.ToLower(v)]++
	}
    return len(letters) == 26
}