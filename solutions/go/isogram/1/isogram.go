package isogram

import "unicode"

func IsIsogram(word string) bool {
	counts := make(map[rune]int)
    for _, v := range word {
        value := unicode.ToLower(v)
        counts[value]++
        if counts[value] > 1 && unicode.IsLetter(value){
            return false
        }
    }
    return true
}
