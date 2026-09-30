package binarysearch

func SearchInts(list []int, key int) int {
    if len(list) == 0 {
		return -1
	}

	leftBorder := 0
	rightBorder := len(list) - 1

	for leftBorder <= rightBorder {
		mid := leftBorder + (rightBorder-leftBorder)/2

		switch {
		case list[mid] == key:
			return mid
		case list[mid] < key:
			leftBorder = mid + 1
		default:
			rightBorder = mid - 1
		}
	}
    return -1
}
