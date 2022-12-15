package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"regexp"
	"strings"
)

var gridSize *aocutils.Gridsize

func main() {

	gridSize = aocutils.NewGridSize()
	sensors := [][2]int{}
	beacons := [][2]int{}

	sensorList := aocutils.ReadInput("test-input.txt")
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

	grid := map[[2]int]rune{}
	for x := gridSize.MinX; x <= gridSize.MaxX; x++ {
		for y := gridSize.MinY; y <= gridSize.MaxY; y++ {
			grid[[2]int{x, y}] = '.'
		}
	}

	for _, sensorPoint := range sensors {
		grid[sensorPoint] = 'S'
	}
	for _, beaconPoint := range beacons {
		grid[beaconPoint] = 'B'
	}

	gridSize.PaintGrid(grid)

	for i := range sensors {
		calculateSensorBeaconBlock(sensors[i], beacons[i], &grid)
	}

}

func calculateSensorBeaconBlock(sensorPoint [2]int, beaconPoint [2]int, grid *map[[2]int]rune) {

	manhattanDistance := aocutils.Abs(sensorPoint[0]-beaconPoint[0]) + aocutils.Abs(sensorPoint[1]-beaconPoint[1])
	fmt.Println("Manhattan between", sensorPoint, "and", beaconPoint, ":", manhattanDistance)
}
