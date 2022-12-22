package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Monkey struct {
	Name           string
	Operation      string
	Value          int64
	DependsOnHuman bool
}

func NewMonkey(name string, operation string) *Monkey {

	value, err := strconv.ParseInt(operation, 10, 64)

	if err == nil {
		return &Monkey{name, operation, value, false}
	}
	return &Monkey{name, operation, int64(math.NaN()), false}
}

func (this *Monkey) Yell1(monkeyGroup map[string]*Monkey) int64 {
	if this.Value != int64(math.NaN()) {
		return this.Value
	}

	operants := strings.Split(this.Operation, " ")

	monkey1 := monkeyGroup[operants[0]].Yell1(monkeyGroup)
	monkey2 := monkeyGroup[operants[2]].Yell1(monkeyGroup)

	switch operants[1] {
	case "/":
		this.Value = monkey1 / monkey2
	case "*":
		this.Value = monkey1 * monkey2
	case "+":
		this.Value = monkey1 + monkey2
	case "-":
		this.Value = monkey1 - monkey2
	}

	return this.Value

}

func (this *Monkey) Yell2(monkeyGroup map[string]*Monkey) {

	operants := strings.Split(this.Operation, " ")

	var monkey1 int64
	var monkey2 int64

	monkey1HumanDep := monkeyGroup[operants[0]].DependsOnHuman
	monkey2HumanDep := monkeyGroup[operants[2]].DependsOnHuman

	switch {
	case monkey1HumanDep:
		monkey2 = monkeyGroup[operants[2]].Yell1(monkeyGroup)
		monkeyGroup[operants[0]].calculateHumnValue(monkey2, monkeyGroup)
	case monkey2HumanDep:
		monkey1 = monkeyGroup[operants[0]].Yell1(monkeyGroup)
		monkeyGroup[operants[2]].calculateHumnValue(monkey1, monkeyGroup)
	}
	fmt.Println(monkeyGroup["humn"].Value)
}

func (this *Monkey) dependsOnHuman(monkeyGroup map[string]*Monkey) bool {
	if this.Name == "humn" {
		this.DependsOnHuman = true
		return true
	}
	if this.Value != int64(math.NaN()) {
		this.DependsOnHuman = false
		return false
	}

	operants := strings.Split(this.Operation, " ")

	monkey1 := monkeyGroup[operants[0]].dependsOnHuman(monkeyGroup)
	monkey2 := monkeyGroup[operants[2]].dependsOnHuman(monkeyGroup)

	if monkey1 || monkey2 {
		this.DependsOnHuman = true
	} else {
		this.Yell1(monkeyGroup)
		this.DependsOnHuman = false
	}

	return this.DependsOnHuman
}

func (this *Monkey) calculateHumnValue(expectedResult int64, monkeyGroup map[string]*Monkey) {
	if this.Name == "humn" {
		this.Value = expectedResult
		return
	}

	operants := strings.Split(this.Operation, " ")

	var monkey1 int64
	var monkey2 int64

	unknownMonkey := 0
	switch {
	case monkeyGroup[operants[0]].DependsOnHuman:
		unknownMonkey = 1
	case monkeyGroup[operants[2]].DependsOnHuman:
		unknownMonkey = 2
	}

	if unknownMonkey == 1 {
		monkey2 = monkeyGroup[operants[2]].Value

		switch operants[1] {
		case "/":
			monkey1 = expectedResult * monkey2
		case "*":
			monkey1 = expectedResult / monkey2
		case "+":
			monkey1 = expectedResult - monkey2
		case "-":
			monkey1 = expectedResult + monkey2
		}
		this.Value = expectedResult
		monkeyGroup[operants[0]].calculateHumnValue(monkey1, monkeyGroup)
		return
	}

	if unknownMonkey == 2 {
		monkey1 = monkeyGroup[operants[0]].Value

		switch operants[1] {
		case "/":
			monkey2 = monkey1 / expectedResult
		case "*":
			monkey2 = expectedResult / monkey1
		case "+":
			monkey2 = expectedResult - monkey1
		case "-":
			monkey2 = monkey1 - expectedResult
		}
		monkeyGroup[operants[2]].calculateHumnValue(monkey2, monkeyGroup)
		this.Value = expectedResult

		return
	}

}
