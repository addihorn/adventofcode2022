package aocutils

import (
	"fmt"
	"math"

	"github.com/inancgumus/screen"
)

type Gridsize struct {
	MinX int
	MinY int
	MaxX int
	MaxY int
}

func NewGridSize() *Gridsize {
	return &Gridsize{math.MaxInt, math.MaxInt, math.MinInt, math.MinInt}
}

func (this *Gridsize) RecalibrateTo(point [2]int) {

	this.MinX = min([]int{this.MinX, point[0]})
	this.MinY = min([]int{this.MinY, point[1]})
	this.MaxX = max([]int{this.MaxX, point[0]})
	this.MaxY = max([]int{this.MaxY, point[1]})
}

func (this *Gridsize) PaintGrid(grid map[[2]int]rune) {
	screen.Clear()
	for y := this.MinY; y <= this.MaxY; y++ {
		for x := this.MinX; x <= this.MaxX; x++ {
			if val, ok := grid[[2]int{x, y}]; ok {
				fmt.Print(string(val))
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}
	//time.Sleep(time.Millisecond * 100)

}
