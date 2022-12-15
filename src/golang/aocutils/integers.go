package aocutils

import "sort"

func min(intValues []int) int {
	sort.Ints(intValues)
	return intValues[0]
}

func max(intValues []int) int {
	sort.Ints(intValues)
	return intValues[len(intValues)-1]
}

func Abs(value int) int {
	if value < 0 {
		return value * -1
	}
	return value
}
