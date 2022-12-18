package main

import (
	"strconv"
	"strings"
)

type Monkey struct {
	Items           []uint64
	Operation       string
	Test            string
	TestTrueTarget  int
	TestFalseTarget int
	ItemsInspected  uint64
}

func NewMonkey(monkeyDescr string) *Monkey {

	/*
		monkeyDesc should look like following:
		Id:
			Starting items: <list of Items>
			Operation: <Operation>
			Test: <Test>
				If true: throw to monkey x
				If False: throw to monkey y

	*/

	monkeyDescDetails := strings.Split(monkeyDescr, "\n")

	startingItemListAsString := strings.Split(monkeyDescDetails[1], ":")[1]
	startingItems := []uint64{}
	for _, item := range strings.Split(startingItemListAsString, ",") {
		itemWorrieLevel, _ := strconv.ParseUint(strings.TrimSpace(item), 10, 64)
		startingItems = append(startingItems, itemWorrieLevel)
	}

	worrienessOperation := strings.TrimSpace(strings.Split(monkeyDescDetails[2], ":")[1])
	throwText := strings.TrimSpace(strings.Split(monkeyDescDetails[3], ":")[1])

	testResult := strings.Split(monkeyDescDetails[4], " ")
	testTrueMonkey, _ := strconv.Atoi(testResult[len(testResult)-1])
	testResult = strings.Split(monkeyDescDetails[5], " ")
	testFalseMonkey, _ := strconv.Atoi(testResult[len(testResult)-1])

	return &Monkey{startingItems, worrienessOperation, throwText, testTrueMonkey, testFalseMonkey, 0}
}

func (this *Monkey) InspectAndThrowItem() (uint64, int) {

	const reliefValue uint64 = 3

	inspectedWitemWorryLevel := this.Items[0]
	operationsAsArray := strings.Split(this.Operation, " ")
	//only indices 2-4 are relevant
	var firstOperant uint64
	firstOperant = 0
	switch operationsAsArray[2] {
	case "old":
		firstOperant = inspectedWitemWorryLevel
	default:
		firstOperant, _ = strconv.ParseUint(operationsAsArray[2], 10, 64)
	}
	var secondOperant uint64
	secondOperant = 0
	switch operationsAsArray[4] {
	case "old":
		secondOperant = inspectedWitemWorryLevel
	default:
		secondOperant, _ = strconv.ParseUint(operationsAsArray[4], 10, 64)
	}

	switch operationsAsArray[3] {
	case "+":
		inspectedWitemWorryLevel = firstOperant + secondOperant
	case "*":
		inspectedWitemWorryLevel = firstOperant * secondOperant
	}

	this.ItemsInspected++
	inspectedWitemWorryLevel = inspectedWitemWorryLevel / reliefValue

	// test throw conditions
	// all throw tests, will test for mod
	testAsArr := strings.Split(this.Test, " ")
	divident, _ := strconv.ParseUint(testAsArr[len(testAsArr)-1], 10, 64)

	this.Items = this.Items[1:len(this.Items)]
	if inspectedWitemWorryLevel%divident == 0 {
		return inspectedWitemWorryLevel, this.TestTrueTarget
	} else {
		return inspectedWitemWorryLevel, this.TestFalseTarget
	}
}

func (this *Monkey) CatchItem(itemWorrieness uint64) {
	this.Items = append(this.Items, itemWorrieness)
}
