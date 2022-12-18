package main

import (
	"example/hello/src/golang/aocutils"
	"fmt"
	"sort"
)

func main() {

	roundsToPlay := 20

	monkeys := []*Monkey{}
	bigMonkeys := []*BigMonkey{}

	monkeyDescriptions := aocutils.ReadInputWithDelimeter("input.txt", "Monkey")
	for _, monkeyDescription := range monkeyDescriptions[1:len(monkeyDescriptions)] {
		monkeys = append(monkeys, NewMonkey(monkeyDescription))
		bigMonkeys = append(bigMonkeys, NewBigMonkey(monkeyDescription))
	}
	/*
		for _, monkey := range monkeys {
			fmt.Println(monkey)
		}
	*/
	for i := 0; i < roundsToPlay; i++ {
		for _, monkey := range monkeys {
			for len(monkey.Items) > 0 {
				itemLevel, targetMonkey := monkey.InspectAndThrowItem()
				monkeys[targetMonkey].CatchItem(itemLevel)

				//fmt.Println("Monkey", mNo, "thows item with value", itemLevel, "to Monkey", targetMonkey)
			}
		}
	}
	itemsInspectedPerMonkey := []int{}

	for _, monkey := range monkeys {
		//fmt.Println(monkey)
		itemsInspectedPerMonkey = append(itemsInspectedPerMonkey, int(monkey.ItemsInspected))
	}

	sort.Sort(
		sort.Reverse(
			sort.IntSlice(itemsInspectedPerMonkey),
		),
	)
	fmt.Println(itemsInspectedPerMonkey)
	fmt.Println("Monkey Business Value:", itemsInspectedPerMonkey[0]*itemsInspectedPerMonkey[1])

	fmt.Println("#### NOW with Big Numbers ######")
	roundsToPlay = 10000

	for i := 0; i < roundsToPlay; i++ {
		for _, monkey := range bigMonkeys {
			for len(monkey.Items) > 0 {
				itemLevel, targetMonkey := monkey.InspectAndThrowItem()
				bigMonkeys[targetMonkey].CatchItem(itemLevel)

			}
		}
	}

	itemsInspectedPerMonkey = []int{}

	for _, monkey := range bigMonkeys {
		itemsInspectedPerMonkey = append(itemsInspectedPerMonkey, int(monkey.ItemsInspected))
	}

	sort.Sort(
		sort.Reverse(
			sort.IntSlice(itemsInspectedPerMonkey),
		),
	)
	fmt.Println(itemsInspectedPerMonkey)
	fmt.Println("Big Monkey Business Value:", itemsInspectedPerMonkey[0]*itemsInspectedPerMonkey[1])

}
