package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

type Monkey interface {
	monkey()
}

type MathMonkey struct {
	left      string
	right     string
	operation func(left, right int) int
}

func (mm MathMonkey) monkey() {}

type NumberMonkey struct {
	value int
}

func (nm NumberMonkey) monkey() {}

type Monkeys map[string]Monkey

func add(left, right int) int {
	return left + right
}

func minus(left, right int) int {
	return left - right
}

func mult(left, right int) int {
	return left * right
}

func div(left, right int) int {
	return left / right
}
func eq(left, right int) int {
	if left == right {
		return 1
	}
	return 0
}

func parseOp(operation string) func(int, int) int {
	switch operation {
	case "+":
		return add
	case "-":
		return minus
	case "*":
		return mult
	case "/":
		return div
	default:
		panic("Not a valid operation")
	}
}

func parseInput(input string) Monkeys {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	monkeys := map[string]Monkey{}

	for _, line := range lines {
		name := line[0:4]
		fields := strings.Fields(line)
		if len(fields) == 2 {
			monkey := NumberMonkey{value: util.Atoi(fields[1])}
			monkeys[name] = monkey
		} else {
			left := fields[1]
			operation := parseOp(fields[2])
			right := fields[3]
			monkey := MathMonkey{left, right, operation}
			monkeys[name] = monkey
		}
	}

	return monkeys
}

func main() {
	util.Time(doPart1, "Part 1")
	util.Time(doPart2, "Part 2")
}

func doPart1() {
	data := parseInput(input)
	fmt.Println("Part 1: ", Part1(data))
}

func doPart2() {
	data := parseInput(input)
	fmt.Println("Part 2: ", Part2(data))
}

func Part1(monkeys Monkeys) int {
	values := map[string]int{}

	for len(values) != len(monkeys) {
		for name, monkey := range monkeys {
			if _, ok := values[name]; ok {
				continue
			}

			switch monkey := monkey.(type) {
			case NumberMonkey:
				values[name] = monkey.value
			case MathMonkey:
				leftVal, ok := values[monkey.left]
				if !ok {
					continue
				}
				rightVal, ok := values[monkey.right]
				if !ok {
					continue
				}
				values[name] = monkey.operation(leftVal, rightVal)
			}
		}
	}

	return values["root"]
}

func part1Alg(monkeys Monkeys) map[string]int {
	values := map[string]int{}
	for len(values) != len(monkeys) {
		for name, monkey := range monkeys {
			if _, ok := values[name]; ok {
				continue
			}

			switch monkey := monkey.(type) {
			case NumberMonkey:
				values[name] = monkey.value
			case MathMonkey:
				leftVal, ok := values[monkey.left]
				if !ok {
					continue
				}
				rightVal, ok := values[monkey.right]
				if !ok {
					continue
				}
				values[name] = monkey.operation(leftVal, rightVal)
			}
		}
	}
	return values
}

func Part2(monkeys Monkeys) int {
	rootMonkey := monkeys["root"]
	newRoot := rootMonkey.(MathMonkey)
	newRoot.operation = eq
	monkeys["root"] = newRoot

	leftMonkey := newRoot.left
	rightMonkey := newRoot.right

	var humnSide string
	monkeys["humn"] = NumberMonkey{value: 0}
	values := part1Alg(monkeys)
	leftVal1 := values[leftMonkey]
	rightVal1 := values[rightMonkey]

	monkeys["humn"] = NumberMonkey{value: 1000000000}
	values = part1Alg(monkeys)
	leftVal2 := values[leftMonkey]
	rightVal2 := values[rightMonkey]

	var target int
	var isDecreasing bool
	if leftVal1 == leftVal2 {
		humnSide = rightMonkey
		target = leftVal1
		isDecreasing = rightVal1 > rightVal2

	} else {
		humnSide = leftMonkey
		target = rightVal1
		isDecreasing = leftVal1 > leftVal2
	}

	// or we do a binary search to find humn value
	// have low = 0, high = 1_000_000_000_000
	low := 0
	high := 1_000_000_000_000_000

	goodGuess := -1
	for {
		mid := (low + high) / 2
		monkeys["humn"] = NumberMonkey{value: mid}
		values := part1Alg(monkeys)
		if values["root"] == 1 {
			goodGuess = mid
			break
		}

		newResult := values[humnSide]
		if isDecreasing {
			if target > newResult {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target < newResult {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}

	lastGoodGuess := goodGuess
	for {
		goodGuess--
		monkeys["humn"] = NumberMonkey{value: goodGuess}
		values := part1Alg(monkeys)
		if values["root"] == 1 {
			lastGoodGuess = goodGuess
			continue
		}
		break
	}

	return lastGoodGuess
}
