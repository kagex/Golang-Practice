package largestseriesproduct

import (
    "errors"
	"unicode"
)
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) || span < 0{
        return 0, errors.New("Error")
    }
    
	for _, r := range digits {
		if !unicode.IsDigit(r) {
			return 0, errors.New("digits input must only contain digits")
		}
	}

	var maxProduct int64 = 0
	for i := 0; i <= len(digits)-span; i++ {
		var product int64 = 1
		for j := i; j < i+span; j++ {
			product *= int64(digits[j] - '0')
		}
		if product > maxProduct {
			maxProduct = product
		}
	}
	return maxProduct, nil
}
