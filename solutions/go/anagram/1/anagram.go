package anagram

import (
    "strings"
	"slices"
)
func Detect(subject string, candidates []string) []string {
	result := make([]string, 0, len(candidates))
	lowerSubject := strings.ToLower(subject)
	subjectRunes := []rune(lowerSubject)
	slices.Sort(subjectRunes)

	for _, candidate := range candidates {
		lowerCandidate := strings.ToLower(candidate)

		if len(subject) != len(candidate) {
			continue
		}
		if lowerCandidate == lowerSubject {
			continue
		}
		candidateRunes := []rune(lowerCandidate)
		slices.Sort(candidateRunes)
		if slices.Equal(subjectRunes, candidateRunes) {
			result = append(result, candidate)
		}
	}
	return result
}
