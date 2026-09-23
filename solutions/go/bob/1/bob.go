package bob

import (
    "strings"
    "unicode"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {return "Fine. Be that way!"}

	isQuestion := remark[len(remark)-1] == '?'
	hasLetter := false
	isYelling := true
	for _, r := range remark {
		if unicode.IsLetter(r) {
			hasLetter = true
			if !unicode.IsUpper(r) {
				isYelling = false
				break
			}
		}
	}
	isYelling = isYelling && hasLetter

	switch {
	case isYelling && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isYelling:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}