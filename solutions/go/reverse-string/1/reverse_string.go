package reversestring

func Reverse(input string) string {
	runes := []rune(input)
    result := ""

    for i:=len(runes)-1; i >= 0; i-- {
        result += string(runes[i])
    }
    return result
}
