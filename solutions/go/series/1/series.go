package series

func All(n int, s string) []string {
	if n <= 0 || n > len(s) {
		return []string{}
	}

	substrings := make([]string, 0, len(s)-n+1)
	for i := 0; i+n <= len(s); i++ {
		substrings = append(substrings, s[i:i+n])
	}
	return substrings
}

func UnsafeFirst(n int, s string) string {
	substrings := All(n,s)
    return substrings[0]
}
