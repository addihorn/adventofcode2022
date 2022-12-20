package main

import (
	"example/hello/src/golang/aocutils"
	"strconv"
)

func main() {
	initialSequence := aocutils.ReadInput("test-input.txt")
	var newSequence = make([]int, len(initialSequence))

	for oldIndex, offsetAsString := range initialSequence {
		offset, _ := strconv.Atoi()
	}
}
