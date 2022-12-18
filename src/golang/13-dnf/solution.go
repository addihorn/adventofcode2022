package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"regexp"
	"strings"
)

func main() {

	pairsOfPackets := aocutils.ReadInputWithDelimeter("test-input.txt", "\n\n")

	for _, packagePair := range pairsOfPackets {
		pairDetails := strings.Split(packagePair, "\n")
		re := regexp.MustCompile(`\[.*\]`)

		fmt.Println(pairDetails)
		for _, packageDetails := range pairDetails {
			fmt.Println("Regex for", packageDetails[1:len(packageDetails)-1], ":", len(re.FindAllString(packageDetails[1:len(packageDetails)-1], -1)))
		}
	}

}
