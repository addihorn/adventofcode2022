package main

type Player struct {
	Position [2]int
	Facing   int //Facing is 0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^)
}

func NewPlayer(startingPosition [2]int) *Player {
	return &Player{startingPosition, 0}
}

func (this *Player) Rotate(rotation string) {
	switch {
	case rotation == "R":
		this.Facing++
	case rotation == "L":
		this.Facing--
	}

	switch {
	case this.Facing == -1:
		this.Facing = 3
	case this.Facing == 4:
		this.Facing = 0
	}
}

func (this *Player) MoveSteps(numberOfSteps int, grid *map[[2]int]rune) {

	for i := 0; i < numberOfSteps; i++ {
		initialPos := this.Position
		switch this.Facing {
		case 0:
			this.Position = this.moveTo([2]int{initialPos[0] + 1, initialPos[1]}, *grid)
		case 1:
			this.Position = this.moveTo([2]int{initialPos[0], initialPos[1] + 1}, *grid)
		case 2:
			this.Position = this.moveTo([2]int{initialPos[0] - 1, initialPos[1]}, *grid)
		case 3:
			this.Position = this.moveTo([2]int{initialPos[0], initialPos[1] - 1}, *grid)
		}

		if initialPos == this.Position {
			//wall reached, no more movement possible
			break
		}
		//PaintGrid(*grid, this)
	}

}

func (this *Player) moveTo(targetPos [2]int, grid map[[2]int]rune) [2]int {

	switch grid[targetPos] {
	case '.':
		return targetPos
	case '#':
		return this.Position
	}

	//check wrap around
	switch this.Facing {
	case 0:
		ok := true
		x := this.Position[0]
		y := this.Position[1]

		for ok {
			_, ok = grid[[2]int{x - 1, y}]
			x--
		}
		return this.moveTo([2]int{x + 1, y}, grid)
	case 1:
		ok := true
		x := this.Position[0]
		y := this.Position[1]

		for ok {
			_, ok = grid[[2]int{x, y - 1}]
			y--
		}
		return this.moveTo([2]int{x, y + 1}, grid)
	case 2:
		ok := true
		x := this.Position[0]
		y := this.Position[1]

		for ok {
			_, ok = grid[[2]int{x + 1, y}]
			x++
		}
		return this.moveTo([2]int{x - 1, y}, grid)
	case 3:
		ok := true
		x := this.Position[0]
		y := this.Position[1]

		for ok {
			_, ok = grid[[2]int{x, y + 1}]
			y++
		}
		return this.moveTo([2]int{x, y - 1}, grid)

	}
	return [2]int{0, 0}

}
