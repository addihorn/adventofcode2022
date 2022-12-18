package main

import (
	"math/big"
	"strconv"
	"strings"
)

type BigMonkey struct {
	Items           []big.Int
	Operation       string
	Test            string
	TestTrueTarget  int
	TestFalseTarget int
	ItemsInspected  uint64
}

func NewBigMonkey(monkeyDescr string) *BigMonkey {

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
	startingItems := []big.Int{}
	for _, item := range strings.Split(startingItemListAsString, ",") {
		itemWorrieLevel, _ := strconv.ParseInt(strings.TrimSpace(item), 10, 64)
		itemWorrieLevelBig := *big.NewInt(itemWorrieLevel)
		startingItems = append(startingItems, itemWorrieLevelBig)
	}

	worrienessOperation := strings.TrimSpace(strings.Split(monkeyDescDetails[2], ":")[1])
	throwText := strings.TrimSpace(strings.Split(monkeyDescDetails[3], ":")[1])

	testResult := strings.Split(monkeyDescDetails[4], " ")
	testTrueMonkey, _ := strconv.Atoi(testResult[len(testResult)-1])
	testResult = strings.Split(monkeyDescDetails[5], " ")
	testFalseMonkey, _ := strconv.Atoi(testResult[len(testResult)-1])

	return &BigMonkey{startingItems, worrienessOperation, throwText, testTrueMonkey, testFalseMonkey, 0}
}

func (this *BigMonkey) InspectAndThrowItem() (big.Int, int) {

	/* thanks to 	https://www.youtube.com/watch?v=k-c_TJ0j0W8 and
				https://www.reddit.com/r/adventofcode/comments/zih7gf/2022_day_11_part_2_what_does_it_mean_find_another/
	for the tips
	*/
	const common_divider_test int64 = 96577
	const common_divider int64 = 9699690

	reliefValue := big.NewInt(common_divider)

	inspectedWitemWorryLevel := this.Items[0]
	operationsAsArray := strings.Split(this.Operation, " ")
	//only indices 2-4 are relevant

	firstOperant := *big.NewInt(0)
	switch operationsAsArray[2] {
	case "old":
		firstOperant = inspectedWitemWorryLevel
	default:
		firstOperantAsInt, _ := strconv.ParseInt(operationsAsArray[2], 10, 64)
		firstOperant = *big.NewInt(firstOperantAsInt)
	}

	secondOperant := *big.NewInt(1)
	switch operationsAsArray[4] {
	case "old":
		secondOperant = inspectedWitemWorryLevel
	default:
		secondOperantAsInt, _ := strconv.ParseInt(operationsAsArray[4], 10, 64)
		secondOperant = *big.NewInt(secondOperantAsInt)
	}

	switch operationsAsArray[3] {
	case "+":
		inspectedWitemWorryLevel.Add(&firstOperant, &secondOperant)
	case "*":
		inspectedWitemWorryLevel.Mul(&firstOperant, &secondOperant)
	}

	this.ItemsInspected++
	inspectedWitemWorryLevel.Mod(&inspectedWitemWorryLevel, reliefValue)

	// test throw conditions
	// all throw tests, will test for mod
	testAsArr := strings.Split(this.Test, " ")
	divident, _ := strconv.ParseInt(testAsArr[len(testAsArr)-1], 10, 64)

	this.Items = this.Items[1:len(this.Items)]
	divValue := big.NewInt(0)
	divValue.Mod(&inspectedWitemWorryLevel, big.NewInt(divident))
	if divValue.Int64() == 0 {
		return inspectedWitemWorryLevel, this.TestTrueTarget
	} else {
		return inspectedWitemWorryLevel, this.TestFalseTarget
	}
}

func (this *BigMonkey) CatchItem(itemWorrieness big.Int) {
	this.Items = append(this.Items, itemWorrieness)
}
