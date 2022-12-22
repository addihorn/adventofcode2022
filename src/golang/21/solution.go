package main

import (
	"example/hello/src/golang/aocutils"
	"strings"
)

func main() {

	monkeys := map[string]*Monkey{}
	monkeyList := aocutils.ReadInput("input.txt")

	for _, monkey := range monkeyList {
		monkeyDetails := strings.Split(monkey, ":")
		monkeyName := strings.TrimSpace(monkeyDetails[0])
		monkeyOperation := strings.TrimSpace(monkeyDetails[1])

		monkeys[monkeyName] = NewMonkey(monkeyName, monkeyOperation)
	}

	//fmt.Println(monkeys["root"].Yell1(monkeys))
	monkeys["root"].dependsOnHuman(monkeys)
	monkeys["root"].Yell2(monkeys)
}
