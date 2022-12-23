package main

type Elf struct {
	Position     [2]int
	proposedTile [2]int
}

var proposalList map[[2]int]int
var elfList map[[2]int]*Elf

func NewElf(startingPosition [2]int) *Elf {
	gridSize.RecalibrateTo(startingPosition)
	return &Elf{startingPosition, [2]int{}}
}

func startNewRound() {
	proposalList = map[[2]int]int{}
	for _, elf := range elfList {
		elf.proposedTile = [2]int{}
	}
	elfMoved = false
}

func (this *Elf) ProposeTile() {

	//first check all surrounding tiles
	x := this.Position[0]
	y := this.Position[1]

	gOk := false
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if !(x == i && y == j) {
				_, ok := elfList[[2]int{i, j}]
				gOk = gOk || ok
			}

		}
	}
	if !gOk {
		this.proposedTile = this.Position
		proposalList[this.proposedTile]++
		return
	}

	//consider other moves
	for _, checkDirection := range orderOfProposal {
		gOk = false
		switch checkDirection {
		case 'N':
			for i := x - 1; i < x+2; i++ {
				_, ok := elfList[[2]int{i, y - 1}]
				gOk = gOk || ok
			}
			if !gOk {
				newPorposal := [2]int{x, y - 1}
				this.proposedTile = newPorposal
				proposalList[newPorposal]++
				return
			}
		case 'S':
			for i := x - 1; i < x+2; i++ {
				_, ok := elfList[[2]int{i, y + 1}]
				gOk = gOk || ok
			}
			if !gOk {
				newPorposal := [2]int{x, y + 1}
				this.proposedTile = newPorposal
				proposalList[newPorposal]++
				return
			}
		case 'W':
			for j := y - 1; j < y+2; j++ {
				_, ok := elfList[[2]int{x - 1, j}]
				gOk = gOk || ok
			}
			if !gOk {
				newPorposal := [2]int{x - 1, y}
				this.proposedTile = newPorposal
				proposalList[newPorposal]++
				return
			}
		case 'E':
			for j := y - 1; j < y+2; j++ {
				_, ok := elfList[[2]int{x + 1, j}]
				gOk = gOk || ok
			}
			if !gOk {
				newPorposal := [2]int{x + 1, y}
				this.proposedTile = newPorposal
				proposalList[newPorposal]++
				return
			}
		}
	}

	if this.proposedTile == [2]int{} {
		this.proposedTile = this.Position
		proposalList[this.proposedTile]++
		return
	}

}

func (this *Elf) MoveToProposedTile() {
	proposedTile := this.proposedTile
	if this.Position == this.proposedTile {
		return
	}

	if proposalList[proposedTile] == 1 {
		delete(elfList, this.Position)
		this.Position = proposedTile
		elfList[this.Position] = this
		gridSize.RecalibrateTo(this.Position)
		elfMoved = true
	}
}
