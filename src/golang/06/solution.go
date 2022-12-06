package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	messages := readInput("input.txt")

	for _, message := range messages {

		startingPosition, startPackage := getStartOfMessageMarker(message, 4)
		messageStartsAt, messageDetails := getStartOfMessageMarker(message, 14)
		fmt.Println(message, startingPosition, startPackage, messageStartsAt, messageDetails)

	}

}

func readInput(fileName string) []string {

	b, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	fileAsString := string(b)
	return strings.Split(fileAsString, "\n")
}

func getStartPackage(message string) (int, string) {
	for i := 3; i < len(message); i++ {
		char1 := message[i-3]
		char2 := message[i-2]
		char3 := message[i-1]
		char4 := message[i]

		if char1 != char2 &&
			char1 != char3 &&
			char1 != char4 &&
			char2 != char3 &&
			char2 != char4 &&
			char3 != char4 {
			return i, message[i : i+4]
		}

	}
	return 0, ""
}

func getStartOfMessageMarker(message string, markerLength int) (int, string) {

	for i := markerLength; i < len(message); i++ {

		potentialMarker := message[i-markerLength : i]

		markerFound := true
		for _, char := range potentialMarker {
			if len(strings.Split(potentialMarker, string(char))) > 2 {
				markerFound = false
				break
			}
		}

		if markerFound {
			return i, potentialMarker
		}

	}
	return 0, ""

}
