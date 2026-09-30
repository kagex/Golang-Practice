package allergies

import "slices"

var allergens = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(allergies uint) []string {
    result := []string{} 
	for i, v := range allergens {
        if allergies & (1 << i) != 0 {
            result = append(result, v)
        }
    }
    return result
}

func AllergicTo(allergies uint, allergen string) bool {
    index := slices.Index(allergens, allergen)
    if index < 0 {
        return false
    }
    return allergies & (1 << index) != 0 
}
