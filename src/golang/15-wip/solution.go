package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"regexp"
	"strings"
)

var gridSize *aocutils.Gridsize
var grid map[[2]int]rune

func main() {

	gridSize = aocutils.NewGridSize()
	sensors := [][2]int{}
	beacons := [][2]int{}
	rowToCheck := 2000000

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
		calculateSensorBeaconBlock(sensors[i], beacons[i])
		gridSize.PaintGrid(grid)
	}

	noBeaconPossibleInRow := 0
	for i := gridSize.MinX; i <= gridSize.MaxX; i++ {
		if grid[[2]int{i, rowToCheck}] == '#' {
			noBeaconPossibleInRow++
		}
	}
	fmt.Println(noBeaconPossibleInRow, "places in row", rowToCheck, "where beacons are impossible")

}

func calculateSensorBeaconBlock(sensorPoint [2]int, beaconPoint [2]int) {

	manhattanDistance := aocutils.Abs(sensorPoint[0]-beaconPoint[0]) + aocutils.Abs(sensorPoint[1]-beaconPoint[1])
	fmt.Println("Manhattan between", sensorPoint, "and", beaconPoint, ":", manhattanDistance)
	fmt.Println(gridSize)

	for i := 0; i <= manhattanDistance; i++ {
		for l := 0; l <= (manhattanDistance - aocutils.Abs(i)); l++ {
			checkX := sensorPoint[0] + i
			checkY := sensorPoint[1] + l
			if _, ok := grid[[2]int{checkX, checkY}]; !ok {
				grid[[2]int{checkX, checkY}] = '#'
				gridSize.RecalibrateTo([2]int{checkX, checkY})
			}

			checkX = sensorPoint[0] - i
			checkY = sensorPoint[1] + l
			if _, ok := grid[[2]int{checkX, checkY}]; !ok {
				grid[[2]int{checkX, checkY}] = '#'
				gridSize.RecalibrateTo([2]int{checkX, checkY})
			}

			checkX = sensorPoint[0] + i
			checkY = sensorPoint[1] - l
			if _, ok := grid[[2]int{checkX, checkY}]; !ok {
				grid[[2]int{checkX, checkY}] = '#'
				gridSize.RecalibrateTo([2]int{checkX, checkY})
			}

			checkX = sensorPoint[0] - i
			checkY = sensorPoint[1] - l
			if _, ok := grid[[2]int{checkX, checkY}]; !ok {
				grid[[2]int{checkX, checkY}] = '#'
				gridSize.RecalibrateTo([2]int{checkX, checkY})
			}
		}
		fmt.Println("Marked impossible Slots in row", i)
	}

}
