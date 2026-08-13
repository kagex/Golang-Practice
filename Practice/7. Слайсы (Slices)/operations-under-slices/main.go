package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"slices"
)

func main() {
	var desiredSize int
	fmt.Println("Ты любишь слайсить? :) \nУкажи размер слайса, а я сгенерирую")
	fmt.Scanln(&desiredSize)

	createdSlice, err := createSlice(desiredSize)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("Сгенерирован слайс размером %d:\n%v", desiredSize, createdSlice)
	fmt.Printf("\n\nСлайс после форматирования:\n%v", filterSlice(createdSlice))

	sizeOfSubSlice := 3
	fmt.Printf("\n\nПодслайс размером %d с отрицательным числом и максимальной суммой элементов:\n%v", sizeOfSubSlice, maxSumWithNegative(createdSlice, sizeOfSubSlice))

	fmt.Printf("\n\nОтсортированный слайс:\n%v", sortByParity(createdSlice))
}

func createSlice(desiredSize int) ([]int, error) {
	switch {
	case desiredSize < 0:
		return []int{}, errors.New("Невозможно создать слайс с отрицательным размером")
	case desiredSize == 0:
		return []int{}, nil
	}

	createdSlice := make([]int, desiredSize)
	for position := range createdSlice {
		createdSlice[position] = rand.IntN(21) - 10
	}
	return createdSlice, nil
}

func filterSlice(numbers []int) []int {
	result := []int{}

	for i := 1; i < len(numbers); i++ {
		temp := numbers[i]
		if temp < numbers[i-1] && (temp%2 == 0 || temp%5 == 0 || temp%6 == 0 || temp%9 == 0) {
			result = append(result, temp)
		}
	}
	return result
}

func maxSumWithNegative(numbers []int, k int) []int {
	if k <= 0 || k > len(numbers) {
		return nil
	}

	startPos := 0
	sum := 0
	negativeExist := false
	negativeInSlices := false

	var maxSum int

	// Инициализация первого окна (собираем первые k элементов)
	for i := 0; i < k; i++ {
		sum += numbers[i]
		if numbers[i] < 0 {
			negativeExist = true
		}
	}

	// Если в первом окне есть отрицательное число, сохраняем его как текущий лучший результат
	if negativeExist {
		maxSum = sum
		negativeInSlices = true
	}

	// Двигаем окно вправо (скользящее окно)
	for i := k; i < len(numbers); i++ {
		// Убираем элемент, который остался позади (слева)
		outElement := numbers[i-k]
		sum -= outElement

		// Если ушедший элемент был отрицательным, нам нужно проверить, есть ли другие отрицательные числа в оставшемся окне.
		if outElement < 0 {
			negativeExist = false // Сбрасываем флаг и проверяем окно заново
			for j := i - k + 1; j <= i; j++ {
				if numbers[j] < 0 {
					negativeExist = true
					break
				}
			}
		}

		// Добавляем новый элемент, который зашел в окно справа
		inElement := numbers[i]
		sum += inElement
		if inElement < 0 {
			negativeExist = true // Если пришло отрицательное, флаг точно true
		}

		// Обновляем максимум, если текущее окно подходит и его сумма больше
		// Строгое сравнение (>) гарантирует, что при равных суммах останется самый первый (левый) подслайс.
		if negativeExist {
			if !negativeInSlices || sum > maxSum {
				maxSum = sum
				startPos = i - k + 1 // Сохраняем позицию начала лучшего слайса
				negativeInSlices = true
			}
		}
	}

	if !negativeInSlices {
		return nil // Если не нашли ни одного подходящего слайса
	}

	return numbers[startPos : startPos+k]
}

func sortByParity(numbers []int) []int {
	result := make([]int, len(numbers))
	copy(result, numbers)

	slices.SortFunc(result, func(a, b int) int {
		aEven := a%2 == 0
		bEven := b%2 == 0

		// Если оба числа чётные — сортируем по убыванию.
		if aEven && bEven {
			return b - a
		}

		// Если оба числа нечётные — сортируем по возрастанию.
		if !aEven && !bEven {
			return a - b
		}

		// Если a чётное, а b нечётное — a должно быть раньше.
		if aEven {
			return -1
		}

		// Если a нечётное, а b чётное — b должно быть раньше.
		return 1
	})

	return result
}
