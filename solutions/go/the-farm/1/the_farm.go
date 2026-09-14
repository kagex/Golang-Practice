package thefarm

import (
    "errors"
    "fmt"
) 

// TODO: define the 'DivideFood' function
func DivideFood (weight FodderCalculator, cowsAmount int) (float64, error) {
    amount, err := weight.FodderAmount(cowsAmount)

    if err != nil {
        return 0, err
    }

    factor, err := weight.FatteningFactor()
    if err != nil {
        return 0, err
    }

    return (amount * factor) / float64(cowsAmount), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(weight FodderCalculator, cows int) (float64, error) {
    if cows <= 0 {
        // Создаем простую ошибку с заданным текстом
        return 0, errors.New("invalid number of cows")
    }
    
    // Если проверка пройдена, вызываем функцию из первого шага
    return DivideFood(weight, cows)
}

// TODO: define the 'ValidateNumberOfCows' function
// 1. Объявляем структуру ошибки
type InvalidCowsError struct {
	cows    int
	message string
}

// 2. Реализуем интерфейс error (обязательно через указатель *InvalidCowsError)
func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// 3. Функция валидации
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "there are no negative cows",
		}
	}
	if cows == 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "no cows don't need food",
		}
	}
	
	// Если всё хорошо, возвращаем nil (ошибки нет)
	return nil
}
