package wordcount

import (
    "strings"
    "unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)
    freq := make(Frequency)

	splitter := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '\''
	}

	words := strings.FieldsFunc(phrase, splitter)

	for _, w := range words {
		word := strings.Trim(w, "'")
		if word != "" {
			freq[strings.ToLower(word)]++
		}
	}
	return freq
}
