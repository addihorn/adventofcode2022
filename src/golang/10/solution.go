package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	values := []int{1}

	program := aocutils.ReadInput("input.txt")

	for _, lineOfCode := range program {
		switch {
		case lineOfCode == "noop":
			values = append(values, values[len(values)-1])
		case strings.Split(lineOfCode, " ")[0] == "addx":
			valueToAdd, _ := strconv.Atoi(strings.Split(lineOfCode, " ")[1])

			values = append(values, values[len(values)-1])
			values = append(values, values[len(values)-1]+valueToAdd)
		default:

		}
	}
	fmt.Println(values)

	signalStrength := 0
	for i := 20; i < len(values); i += 40 {
		signalStrength += (values[i-1] * i)
	}
	fmt.Println("SignalStrength:", signalStrength)

	for row := 0; row < 6; row++ {
		for pixel := 0; pixel < 40; pixel++ {
			spriteToValidate := row*40 + pixel
			spritePosHorizontal := values[spriteToValidate]

			if pixel >= spritePosHorizontal-1 && pixel <= spritePosHorizontal+1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ") //print Space instead of Point for readability
			}
		}
		fmt.Println()
	}

}
