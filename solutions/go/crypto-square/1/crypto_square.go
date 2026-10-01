package cryptosquare

import (
    "unicode"
    "strings"
)
func Encode(pt string) string {
    var sb strings.Builder
	for _, v := range pt {
        if unicode.IsLetter(v) || unicode.IsDigit(v) {
			sb.WriteRune(unicode.ToLower(v))
        }
    }
    normalized := sb.String()
    length:= len(normalized)
        
    c := 1
    for c*c < length {
    	c++
    }
    r := (length + c - 1) / c
    
    sb.Reset()
    for j := 0; j < c; j++ {
    	if j > 0 {
    		sb.WriteByte(' ')
    	}
    	for i := 0; i < r; i++ {
    		index := j + i*c
    		if index < length {
    			sb.WriteByte(normalized[index])
    		} else {
    			sb.WriteByte(' ')
    		}
    	}
    }
    return sb.String()
}
