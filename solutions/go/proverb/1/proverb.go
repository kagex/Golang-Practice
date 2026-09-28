package proverb

import "fmt"

func Proverb(rhyme []string) []string {
    proverbStrings := []string{}
    length := len(rhyme)
    
    if length == 0 {
        return proverbStrings
    }

    for i:= 0; i < length; i++ {
        if i == length - 1 {
        	proverbStrings = append(proverbStrings, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
   		} else {
        	proverbStrings = append(proverbStrings, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i],rhyme[i+1]))
        }
    }
    return proverbStrings
}
