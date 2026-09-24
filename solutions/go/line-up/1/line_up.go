package lineup

import "fmt"

func Format(name string, number int) string {
    switch {
    case number%10 == 1 && number % 100 != 11:
        return fmt.Sprintf("%s, you are the %dst customer we serve today. Thank you!", name, number)
    case number%10 == 2 && number % 100 != 12:
        return fmt.Sprintf("%s, you are the %dnd customer we serve today. Thank you!", name, number)
    case number%10 == 3 && number % 100 != 13:
        return fmt.Sprintf("%s, you are the %drd customer we serve today. Thank you!", name, number)
    default:
        return fmt.Sprintf("%s, you are the %dth customer we serve today. Thank you!", name, number)
    }
}
