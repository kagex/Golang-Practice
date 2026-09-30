package bafflingbirthdays

import (
    "time"
	"math/rand/v2"
)
func SharedBirthday(dates []time.Time) bool {
    birthdays := map[string]int{}
        
	for _, t := range dates {
        date := t.Format("01-02")
        if _, ok := birthdays[date]; ok {
			return true
		}
		birthdays[date]++
    }      
    return false
}

func RandomBirthdates(size int) []time.Time {
    dates := make([]time.Time, size)
    for i:=0; i < size; i++ {
		year := notLeapYear()
        day := rand.IntN(365)+1
        dates[i] = time.Date(year, 1, day, 0, 0, 0, 0, time.UTC)
    }
    return dates
}

func EstimatedProbability(size int) float64 {
	const experiments = 10000
	matches := 0

	for i := 0; i < experiments; i++ {
		dates := RandomBirthdates(size)
		if SharedBirthday(dates) {
			matches++
		}
	}

	return float64(matches) / float64(experiments) * 100
}

func notLeapYear() int {
    currentYear := time.Now().Year()
	for {
        year := currentYear - rand.IntN(101)
        if year%4 != 0 || (year%100 == 0 && year%400 != 0) {
            return year
        }
    }
}