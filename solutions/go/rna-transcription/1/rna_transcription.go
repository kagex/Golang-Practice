package rnatranscription

var transcribe = map[rune]rune {
    'G':'C',
    'C':'G',
    'T':'A',
    'A':'U',
}
func ToRNA(dna string) string {
	runes := []rune(dna)
    result := ""
    for _, v := range runes {
        result += string(transcribe[v])
    }
    return result
}
