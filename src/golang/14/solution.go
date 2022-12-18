package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type (
	Gridsize struct {
		minX int
		minY int
		maxX int
		maxY int
	}
	Rockstructure struct {
		wall [][2]int
	}
)

var Cavestructure [][]rune
var gridSize *Gridsize

func (this *Gridsize) recalibrateTo(point [2]int) {

	this.minX = min([]int{this.minX, point[0]})
	this.minY = min([]int{this.minY, point[1]})
	this.maxX = max([]int{this.maxX, point[0]})
	this.maxY = max([]int{this.maxY, point[1]})
}
func min(intValues []int) int {
	sort.Ints(intValues)
	return intValues[0]
}

func max(intValues []int) int {
	sort.Ints(intValues)
	return intValues[len(intValues)-1]
}

func NewRockstructure(startingPoint [2]int) *Rockstructure {

	return &Rockstructure{[][2]int{startingPoint}}
}

func (this *Rockstructure) drawPathTo(point [2]int) {
	currentEndOfWall := this.wall[len(this.wall)-1]

	currentX := currentEndOfWall[0]
	currentY := currentEndOfWall[1]

	switch {
	case currentX > point[0]:
		for x := currentX - 1; x >= point[0]; x-- {
			this.wall = append(this.wall, [2]int{x, currentY})
		}
	case currentX < point[0]:
		for x := currentX + 1; x <= point[0]; x++ {
			this.wall = append(this.wall, [2]int{x, currentY})
		}
	case currentY > point[1]:
		for y := currentY - 1; y >= point[1]; y-- {
			this.wall = append(this.wall, [2]int{currentX, y})
		}
	case currentY < point[1]:
		for y := currentY + 1; y <= point[1]; y++ {
			this.wall = append(this.wall, [2]int{currentX, y})
		}
	}

}

func calculatePointAsInt(pointAsString []string) [2]int {
	var point [2]int = [2]int{0, 0}

	point[0], _ = strconv.Atoi(pointAsString[0])
	point[1], _ = strconv.Atoi(pointAsString[1])

	return point

}

func growCave(indexToAttach int) {
	newColumn := make([][]rune, 1)

	newColumn[0] = make([]rune, len(Cavestructure[0]))
	for y := range newColumn[0] {
		newColumn[0][y] = ' '
	}
	newColumn[0][len(newColumn[0])-1] = '#'

	if indexToAttach == 0 {
		Cavestructure = append(newColumn, Cavestructure...)
		gridSize.minX--
	} else {
		Cavestructure = append(Cavestructure, newColumn...)
		gridSize.maxX--
	}

}

func main() {

	wallsOfRock := aocutils.ReadInput("input.txt")
	allRockStructures := []*Rockstructure{}
	gridSize = &Gridsize{500, 0, 500, 0}
	for _, wallOfRock := range wallsOfRock {
		wallEdges := strings.Split(wallOfRock, "->")
		startingPoint := calculatePointAsInt(strings.Split(strings.TrimSpace(wallEdges[0]), ","))
		gridSize.recalibrateTo(startingPoint)
		currentStructure := NewRockstructure(startingPoint)

		for _, wallEdge := range wallEdges[1:len(wallEdges)] {
			newPoint := calculatePointAsInt(strings.Split(strings.TrimSpace(wallEdge), ","))
			gridSize.recalibrateTo(newPoint)
			currentStructure.drawPathTo(newPoint)
		}

		allRockStructures = append(allRockStructures, currentStructure)
		//fmt.Println(currentStructure)
	}

	//fmt.Println(gridSize)

	// init Grid
	// part1
	Cavestructure = make([][]rune, gridSize.maxX-gridSize.minX+2)

	for x := range Cavestructure {
		Cavestructure[x] = make([]rune, gridSize.maxY-gridSize.minY+1)
		for y := range Cavestructure[x] {
			Cavestructure[x][y] = ' '
		}
	}

	for _, rockStructure := range allRockStructures {
		for _, point := range rockStructure.wall {
			x := point[0] - gridSize.minX + 1
			y := point[1] - gridSize.minY
			Cavestructure[x][y] = '#'
		}
	}

	sandSpawnPoint := [2]int{500, 0}

	Cavestructure[sandSpawnPoint[0]-gridSize.minX+1][sandSpawnPoint[1]] = 'x'

	sandCorns := 0
	//PaintCave(Cavestructure)
	for {
		_, noWayToGo := checkSandRestingPosition([2]int{sandSpawnPoint[0] - gridSize.minX + 1, 0})
		if noWayToGo {
			break
		}
		//PaintCave(Cavestructure)
		sandCorns++
	}
	PaintCave(Cavestructure)
	fmt.Println("Part1: I created", sandCorns, "units of sand")
	//part2
	Cavestructure = make([][]rune, gridSize.maxX-gridSize.minX+2)
	for x := range Cavestructure {
		Cavestructure[x] = make([]rune, gridSize.maxY-gridSize.minY+3)
		for y := range Cavestructure[x] {
			Cavestructure[x][y] = ' '
		}
		//build floor
		Cavestructure[x][len(Cavestructure[x])-1] = '#'
	}

	for _, rockStructure := range allRockStructures {
		for _, point := range rockStructure.wall {
			x := point[0] - gridSize.minX + 1
			y := point[1] - gridSize.minY
			Cavestructure[x][y] = '#'
		}
	}

	Cavestructure[sandSpawnPoint[0]-gridSize.minX+1][sandSpawnPoint[1]] = 'x'

	sandCorns = 0
	//PaintCave(Cavestructure)
	for {
		_, noWayToGo := checkSandRestingPosition2([2]int{sandSpawnPoint[0] - gridSize.minX + 1, 0})
		if noWayToGo {
			break
		}
		//PaintCave(Cavestructure)
		sandCorns++
	}
	//PaintCaveToFile(Cavestructure)
	fmt.Println("Part2: I created", sandCorns+1, "units of sand")
}

func checkSandRestingPosition(sandUnitPosition [2]int) ([2]int, bool) {
	//sand falls down from point until bottom is all full
	x := sandUnitPosition[0]
	y := sandUnitPosition[1]

	//check for indexOutOfBound
	if x+2 > len(Cavestructure) || x-1 < 0 || y+2 > len(Cavestructure[0]) {
		return sandUnitPosition, true
	}

	switch {
	case Cavestructure[x][y+1] == ' ': //fall down
		return checkSandRestingPosition([2]int{x, y + 1})
	case Cavestructure[x-1][y+1] == ' ': //fall to the left
		return checkSandRestingPosition([2]int{x - 1, y + 1})
	case Cavestructure[x+1][y+1] == ' ': //fall to the right
		return checkSandRestingPosition([2]int{x + 1, y + 1})
	default: //no way to fall
		Cavestructure[x][y] = 'o'
		return sandUnitPosition, false
	}
}

func checkSandRestingPosition2(sandUnitPosition [2]int) ([2]int, bool) {
	//sand falls down from point until bottom is all full
	x := sandUnitPosition[0]
	y := sandUnitPosition[1]

	//check for indexOutOfBound
	switch {
	case x+2 > len(Cavestructure):
		growCave(x)
	case x-1 < 0:
		x++
		growCave(0)
	}

	switch {
	case Cavestructure[x][y+1] == ' ': //fall down
		return checkSandRestingPosition2([2]int{x, y + 1})
	case Cavestructure[x-1][y+1] == ' ': //fall to the left
		return checkSandRestingPosition2([2]int{x - 1, y + 1})
	case Cavestructure[x+1][y+1] == ' ': //fall to the right
		return checkSandRestingPosition2([2]int{x + 1, y + 1})
	default: //no way to fall
		if y == 0 {
			return sandUnitPosition, true
		}
		Cavestructure[x][y] = 'o'
		return sandUnitPosition, false
	}
}

func PaintCave(grid [][]rune) {

	output := ""
	/*
		for _, pseudoRow := range grid {
			for _, value := range pseudoRow {
				output = output + string(value)
			}
			output = output + "\n"
		}
	*/
	//screen.Clear()
	for row := 0; row < len(grid[0]); row++ {

		for col := 0; col < len(grid); col++ {
			output = output + string(grid[col][row])

		}
		output = output + "\n"
	}

	fmt.Println(output)
	time.Sleep(time.Millisecond * 100)
}

func PaintCaveToFile(grid [][]rune) {

	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	output := ""
	/*
		for _, pseudoRow := range grid {
			for _, value := range pseudoRow {
				output = output + string(value)
			}
			output = output + "\n"
		}
	*/
	for row := 0; row < len(grid[0]); row++ {

		for col := 0; col < len(grid); col++ {
			output = output + string(grid[col][row])

		}
		output = output + "\n"
	}

	_, err2 := f.WriteString(output)

	if err2 != nil {
		log.Fatal(err2)
	}

}
