package main

import (
	"fmt"
	"strings"
)

func main() {
	printDiamond(10)
}

func printDiamond(n int) {
	var sb strings.Builder

	sb.WriteString("Мой бриллиант:\n")
	// Количество строк в ромбе (нашем бриллианте) равно 2*n-1
	for row := 0; row < 2*n-1; row++ {
		i := row //индекс для симметрии

		if row >= n { // После середины ромба отражаем индекс, чтобы ромб был симметричен
			i = 2*n - 2 - row
		}

		outSpaces := strings.Repeat(" ", n-1-i) // Количество пробелов слева

		if i == 0 {
			// Самый верх (i=0) и самый низ (i=0):
			// печатаем внешние пробелы и ровно одну решетку
			fmt.Fprintf(&sb, "%s#\n", outSpaces)
		} else {
			// Все остальные строки:
			// вычисляем нечетное количество внутренних пробелов (2*i - 1)
			// печатаем решетки по краям
			inSpaces := strings.Repeat(" ", 2*i-1)
			fmt.Fprintf(&sb, "%s#%s#\n", outSpaces, inSpaces)
		}
	}
	// Выводим всю собранную в буфере строку
	fmt.Print(sb.String())
}
