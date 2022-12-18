package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	room := map[[3]int]rune{}
	lavaBlockCoordsAsString := aocutils.ReadInput("input.txt")
	roomSize := aocutils.NewRoomSize()
	for _, lavaBlock := range lavaBlockCoordsAsString {
		coords := strings.Split(lavaBlock, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		room[[3]int{x, y, z}] = '#'
		roomSize.RecalibrateTo([3]int{x, y, z})
	}

	surfaceArea := 0
	for x := roomSize.MinX; x <= roomSize.MaxX; x++ {
		for y := roomSize.MinY; y <= roomSize.MaxY; y++ {
			for z := roomSize.MinZ; z <= roomSize.MaxZ; z++ {
				surfaceArea = surfaceArea + checkAreaOfPoint(room, [3]int{x, y, z})
			}
		}
	}

	fmt.Println("P1: Surface Area Of Lava Blob:", surfaceArea)

}

func checkAreaOfPoint(room map[[3]int]rune, point [3]int) int {
	x := point[0]
	y := point[1]
	z := point[2]

	if _, ok := room[point]; !ok {
		return 0
	}

	surroundingSurface := 0

	if _, ok := room[[3]int{x + 1, y, z}]; !ok {
		surroundingSurface++
	}
	if _, ok := room[[3]int{x - 1, y, z}]; !ok {
		surroundingSurface++
	}
	if _, ok := room[[3]int{x, y + 1, z}]; !ok {
		surroundingSurface++
	}
	if _, ok := room[[3]int{x, y - 1, z}]; !ok {
		surroundingSurface++
	}
	if _, ok := room[[3]int{x, y, z + 1}]; !ok {
		surroundingSurface++
	}
	if _, ok := room[[3]int{x, y, z - 1}]; !ok {
		surroundingSurface++
	}

	return surroundingSurface

}
