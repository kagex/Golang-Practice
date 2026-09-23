package eliudseggs

func EggCount(displayValue int) int {
	count := 0
	for displayValue > 0 {
		displayValue &= displayValue - 1
		count++
	}
	return count
}
