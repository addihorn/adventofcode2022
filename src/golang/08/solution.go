package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
)

func main() {
	treesByRow := aocutils.ReadInput("input.txt")

	visibleTrees := 0
	maxScenicScore := 0
	for rowNumber, treeRow := range treesByRow {
		for treeNumber, tree := range treeRow {
			treeVisibility, treeScenicScore := isTreeVisible(int(tree), rowNumber, treeNumber, treesByRow)
			visibleTrees += treeVisibility
			maxScenicScore = max(maxScenicScore, treeScenicScore)
		}
	}

	fmt.Println("Visible Trees:", visibleTrees, "max Scenic Score:", maxScenicScore)
}

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

func isTreeVisible(treeHight int, treeRow int, treeColumn int, forrest []string) (int, int) {

	visibleFromTop := true
	visibleFromBottom := true
	visibleFromRight := true
	visibleFromLeft := true

	scenicScore := 1
	treesCounted := 0
	//check top

	for i := treeRow - 1; i >= 0; i-- {
		treesCounted++
		if int(forrest[i][treeColumn]) >= treeHight {
			visibleFromTop = false
			break
		}
	}

	scenicScore *= treesCounted
	treesCounted = 0

	//check bottom
	for _, row := range forrest[treeRow+1 : len(forrest)] {
		treesCounted++
		if int(row[treeColumn]) >= treeHight {
			visibleFromBottom = false
			break
		}
	}
	scenicScore *= treesCounted
	treesCounted = 0

	//check left
	for i := treeColumn - 1; i >= 0; i-- {
		treesCounted++
		if int(forrest[treeRow][i]) >= treeHight {
			visibleFromLeft = false
			break
		}
	}

	scenicScore *= treesCounted
	treesCounted = 0

	//check right
	for _, col := range forrest[treeRow][treeColumn+1 : len(forrest[treeRow])] {
		treesCounted++
		if int(col) >= treeHight {
			visibleFromRight = false
			break
		}
	}
	scenicScore *= treesCounted
	treesCounted = 0

	if visibleFromTop || visibleFromBottom || visibleFromRight || visibleFromLeft {
		return 1, scenicScore
	} else {
		return 0, scenicScore
	}

}
