package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Print(err)
	}

	// fmt.Println(b)
	textFileContent := string(b)
	//fmt.Println(textFileContent)

	elveList := strings.Split(textFileContent, "\n\n")
	fmt.Println("Number of elves carrying stuff: ", len(elveList))

	caloriesCarried := []int{}

	for _, s := range elveList {
		caloriesCarried = append(caloriesCarried, sum(strings.Split(s, "\n")))
	}

	sort.Sort(
		sort.Reverse(
			sort.IntSlice(caloriesCarried),
		),
	)

	fmt.Println("Max Calories carried by one elve: ", caloriesCarried[0])

	carriedByXStrongest := caloriesCarried
	for i := 1; i < len(caloriesCarried); i++ {
		carriedByXStrongest[i] = carriedByXStrongest[i-1] + carriedByXStrongest[i]
	}

	fmt.Println("Calories carried by 3 strongest elves: ", carriedByXStrongest[3-1])
}

func sum(array []string) int {
	result := 0
	for _, v := range array {
		intVar, err := strconv.Atoi(v)
		if err != nil {
			//fmt.Print(err)
		}
		result += intVar
	}
	return result
}
