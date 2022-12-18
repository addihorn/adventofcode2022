package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
)

var gridSize *aocutils.Gridsize
var grid map[[2]int]rune

//var targetRowBlock map[[2]int]rune

func main() {

	fmt.Println(math.MaxInt)
	gridSize = aocutils.NewGridSize()
	//targetRowBlock = map[[2]int]rune{}
	sensors := [][2]int{}
	beacons := [][2]int{}
	rowToCheckP1 := 2000000
	//rowToCheckP1 := 10
	maxX := 4000000 //20
	//maxX := //20
	maxY := 4000000 //20
	//maxY := //20

	sensorList := aocutils.ReadInput("input.txt")
	for _, sensor := range sensorList {
		sensorData := strings.Split(sensor, ":")[0]
		beaconData := strings.Split(sensor, ":")[1]

		findX := regexp.MustCompile("x=-?\\d+")
		findY := regexp.MustCompile("y=-?\\d+")

		sensorPoint := [2]int{
			aocutils.CString2Int(strings.Split(findX.FindString(sensorData), "=")[1]),
			aocutils.CString2Int(strings.Split(findY.FindString(sensorData), "=")[1])}

		beaconPoint := [2]int{
			aocutils.CString2Int(strings.Split(findX.FindString(beaconData), "=")[1]),
			aocutils.CString2Int(strings.Split(findY.FindString(beaconData), "=")[1])}

		gridSize.RecalibrateTo(sensorPoint)
		gridSize.RecalibrateTo(beaconPoint)

		sensors = append(sensors, sensorPoint)
		beacons = append(beacons, beaconPoint)
	}
	fmt.Println("Sensor-Liste:", sensors, "\nBeacon-Liste:", beacons)
	fmt.Println(gridSize)

	grid = map[[2]int]rune{}

	for _, sensorPoint := range sensors {
		grid[sensorPoint] = 'S'
	}
	for _, beaconPoint := range beacons {
		grid[beaconPoint] = 'B'
	}

	//gridSize.PaintGrid(grid)

	for i := range sensors {
		//calculateSensorBeaconBlock(sensors[i], beacons[i])
		calculateSensorBeaconBlockOnRowOnly(sensors[i], beacons[i], rowToCheckP1)
		//gridSize.PaintGrid(grid)
	}

	beaconImpossible := 0
	for i := gridSize.MinX; i <= gridSize.MaxX; i++ {
		if grid[[2]int{i, rowToCheckP1}] == '#' {
			beaconImpossible++
		}
	}

	fmt.Println("Part1: Impossible Beacon-Positions in row", rowToCheckP1, ":", beaconImpossible)

	gridSize.MinX = 0
	gridSize.MaxX = maxX
	gridSize.MinY = 0
	gridSize.MaxY = maxY

	//gridSize.PaintGrid(grid)

	//rowWithBeacon := 0
	for rowTocheck := gridSize.MinY; rowTocheck <= gridSize.MaxY; rowTocheck++ {

		startPoints := []int{}
		intervals := map[int]int{}

		for i := range sensors {
			//calculateSensorBeaconBlock(sensors[i], beacons[i])
			distanceAffectsRow, start, end := calculateSensorBeaconWithStartAndEnd(sensors[i], beacons[i], rowTocheck)
			if distanceAffectsRow {
				startPoints = append(startPoints, start)

				intervals[start] = aocutils.Max([]int{end, intervals[start]})
			}
			//gridSize.PaintGrid(grid)
		}
		sort.Ints(startPoints)
		//fmt.Println(startPoints)

		maxEnd := -math.MaxInt
		for i, point := range startPoints[0 : len(startPoints)-1] {
			maxEnd = aocutils.Max([]int{intervals[point], maxEnd})
			if startPoints[i+1]-maxEnd > 1 {
				//gap found
				fmt.Println(intervals)
				fmt.Println("I found a gap in row:", rowTocheck)
				fmt.Println("End of first slice:", maxEnd, "Start of new slice:", startPoints[i+1])
				fmt.Println("P2: Signal strength:", int64((maxEnd+1)*4000000+rowTocheck))
			}
		}
	}

}

func calculateSensorBeaconWithStartAndEnd(sensorPoint [2]int, beaconPoint [2]int, targetRow int) (bool, int, int) {
	manhattanDistance := aocutils.Abs(sensorPoint[0]-beaconPoint[0]) + aocutils.Abs(sensorPoint[1]-beaconPoint[1])
	//fmt.Println("Manhattan between", sensorPoint, "and", beaconPoint, ":", manhattanDistance)

	//gridSize.RecalibrateTo([2]int{sensorPoint[0] + manhattanDistance, sensorPoint[1] + manhattanDistance})
	//gridSize.RecalibrateTo([2]int{sensorPoint[0] - manhattanDistance, sensorPoint[1] - manhattanDistance})

	if (targetRow < sensorPoint[1]-manhattanDistance) || (targetRow > sensorPoint[1]+manhattanDistance) {
		return false, 0, 0
	}

	targetRowOffset := aocutils.Abs(sensorPoint[1] - targetRow)
	return true, sensorPoint[0] - (manhattanDistance - targetRowOffset), sensorPoint[0] + (manhattanDistance - targetRowOffset)

}

func calculateSensorBeaconBlockOnRowOnly(sensorPoint [2]int, beaconPoint [2]int, targetRow int) {
	manhattanDistance := aocutils.Abs(sensorPoint[0]-beaconPoint[0]) + aocutils.Abs(sensorPoint[1]-beaconPoint[1])
	//fmt.Println("Manhattan between", sensorPoint, "and", beaconPoint, ":", manhattanDistance)

	gridSize.RecalibrateTo([2]int{sensorPoint[0] + manhattanDistance, sensorPoint[1] + manhattanDistance})
	gridSize.RecalibrateTo([2]int{sensorPoint[0] - manhattanDistance, sensorPoint[1] - manhattanDistance})

	if (targetRow < sensorPoint[1]-manhattanDistance) && (targetRow > sensorPoint[1]+manhattanDistance) {
		return
	}

	targetRowOffset := aocutils.Abs(sensorPoint[1] - targetRow)
	targetRowLeftX := sensorPoint[0] - (manhattanDistance - targetRowOffset)
	targetRowRightX := sensorPoint[0] + (manhattanDistance - targetRowOffset)

	for x := targetRowLeftX; x <= targetRowRightX; x++ {
		setInpossibleSpot([2]int{x, targetRow})
	}

}

func calculateSensorBeaconBlock(sensorPoint [2]int, beaconPoint [2]int) {

	manhattanDistance := aocutils.Abs(sensorPoint[0]-beaconPoint[0]) + aocutils.Abs(sensorPoint[1]-beaconPoint[1])
	//fmt.Println("Manhattan between", sensorPoint, "and", beaconPoint, ":", manhattanDistance)
	//fmt.Println(gridSize)

	//gridSize.RecalibrateTo([2]int{sensorPoint[0] + manhattanDistance, sensorPoint[1] + manhattanDistance})
	//gridSize.RecalibrateTo([2]int{sensorPoint[0] - manhattanDistance, sensorPoint[1] - manhattanDistance})

	for i := 0; i <= manhattanDistance; i++ {
		for l := 0; l <= (manhattanDistance - aocutils.Abs(i)); l++ {
			checkX := sensorPoint[0] + i
			checkY := sensorPoint[1] + l
			setInpossibleSpot([2]int{checkX, checkY})

			checkX = sensorPoint[0] - i
			checkY = sensorPoint[1] + l
			setInpossibleSpot([2]int{checkX, checkY})

			checkX = sensorPoint[0] + i
			checkY = sensorPoint[1] - l
			setInpossibleSpot([2]int{checkX, checkY})

			checkX = sensorPoint[0] - i
			checkY = sensorPoint[1] - l
			setInpossibleSpot([2]int{checkX, checkY})
		}
		fmt.Println("Marked impossible Slots in row", i)
	}

}

func setInpossibleSpot(point [2]int) {
	if (gridSize.MinX > point[0]) || (gridSize.MaxX < point[0]) {
		return
	}
	if _, ok := grid[point]; !ok {
		grid[point] = '#'
	}
	//gridSize.PaintGrid(grid)
}
