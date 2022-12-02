package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	b, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	gamePlan := string(b)

	gamePlanAsArray := strings.Split(gamePlan, "\n")
	fmt.Println(len(gamePlanAsArray))

	ownScore := 0
	opponentScore := 0

	for _, play := range gamePlanAsArray {
		opponentPlay := string(play[0])
		ownPlay := string(play[2])

		opponentGameScore, ownGameScore := calculateScores(opponentPlay, ownPlay)
		opponentScore += opponentGameScore
		ownScore += ownGameScore
		fmt.Println("Opponent: ", opponentPlay, opponentGameScore, opponentScore, " - own play: ", ownPlay, ownGameScore, ownScore)

	}

	for _, play := range gamePlanAsArray {
		opponentPlay := string(play[0])
		ownPlay := string(play[2])

		opponentGameScore, ownGameScore := calculateScores(opponentPlay, ownPlay)
		opponentScore += opponentGameScore
		ownScore += ownGameScore
		fmt.Println("Opponent: ", opponentPlay, opponentGameScore, opponentScore, " - own play: ", ownPlay, ownGameScore, ownScore)

	}

	ownScore = 0
	opponentScore = 0

	for _, play := range gamePlanAsArray {
		opponentPlay := string(play[0])
		ownPlay := string(play[2])

		opponentGameScore, ownGameScore := calculateScoresPart2(opponentPlay, ownPlay)
		opponentScore += opponentGameScore
		ownScore += ownGameScore
		fmt.Println("Opponent: ", opponentPlay, opponentGameScore, opponentScore, " - own play: ", ownPlay, ownGameScore, ownScore)

	}
}

func calculateScores(opponentPlay string, ownPlay string) (int, int) {

	opponentPoints := 0
	ownPoints := 0

	switch opponentPlay {
	case "A":
		opponentPoints += 1
	case "B":
		opponentPoints += 2
	case "C":
		opponentPoints += 3
	default:
	}

	switch ownPlay {
	case "X":
		ownPoints += 1
		switch opponentPoints {
		case 1:
			ownPoints += 3
			opponentPoints += 3
		case 2:
			opponentPoints += 6
		case 3:
			ownPoints += 6
		}

	case "Y":
		ownPoints += 2

		switch opponentPoints {
		case 2:
			ownPoints += 3
			opponentPoints += 3
		case 3:
			opponentPoints += 6
		case 1:
			ownPoints += 6
		}

	case "Z":
		ownPoints += 3
		switch opponentPoints {
		case 3:
			ownPoints += 3
			opponentPoints += 3
		case 1:
			opponentPoints += 6
		case 2:
			ownPoints += 6
		}
	default:
	}

	return opponentPoints, ownPoints
}

func calculateScoresPart2(opponentPlay string, ownPlay string) (int, int) {

	opponentPoints := 0
	ownPoints := 0

	switch opponentPlay {
	case "A":
		opponentPoints += 1
	case "B":
		opponentPoints += 2
	case "C":
		opponentPoints += 3
	default:
	}

	switch ownPlay {
	case "X":

		switch opponentPoints {
		case 1:
			ownPoints += 3
		case 2:
			ownPoints += 1
		case 3:
			ownPoints += 2
		}
		ownPoints += 0
		opponentPoints += 6

	case "Y":
		ownPoints += opponentPoints
		ownPoints += 3
		opponentPoints += 3

	case "Z":

		switch opponentPoints {
		case 1:
			ownPoints += 2
		case 2:
			ownPoints += 3
		case 3:
			ownPoints += 1
		}
		ownPoints += 6
		opponentPoints += 0

	default:
	}

	return opponentPoints, ownPoints
}
