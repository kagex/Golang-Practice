package twelvedays

import (
    "fmt"
    "strings"
)
    
var days = []string{
    "first",
    "second",
    "third",
    "fourth",
    "fifth",
    "sixth",
    "seventh",
    "eighth",
    "ninth",
    "tenth",
    "eleventh",
    "twelfth",
}

var gifts = []string{
    "a Partridge in a Pear Tree",
    "two Turtle Doves",
    "three French Hens",
    "four Calling Birds",
    "five Gold Rings",
    "six Geese-a-Laying",
    "seven Swans-a-Swimming",
    "eight Maids-a-Milking",
    "nine Ladies Dancing",
    "ten Lords-a-Leaping",
    "eleven Pipers Piping",
    "twelve Drummers Drumming",
}

func Verse(i int) string {
    i = i-1
    var sb strings.Builder
    fmt.Fprintf(&sb, "On the %s day of Christmas my true love gave to me: ", days[i])
    for pos:=i; pos >= 0; pos-- {
        switch {
        case pos == 0 && i == 0:
            fmt.Fprintf(&sb, "%s.", gifts[pos])
        case pos == 0:
            fmt.Fprintf(&sb, "and %s.", gifts[pos])
        default:
            fmt.Fprintf(&sb, "%s, ", gifts[pos])
        }   
	}
    return sb.String()
}

func Song() string {
    var sb strings.Builder
	for i:= 1; i <= 12; i++ {
        fmt.Fprintf(&sb, "%s", Verse(i))
    	if i != 12 {
			sb.WriteString("\n")
        }
    }
    return sb.String()
}
