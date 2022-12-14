package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/inancgumus/screen"
)

type HightPoint struct {
	value         rune
	position      [2]int
	upPossible    bool
	rightPossible bool
	downPossible  bool
	leftPossible  bool
	stepsToTarget int
}

func NewPoint(col int, row int, hightProfile []string) *HightPoint {

	this := &HightPoint{0, [2]int{col, row}, false, false, false, false, math.MaxInt}

	this.value = getSanitizedValue(rune(hightProfile[row][col]))

	if row > 0 {
		this.upPossible = this.value-getSanitizedValue(rune(hightProfile[row-1][col])) > -2
	}
	if row < len(hightProfile)-1 {
		this.downPossible = this.value-getSanitizedValue(rune(hightProfile[row+1][col])) > -2
	}

	if col > 0 {
		this.leftPossible = this.value-getSanitizedValue(rune(hightProfile[row][col-1])) > -2
	}
	if col < len(hightProfile[0])-1 {
		this.rightPossible = this.value-getSanitizedValue(rune(hightProfile[row][col+1])) > -2
	}

	return this
}

func (this *HightPoint) MarkAsDead() {
	this.downPossible = false
	this.leftPossible = false
	this.rightPossible = false
	this.upPossible = false
}

func (this *HightPoint) isDead() bool {
	return !(this.downPossible || this.leftPossible || this.rightPossible || this.upPossible || this.stepsToTarget < math.MaxInt)
}

func getSanitizedValue(originalValue rune) rune {
	switch originalValue {
	case 'S':
		return 'a'
	case 'E':
		return 'z'
	default:
		return originalValue
	}
}

func positionAlreadyVisited(trail [][2]int, position [2]int) bool {

	for _, step := range trail {
		if position == step {
			return true
		}
	}
	return false

}

func main() {
	heightProfile := aocutils.ReadInput("input.txt")

	profileMap := make([][]*HightPoint, len(heightProfile[0]))

	for i := range profileMap {
		profileMap[i] = make([]*HightPoint, len(heightProfile))
	}

	for rowNo, hightPoints := range heightProfile {
		for colNo, _ := range hightPoints {
			profileMap[colNo][rowNo] = NewPoint(colNo, rowNo, heightProfile)
		}
	}

	startPoint := profileMap[0][0]
	endPoint := profileMap[0][0]

	for rowNo, hightPoints := range heightProfile {
		for colNo, hightValue := range hightPoints {
			switch hightValue {
			case 'S':
				startPoint = profileMap[colNo][rowNo]
			case 'E':
				endPoint = profileMap[colNo][rowNo]
			default:
			}
		}
	}

	fmt.Println(*startPoint)
	fmt.Println(*endPoint)

	fmt.Println(moveToTarget(startPoint, endPoint, profileMap, [][2]int{}))

}

func moveToTarget(currentPosition *HightPoint, targetPosition *HightPoint, grid [][]*HightPoint, trail [][2]int) (int, bool) {

	moveUpBias, moveRightBias, moveDownBias, moveLeftBias := calculateStepBias(currentPosition, targetPosition, grid)
	//fmt.Println(currentPosition, moveUpBias, moveRightBias, moveDownBias, moveLeftBias)

	currentCol := currentPosition.position[0]
	currentRow := currentPosition.position[1]
	trail = append(trail, currentPosition.position)
	//PaintPath(grid, trail)
	if currentPosition.position == targetPosition.position {
		currentPosition.stepsToTarget = 0
		return 0, true
	}
	if positionAlreadyVisited(trail[0:len(trail)-1], currentPosition.position) {
		return 9999, false
	}

	if currentPosition.stepsToTarget < math.MaxInt {
		return currentPosition.stepsToTarget, true
	}

	endFound := false
	numberOfStepsArray := []int{math.MaxInt - 1}
	//fmt.Println(numberOfStepsArray)
	for moveUpBias != 0 || moveDownBias != 0 || moveLeftBias != 0 || moveRightBias != 0 {
		stepsTaken := 0
		endFoundNew := false

		switch {
		case moveUpBias >= moveDownBias &&
			moveUpBias >= moveLeftBias &&
			moveUpBias >= moveRightBias:
			newStartPosition := grid[currentCol][currentRow-1]
			//newStartPosition.downPossible = false
			stepsTaken, endFoundNew = moveToTarget(newStartPosition, targetPosition, grid, trail)
			currentPosition.upPossible = false
			//newStartPosition.downPossible = true
		case moveRightBias >= moveDownBias &&
			moveRightBias >= moveLeftBias &&
			moveRightBias >= moveUpBias:
			newStartPosition := grid[currentCol+1][currentRow]
			//newStartPosition.leftPossible = false
			stepsTaken, endFoundNew = moveToTarget(newStartPosition, targetPosition, grid, trail)
			currentPosition.rightPossible = false
		case moveDownBias >= moveLeftBias &&
			moveDownBias >= moveUpBias &&
			moveDownBias >= moveRightBias:
			newStartPosition := grid[currentCol][currentRow+1]
			//newStartPosition.upPossible = false

			stepsTaken, endFoundNew = moveToTarget(newStartPosition, targetPosition, grid, trail)
			currentPosition.downPossible = false
			//newStartPosition.upPossible = true
		case moveLeftBias >= moveUpBias &&
			moveLeftBias >= moveRightBias &&
			moveLeftBias >= moveDownBias:
			newStartPosition := grid[currentCol-1][currentRow]
			//newStartPosition.rightPossible = false

			stepsTaken, endFoundNew = moveToTarget(newStartPosition, targetPosition, grid, trail)
			currentPosition.leftPossible = false
			//newStartPosition.rightPossible = true
		default:

		}

		moveUpBias, moveRightBias, moveDownBias, moveLeftBias = calculateStepBias(currentPosition, targetPosition, grid)
		endFound = endFound || endFoundNew
		if endFound {
			numberOfStepsArray = append(numberOfStepsArray, stepsTaken+1)

			sort.Ints(numberOfStepsArray)
			currentPosition.stepsToTarget = numberOfStepsArray[0]
		}
		//fmt.Println(numberOfStepsArray)
	}

	sort.Ints(numberOfStepsArray)
	if endFound {
		//PaintPathBackwards(grid)
		fmt.Println(currentPosition, "returning steps taken", numberOfStepsArray)
	}
	return numberOfStepsArray[0], endFound

}

func PaintPathBackwards(grid [][]*HightPoint) {
	screen.Clear()
	output := ""
	for row := 0; row < len(grid[0]); row++ {

		for col := 0; col < len(grid); col++ {
			switch {
			case grid[col][row].stepsToTarget < math.MaxInt:
				output = output + "#"
			case grid[col][row].isDead():
				output = output + "."
			default:
				output = output + " "
			}
		}
		output = output + "\n"
	}
	fmt.Println(output)
	time.Sleep(time.Millisecond * 10)
}

func PaintPath(grid [][]*HightPoint, trail [][2]int) {

	screen.Clear()
	output := ""
	for row := 0; row < len(grid[0]); row++ {

		for col := 0; col < len(grid); col++ {
			switch {
			case positionAlreadyVisited(trail, grid[col][row].position):
				output = output + "#"
			case grid[col][row].isDead():
				output = output + "."
			default:
				output = output + " "
			}
		}
		output = output + "\n"
	}
	fmt.Println(output)
	time.Sleep(time.Millisecond * 10)
}

func calculateStepBias(currentPosition *HightPoint, targetPosition *HightPoint, grid [][]*HightPoint) (int, int, int, int) {

	moveUpBias := currentPosition.position[1] - targetPosition.position[1]
	moveDownBias := moveUpBias * -1
	moveLeftBias := currentPosition.position[0] - targetPosition.position[0]
	moveRightBias := moveLeftBias * -1

	biasSanitizationValue := min([]int{moveUpBias, moveDownBias, moveLeftBias, moveRightBias})*-1 + 1

	moveUpBias = (moveUpBias + biasSanitizationValue) * cBool2Int(currentPosition.upPossible)
	moveDownBias = (moveDownBias + biasSanitizationValue) * cBool2Int(currentPosition.downPossible)
	moveLeftBias = (moveLeftBias + biasSanitizationValue) * cBool2Int(currentPosition.leftPossible)
	moveRightBias = (moveRightBias + biasSanitizationValue) * cBool2Int(currentPosition.rightPossible)

	return moveUpBias, moveRightBias, moveDownBias, moveLeftBias

}

func cBool2Int(value bool) int {
	if value {
		return 1
	}
	return 0
}

func min(intValues []int) int {
	sort.Ints(intValues)
	return intValues[0]
}
