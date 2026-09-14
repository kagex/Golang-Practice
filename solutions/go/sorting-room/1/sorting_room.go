package sorting

import (
    "strconv"
    "fmt"
)
// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fn, ok := fnb.(FancyNumber)
	
	if !ok {
		return 0
	}
    
	number, _ := strconv.Atoi(fn.Value())
	
	return number
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
    switch v := i.(type) {
        case int:
        // int нужно привести к float64, потому что DescribeNumber принимает float64
        return DescribeNumber(float64(v))

        case float64:
        // float64 уже подходит
        return DescribeNumber(v)

        case NumberBox:
        // v здесь имеет тип NumberBox (интерфейс)
        return DescribeNumberBox(v)

        case FancyNumberBox:
        // v здесь имеет тип FancyNumberBox (интерфейс)
        return DescribeFancyNumberBox(v)

        default:
        // Ни один тип не подошёл
        return "Return to sender"
    }
}
