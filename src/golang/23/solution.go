package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"time"

	"github.com/inancgumus/screen"
)

var gridSize *aocutils.Gridsize
var orderOfProposal [4]rune
var elfMoved bool

// const inputFile = "test-input2.txt"
const inputFile = "input.txt"

func main() {

	numberOfRounds := math.MaxInt

	gridSize = aocutils.NewGridSize()
	orderOfProposal = [4]rune{'N', 'S', 'W', 'E'}
	elfList = map[[2]int]*Elf{}

	inputData := aocutils.ReadInput(inputFile)
	for row, inputRow := range inputData {
		for col, tileData := range inputRow {
			if tileData == '#' {
				newElfStart := [2]int{col + 1, row + 1}
				elfList[newElfStart] = NewElf(newElfStart)
			}
		}
	}

	paintGrid()
	for i := 0; i < numberOfRounds; i++ {
		startNewRound()

		//each elf porposes his tile
		for _, elf := range elfList {
			elf.ProposeTile()
		}

		//each elf moves to proposed tile
		for _, elf := range elfList {
			elf.MoveToProposedTile()
		}

		//order of proposal rotates
		firstProposal := orderOfProposal[0]
		for i := 0; i < len(orderOfProposal)-1; i++ {
			orderOfProposal[i] = orderOfProposal[i+1]
		}
		orderOfProposal[len(orderOfProposal)-1] = firstProposal
		//paintGrid()
		//fmt.Println(len(elfList))

		if !elfMoved {
			fmt.Println("Part2: First round without moving elves:", i+1)
			break
		}
	}
	paintGrid()
	fmt.Println(gridSize, gridSize.MaxX-gridSize.MinX, gridSize.MaxY-gridSize.MinY, len(elfList))
	fmt.Println("Number of empty tiles:", (gridSize.MaxX-gridSize.MinX+1)*(gridSize.MaxY-gridSize.MinY+1)-len(elfList))

}

func paintGrid() {

	screen.Clear()
	output := ""
	for y := gridSize.MinY; y <= gridSize.MaxY; y++ { //
		for x := gridSize.MinX; x <= gridSize.MaxX; x++ {
			if _, ok := elfList[[2]int{x, y}]; ok {
				output = output + "#"

			} else {
				output = output + "."
			}
		}
		output = output + "\n"
	}
	fmt.Println(output)
	time.Sleep(time.Millisecond * 100)
}
