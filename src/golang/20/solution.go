package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"strconv"
)

type EncryptionValue struct {
	CurrentPosition int
	OffseValue      int64
}

func main() {

	// Part2
	p2decryptionKey := int64(811589153)
	p2NumberOfMixes := 10

	//Part 1 with Part2-Code
	//p2decryptionKey := int64(1)
	//p2NumberOfMixes := 1

	initialSequence := aocutils.ReadInput("input.txt")
	var newSequence = make([]*EncryptionValue, len(initialSequence))
	var initialIndizes = make(map[int]*EncryptionValue)
	//initialize new sequence

	for index, value := range initialSequence {
		//offset, _ := strconv.Atoi(offsetAsString)
		offset, _ := strconv.Atoi(value)
		encryptionValue := &EncryptionValue{index, int64(offset) * p2decryptionKey}
		newSequence[index] = encryptionValue
		initialIndizes[index] = encryptionValue
	}

	//fmt.Println(newIndizes)
	for x := 0; x < p2NumberOfMixes; x++ {
		for initialIndex, _ := range initialSequence {

			encryptionValue := initialIndizes[initialIndex]
			//fmt.Println(rollOverSanitizedOffset)

			newIndex := int64(encryptionValue.CurrentPosition) + encryptionValue.OffseValue

			for newIndex < 0 || newIndex >= int64(len(initialSequence)) {
				if newIndex >= int64(len(initialSequence)) {
					overflowCorrection := newIndex / int64(len(initialSequence))
					newIndex = newIndex%int64(len(initialSequence)) + overflowCorrection
				}
				if newIndex < 0 {
					overflowCorrection := newIndex/int64(len(initialSequence)) - 1
					newIndex = int64(len(initialSequence)) + newIndex%int64(len(initialSequence)) + overflowCorrection
				}
			}

			newIndexInt := int(newIndex)
			oldIndex := encryptionValue.CurrentPosition

			firstHalf := make([]*EncryptionValue, oldIndex)
			secondHalf := make([]*EncryptionValue, len(newSequence)-1-oldIndex)

			firstHalfNewIndex := make([]*EncryptionValue, newIndex)
			secondHalfNewIndex := make([]*EncryptionValue, len(newSequence)-1-newIndexInt)

			copy(firstHalf, newSequence[0:oldIndex])
			copy(secondHalf, newSequence[oldIndex+1:len(newSequence)])

			sliceWithoutElement := append(firstHalf, secondHalf...)

			copy(firstHalfNewIndex, sliceWithoutElement[0:newIndex])
			copy(secondHalfNewIndex, sliceWithoutElement[newIndex:len(sliceWithoutElement)])

			firstHalfNewIndex = append(firstHalfNewIndex, []*EncryptionValue{encryptionValue}...)
			newSequence = append(firstHalfNewIndex, secondHalfNewIndex...)
			//fmt.Println(newSequence)
			for newIndex, this := range newSequence {
				this.CurrentPosition = newIndex
				//fmt.Print(this.OffseValue, " ")
			}
			//fmt.Println()
		}
	}

	indexZero := 0
	for index, value := range newSequence {
		if value.OffseValue == 0 {
			indexZero = index
			break
		}
	}

	indizesToCheck := [3]int{(indexZero + 1000) % len(initialSequence), (indexZero + 2000) % len(initialSequence), (indexZero + 3000) % len(initialSequence)}
	fmt.Println(indexZero, indizesToCheck)
	fmt.Println(newSequence[indizesToCheck[0]].OffseValue, newSequence[indizesToCheck[1]].OffseValue, newSequence[indizesToCheck[2]].OffseValue, "=",
		newSequence[indizesToCheck[0]].OffseValue+newSequence[indizesToCheck[1]].OffseValue+newSequence[indizesToCheck[2]].OffseValue)

}
