package main

import (
	"example/hello/src/golang/05/stack"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	startingPositions := createStartingConfiguration()
	fmt.Println(startingPositions, startingPositions[0].Peek(), startingPositions[2].Peek(), startingPositions[8].Peek())
	craneMoves := readInput("input.txt")

	// CrateMover 9000
	loadingZone9000 := startingPositions
	for _, craneMove := range craneMoves {
		craneMoveDetails := strings.Split(craneMove, " ")
		moveFrom, _ := strconv.Atoi(craneMoveDetails[3])
		moveTo, _ := strconv.Atoi(craneMoveDetails[5])
		moveNoCrates, _ := strconv.Atoi(craneMoveDetails[1])

		for i := 0; i < moveNoCrates; i++ {
			crateId := loadingZone9000[moveFrom-1].Pop()
			loadingZone9000[moveTo-1].Push(crateId)
			//fmt.Println("Moved Crate", crateId, "from", moveFrom, "to", moveTo)
		}
	}
	fmt.Println("Top of Stacks for CrateMover 9000")
	for _, crates := range loadingZone9000 {
		fmt.Print(crates.Peek())
	}
	fmt.Println()

	// CrateMover 9001
	loadingZone9001 := startingPositions
	for _, craneMove := range craneMoves {
		craneMoveDetails := strings.Split(craneMove, " ")
		moveFrom, _ := strconv.Atoi(craneMoveDetails[3])
		moveTo, _ := strconv.Atoi(craneMoveDetails[5])
		moveNoCrates, _ := strconv.Atoi(craneMoveDetails[1])

		bufferStack9001 := stack.New()
		for i := 0; i < moveNoCrates; i++ {
			createId := loadingZone9001[moveFrom-1].Pop()
			bufferStack9001.Push(createId)
		}

		for ok := true; ok; ok = (bufferStack9001.Peek() != nil) {
			createId := bufferStack9001.Pop()
			loadingZone9001[moveTo-1].Push(createId)
			//fmt.Println("Moved Crate", crateId, "from", moveFrom, "to", moveTo)
		}

	}

	fmt.Println("Top of Stacks for CrateMover 9001")
	for _, crates := range loadingZone9001 {
		fmt.Print(crates.Peek())
	}
	fmt.Println()

}

func readInput(fileName string) []string {

	b, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	fileAsString := string(b)
	return strings.Split(fileAsString, "\n")
}

func createStartingConfiguration() [9]stack.Stack {

	startingPostitionsPerRow := readInput("startingPositions.txt")
	fmt.Println(startingPostitionsPerRow[0])

	valueIndices := []int{1, 5, 9, 13, 17, 21, 25, 29, 33}
	var startingPositions [9]stack.Stack

	for row := len(startingPostitionsPerRow) - 2; row >= 0; row-- {
		for i, index := range valueIndices {
			//fmt.Println(index, string(startingPostitionsPerRow[row][index]))
			valueToPush := startingPostitionsPerRow[row][index]
			if valueToPush != 32 {
				startingPositions[i].Push(string(valueToPush))
			}

		}
	}
	return startingPositions

}
