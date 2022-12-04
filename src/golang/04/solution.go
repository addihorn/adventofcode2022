package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	pairings := readInput("input.txt")
	fmt.Println(pairings)

	fullyContainedPairings := 0
	overlappingPairings := 0
	for _, pairing := range pairings {
		worksationsPerElve := strings.Split(pairing, ",")

		elve1Bounds := strings.Split(worksationsPerElve[0], "-")
		elve2Bounds := strings.Split(worksationsPerElve[1], "-")

		elve1LB, _ := strconv.Atoi(elve1Bounds[0])
		elve1UB, _ := strconv.Atoi(elve1Bounds[1])

		elve2LB, _ := strconv.Atoi(elve2Bounds[0])
		elve2UB, _ := strconv.Atoi(elve2Bounds[1])

		if (elve1LB <= elve2LB && elve1UB >= elve2UB) || //elve 2 fully contained in elve 1
			(elve2LB <= elve1LB && elve2UB >= elve1UB) {
			fullyContainedPairings += 1
		}

		if (elve1LB <= elve2LB && elve1UB >= elve2UB) || //elve 2 fully contained in elve 1
			(elve2LB <= elve1LB && elve2UB >= elve1UB) || //elve 1 fully contained in elve 2
			(elve1UB >= elve2LB && elve2UB >= elve1UB) || //elve 2 overlaps at end of elve 1
			(elve2UB >= elve1LB && elve1UB >= elve2UB) {
			overlappingPairings += 1
		}
		fmt.Println(worksationsPerElve, fullyContainedPairings, overlappingPairings)

	}

}

func readInput(fileName string) []string {

	b, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	fileAsString := string(b)
	return strings.Split(fileAsString, "\n")
}
