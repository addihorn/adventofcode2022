package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/inancgumus/screen"
)

var mapSize *aocutils.Gridsize

const inputFile = "input.txt"

func main() {

	mapSize = aocutils.NewGridSize()

	inputData := aocutils.ReadInputWithDelimeter(inputFile, "\n\n")
	mapData, startingPosition := parseInputMap(inputData[0])
	instructionSet := parseInstructions(inputData[1])
	fmt.Println(instructionSet)

	player := NewPlayer(startingPosition)
	PaintGrid(mapData, player)
	for _, instruction := range instructionSet {
		switch instruction {
		case "R", "L":
			player.Rotate(instruction)
		default:
			numberOfSteps, _ := strconv.Atoi(instruction)
			player.MoveSteps(numberOfSteps, &mapData)
		}

	}
	PaintGrid(mapData, player)
	row := player.Position[1]
	col := player.Position[0]
	facing := player.Facing

	fmt.Println("Part1:", row, col, facing, (1000*row)+(4*col)+facing)

}

func parseInputMap(mapAsString string) (map[[2]int]rune, [2]int) {

	mapData := map[[2]int]rune{}
	startingPosition := [2]int{}
	mapRows := strings.Split(mapAsString, "\n")
	mapSize.RecalibrateTo([2]int{1, 1})

	for row, mapRow := range mapRows {
		for col, mapSpace := range mapRow {
			mapSize.RecalibrateTo([2]int{col + 1, row + 1})
			switch mapSpace {
			case '.':
				mapData[[2]int{col + 1, row + 1}] = '.'
				if row == 0 && startingPosition == [2]int{} {
					startingPosition = [2]int{col + 1, row + 1}
				}
			case '#':
				mapData[[2]int{col + 1, row + 1}] = '#'
			}
		}
	}
	return mapData, startingPosition
}

func PaintGrid(mapData map[[2]int]rune, player *Player) {
	screen.Clear()
	output := ""
	playerFacing := ""
	switch player.Facing {
	case 0:
		playerFacing = ">"
	case 1:
		playerFacing = "v"
	case 2:
		playerFacing = "<"
	case 3:
		playerFacing = "^"
	}

	for y := mapSize.MinY; y <= mapSize.MaxY; y++ { //
		for x := mapSize.MinX; x <= mapSize.MaxX; x++ {
			if val, ok := mapData[[2]int{x, y}]; ok {
				if [2]int{x, y} == player.Position {
					output = output + string(playerFacing)
				} else {
					output = output + string(val)
				}

			} else {
				output = output + " "
			}

		}
		output = output + "\n"
	}
	fmt.Println(output)
	time.Sleep(time.Millisecond * 100)
}

func parseInstructions(instructionsAsString string) []string {
	instructionSet := []string{}
	steps := ""
	for _, char := range instructionsAsString {
		switch char {
		case 'R', 'L':
			instructionSet = append(instructionSet, []string{steps, string(char)}...)
			steps = ""
		default:
			steps = steps + string(char)
		}
	}
	return instructionSet
}
