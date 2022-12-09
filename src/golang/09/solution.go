package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Knot struct {
	PosX  int
	PosY  int
	Trail []string
}

func NewPlankKnot() *Knot {
	return &Knot{0, 0, []string{"0,0"}}
}

func (this *Knot) MoveHorizontal(step int) {
	this.PosX += step
	s := []string{strconv.Itoa(this.PosX), strconv.Itoa(this.PosY)}
	this.Trail = append(this.Trail, strings.Join(s, ","))
}

func (this *Knot) MoveVertical(step int) {
	this.PosY += step
	s := []string{strconv.Itoa(this.PosX), strconv.Itoa(this.PosY)}
	this.Trail = append(this.Trail, strings.Join(s, ","))
}

// Diagonal Steps

func (this *Knot) MoveDiagonal(horizontal int, vertical int) {
	this.PosX += horizontal
	this.PosY += vertical
	s := []string{strconv.Itoa(this.PosX), strconv.Itoa(this.PosY)}
	this.Trail = append(this.Trail, strings.Join(s, ","))
}

func (this *Knot) distanceTo(target *Knot) (int, int) {
	distanceX := target.PosX - this.PosX
	distanceY := target.PosY - this.PosY

	return distanceX, distanceY
}

func (this *Knot) isAdjecent(target *Knot) bool {
	distanceX, distanceY := this.distanceTo(target)

	return (math.Abs(float64(distanceX)) <= 1 &&
		math.Abs(float64(distanceY)) <= 1)

}

func (this *Knot) followTarget(target *Knot) {

	if this.isAdjecent(target) {
		return
	}

	distanceX, distanceY := this.distanceTo(target)
	switch {
	case distanceY == 0:
		step := distanceX / int(math.Abs(float64(distanceX)))
		this.MoveHorizontal(step)
	case distanceX == 0:
		step := distanceY / int(math.Abs(float64(distanceY)))
		this.MoveVertical(step)
	default:
		stepX := distanceX / int(math.Abs(float64(distanceX)))
		stepY := distanceY / int(math.Abs(float64(distanceY)))
		this.MoveDiagonal(stepX, stepY)
	}

}

const ropeLength = 10

func main() {

	var rope [ropeLength]*Knot
	for i := 0; i < len(rope); i++ {
		rope[i] = NewPlankKnot()
	}

	head := rope[0]
	tail := rope[len(rope)-1]

	headMovements := aocutils.ReadInput("input.txt")
	uniqueVisitedSpots := make(map[string]int)

	for _, headMove := range headMovements {
		moveDetails := strings.Split(headMove, " ")
		direction := moveDetails[0]
		steps, _ := strconv.Atoi(moveDetails[1])
		for steps > 0 {
			switch direction {
			case "R":
				head.MoveHorizontal(1)
			case "L":
				head.MoveHorizontal(-1)
			case "U":
				head.MoveVertical(1)
			case "D":
				head.MoveVertical(-1)
			default:
			}
			//the rope follows head
			//each knot follows its leading knot
			for i, knot := range rope[1:len(rope)] {
				knot.followTarget(rope[i])
			}
			steps--
			uniqueVisitedSpots[tail.Trail[len(tail.Trail)-1]]++
		}
	}
	//	fmt.Println(head)
	//fmt.Println(tail)
	fmt.Println("Tail Positions:", len(uniqueVisitedSpots))
}
