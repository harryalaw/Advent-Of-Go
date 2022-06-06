package main

import (
	util "2015/Util"
	"io/ioutil"
	"strconv"
	"strings"
)

func solve(part func() int) int {
	value := part()
	println(value)
	return value
}

func main() {
	util.Time(func() { solve(Part1) }, "Part1")
	util.Time(func() { solve(Part2) }, "Part2")
}

func parseInput() []*instruction {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	instructions := make([]*instruction, 0, len(lines))

	for _, line := range lines {
		instruction := MapStringToInstruction(line)
		instructions = append(instructions, instruction)
	}

	return instructions
}

type instruction struct {
	operator string
	args     []string
	value    string
}

func MapStringToInstruction(line string) *instruction {
	fields := strings.Fields(line)

	switch len(fields) {
	case 3:
		return &instruction{
			args:     fields[0:1],
			operator: "ASSIGN",
			value:    fields[2],
		}
	case 4:
		return &instruction{
			args:     fields[1:2],
			operator: fields[0],
			value:    fields[3],
		}
	case 5:
		return &instruction{
			args:     []string{fields[0], fields[2]},
			operator: fields[1],
			value:    fields[4],
		}
	}

	return &instruction{}
}
func isNumeric(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

func allArgsAreKnown(args *[]string, knownValues *map[string]int) bool {
	for _, arg := range *args {
		_, isPresent := (*knownValues)[arg]
		if !(isPresent || isNumeric(arg)) {
			return false
		}
	}

	return true
}

func fetchValue(arg string, knownValues *map[string]int) int {
	value, isPresent := (*knownValues)[arg]
	if isPresent {
		return value
	}
	num, _ := strconv.Atoi(arg)

	return num
}

func Part1() int {

	instructions := parseInput()

	knownValues := make(map[string]int)

	return CoreLogic(instructions, knownValues)
}

func Part2() int {
	instructions := parseInput()

	knownValues := make(map[string]int)
	knownValues["b"] = 46065

	return CoreLogic(instructions, knownValues)
}

func CoreLogic(instructions []*instruction, knownValues map[string]int) int {
	for len(knownValues) < len(instructions) {
		for _, instruction := range instructions {
			_, isPresent := knownValues[instruction.value]
			if isPresent {
				continue
			}
			if !allArgsAreKnown(&instruction.args, &knownValues) {
				continue
			}
			switch instruction.operator {
			case "ASSIGN":
				num := fetchValue(instruction.args[0], &knownValues)
				knownValues[instruction.value] = num
			case "NOT":
				num := fetchValue(instruction.args[0], &knownValues)

				knownValues[instruction.value] = 65535 ^ num
			case "OR":
				num1 := fetchValue(instruction.args[0], &knownValues)
				num2 := fetchValue(instruction.args[1], &knownValues)

				knownValues[instruction.value] = num1 | num2
			case "AND":
				num1 := fetchValue(instruction.args[0], &knownValues)
				num2 := fetchValue(instruction.args[1], &knownValues)

				knownValues[instruction.value] = num1 & num2
			case "RSHIFT":
				num1 := fetchValue(instruction.args[0], &knownValues)
				num2 := fetchValue(instruction.args[1], &knownValues)

				knownValues[instruction.value] = num1 >> num2
			case "LSHIFT":
				num1 := fetchValue(instruction.args[0], &knownValues)
				num2 := fetchValue(instruction.args[1], &knownValues)

				knownValues[instruction.value] = num1 << num2
			}
		}
	}
	return knownValues["a"]
}
