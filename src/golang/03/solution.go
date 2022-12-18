package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	rucksackContents := readInput()

	prioritySum := 0

	for _, rucksack := range rucksackContents {
		firstCompartment := rucksack[0:(len(rucksack) / 2)]
		secondCompartment := rucksack[(len(rucksack) / 2):len(rucksack)]
		repeatedChar := findRepeatedChar(firstCompartment, secondCompartment)

		priority := calculatePriorityOfChar(repeatedChar)
		prioritySum += priority
		fmt.Println(firstCompartment, len(firstCompartment), secondCompartment, len(secondCompartment))
		fmt.Println("same Item:", string(repeatedChar), "item priority:", priority, "Sum of priorities:", prioritySum)

	}

	prioritySum = 0
	for i := 1; i <= len(rucksackContents)/3; i++ {
		groupRucksacks := rucksackContents[3*i-3 : 3*i]
		repeatedChar := findReaptedChar(groupRucksacks)
		priority := calculatePriorityOfChar(repeatedChar)
		prioritySum += priority
		fmt.Println(groupRucksacks, "same Item:", string(repeatedChar), "priority:", priority, "Sum:", prioritySum)
	}

}

func readInput() []string {

	b, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileAsString := string(b)
	return strings.Split(fileAsString, "\n")
}

func findRepeatedChar(first string, second string) rune {
	for _, char := range first {
		if len(strings.Split(second, string(char))) > 1 {
			return char
		}
	}
	return 0
}

func findReaptedChar(charArray []string) rune {
	for _, char := range charArray[0] {
		if len(strings.Split(charArray[1], string(char))) > 1 &&
			len(strings.Split(charArray[2], string(char))) > 1 {
			return char
		}
	}
	return 0
}

func calculatePriorityOfChar(char rune) int {
	var priority int
	priority = 0
	if 65 <= char && char <= 90 {
		priority = int(char) - 38
	}
	if 97 <= char && char <= 122 {
		priority = int(char) - 96
	}
	return priority
}
