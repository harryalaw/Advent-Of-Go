package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

type Operation func(int) int
type Predicate func(int) bool

type Monkey struct {
	items        []int
	op           Operation
	test         Predicate
	targets      [2]int
	interactions int
}

func parseOperation(input string) Operation {
	symbols := strings.Fields(input)
	operator := symbols[3]
	if symbols[4] == "old" {
		if operator == "+" {
			return func(a int) int {
				return a + a
			}
		}
		if operator == "*" {
			return func(a int) int {
				return a * a
			}
		}
	}
	amount := util.Atoi(symbols[4])

	switch operator {
	case "+":
		return func(a int) int {
			return a + amount
		}
	case "*":
		return func(a int) int {
			return a * amount
		}
	default:
		return func(a int) int {
			return -1
		}
	}
}

func parsePredicate(input string) (Predicate, int) {
	symbols := strings.Fields(input)
	amount := util.Atoi(symbols[2])

	return func(a int) bool {
		return a%amount == 0
	}, amount
}

func parseInput(input string) ([]*Monkey, int) {
	monkeyText := strings.Split(strings.TrimSpace(input), "\n\n")
	monkeys := make([]*Monkey, len(monkeyText))
	lcm := 1

	for i, monkey := range monkeyText {
		lines := strings.Split(strings.TrimSpace(monkey), "\n")

		items := strings.Split(strings.Split(lines[1], ": ")[1], ", ")
		intItems := make([]int, len(items))
		for j, item := range items {
			intItems[j] = util.Atoi(item)
		}

		op := parseOperation(strings.Split(lines[2], ": ")[1])

		test, factor := parsePredicate(strings.Split(lines[3], ": ")[1])

		lcm = util.Lcm(lcm, factor)

		targets := [2]int{}
		targets[0] = util.Atoi(strings.Fields(lines[4])[5])
		targets[1] = util.Atoi(strings.Fields(lines[5])[5])

		monkeys[i] = &Monkey{
			items:        intItems,
			op:           op,
			test:         test,
			targets:      targets,
			interactions: 0,
		}
	}

	return monkeys, lcm
}

func main() {
	util.Time(doPart1, "Part 1")
	util.Time(doPart2, "Part 2")
}

func doPart1() {
	data, _ := parseInput(input)
	fmt.Println("Part 1: ", Part1(data))
}

func doPart2() {
	data, lcm := parseInput(input)
	fmt.Println("Part 2: ", Part2(data, lcm))
}

func Part1(monkeys []*Monkey) int {
	/*
		 for 20 rounds
		 	for each monkey
				Inspect each item
					increment count
				do operation
				divide by 3
				throw to next monkey

		mult two most active monkeys
	*/

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				monkey.interactions += 1
				item = monkey.op(item)
				item /= 3
				var targetMonkey int
				if monkey.test(item) {
					targetMonkey = monkey.targets[0]
				} else {
					targetMonkey = monkey.targets[1]
				}
				monkeys[targetMonkey].items = append(monkeys[targetMonkey].items, item)
			}
			monkey.items = []int{}
		}
	}

	mostInteractions := []int{0, 0}
	for _, monkey := range monkeys {
		if monkey.interactions > mostInteractions[0] {
			mostInteractions[1] = mostInteractions[0]
			mostInteractions[0] = monkey.interactions
		} else if monkey.interactions > mostInteractions[1] {
			mostInteractions[1] = monkey.interactions
		}
	}

	return mostInteractions[0] * mostInteractions[1]
}

func printMonkeys(monkeys []*Monkey) {
	for i, monkey := range monkeys {
		fmt.Printf("Monkey %d: ", i)
		for _, item := range monkey.items {
			fmt.Printf("%d ", item)
		}
		fmt.Printf("\n")
	}
}

func Part2(monkeys []*Monkey, lcm int) int {
	// not 20439655544
	// how to handle overflows?
	// a bigint package?
	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				monkey.interactions += 1
				item = monkey.op(item)
				item = item % lcm
				var targetMonkey int
				if monkey.test(item) {
					targetMonkey = monkey.targets[0]
				} else {
					targetMonkey = monkey.targets[1]
				}
				monkeys[targetMonkey].items = append(monkeys[targetMonkey].items, item)
			}
			monkey.items = []int{}
		}
	}

	mostInteractions := []int{0, 0}
	for _, monkey := range monkeys {
		if monkey.interactions > mostInteractions[0] {
			mostInteractions[1] = mostInteractions[0]
			mostInteractions[0] = monkey.interactions
		} else if monkey.interactions > mostInteractions[1] {
			mostInteractions[1] = monkey.interactions
		}
	}

	return mostInteractions[0] * mostInteractions[1]
}
