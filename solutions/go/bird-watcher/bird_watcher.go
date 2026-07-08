package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    count := 0
	for i:= 0; i < len(birdsPerDay); i++ {
        count += birdsPerDay[i]
    }
    return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	count := 0
    startIndex := 7 * (week - 1)
    endIndex := 7 + startIndex
    
    for i:= startIndex; i < endIndex; i++ {
        count += birdsPerDay[i]
    }
    return count
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    
    for i:= 0; i < len(birdsPerDay)/2; i++ {
        birdsPerDay[i*2] += 1 
    }
	return birdsPerDay
}
