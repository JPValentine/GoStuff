package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count:=0
	for i:=0;len(birdsPerDay)>i;i++{
		count+=birdsPerDay[i]
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	wEnd:=(week*7)
	wStart:=wEnd-7
	birdWeek:=birdsPerDay[wStart:wEnd]
	return TotalBirdCount(birdWeek) 
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i:=0;len(birdsPerDay)>i;i++{
		if i%2==0 || i==0{
			birdsPerDay[i]=birdsPerDay[i]+1
		}
	}//for
	return birdsPerDay
}
