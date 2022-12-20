package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"strconv"
)

func main() {

	initialSequence := aocutils.ReadInput("input.txt")
	var newIndizes = make(map[int]int, len(initialSequence))
	var newSequence = make([]int, len(initialSequence))
	//initialize new sequence

	for index, _ := range initialSequence {
		//offset, _ := strconv.Atoi(offsetAsString)
		newIndizes[index] = index
	}

	//fmt.Println(newIndizes)
	for initialIndex, offsetAsString := range initialSequence {
		offset, _ := strconv.Atoi(offsetAsString)
		rollOverSanitizedOffset := offset % len(initialSequence)
		//fmt.Println(rollOverSanitizedOffset)
		if rollOverSanitizedOffset < 0 {
			rollOverSanitizedOffset = len(initialSequence) - 1 + rollOverSanitizedOffset
		}
		newIndex := newIndizes[initialIndex] + rollOverSanitizedOffset

		if newIndex > len(initialSequence) {
			newIndex = newIndex - (len(initialSequence) - 1)
		}

		oldIndex := newIndizes[initialIndex]

		for x, _ := range initialSequence {
			validateIndex := x
			if x == initialIndex {
				newIndizes[validateIndex] = newIndex
				continue
			}

			if newIndizes[validateIndex] < oldIndex && newIndizes[validateIndex] > newIndex {
				newIndizes[validateIndex]++
				continue
			}
			if newIndizes[validateIndex] > oldIndex && newIndizes[validateIndex] < newIndex {
				newIndizes[validateIndex]--
				continue
			}

		}

		for index, value := range initialSequence {
			newIndex := newIndizes[index]
			offset, _ := strconv.Atoi(value)
			newSequence[newIndex] = offset
		}

		//fmt.Println(newSequence)
	}

	indexZero := 0
	for index, value := range newSequence {
		if value == 0 {
			indexZero = index
			break
		}
	}

	indizesToCheck := [3]int{(indexZero + 1000) % len(initialSequence), (indexZero + 2000) % len(initialSequence), (indexZero + 3000) % len(initialSequence)}
	fmt.Println(indexZero, indizesToCheck)
	fmt.Println(newSequence[indizesToCheck[0]], newSequence[indizesToCheck[1]], newSequence[indizesToCheck[2]], "=",
		newSequence[indizesToCheck[0]]+newSequence[indizesToCheck[1]]+newSequence[indizesToCheck[2]])

}
