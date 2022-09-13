package main

import (
	"github.com/harryalaw/advent-of-go/util"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	util.Time(part1, "Part1")
	util.Time(part2, "Part2")
}

func part1() {
	instructions := parseInput()

	_, b := performOperations(0, 0, instructions)

	fmt.Println(b)
}

func part2() {
	instructions := parseInput()

	_, b := performOperations(1, 0, instructions)

	fmt.Println(b)
}

type instruction struct {
	name     string
	value    string
	third    int
	jmpValue int
}

func parseInput() []instruction {
	data, err := ioutil.ReadFile("data.txt")

	if err != nil {
		panic(err)
	}

	noCommas := strings.ReplaceAll(string(data), ",", "")
	lines := strings.Split(noCommas, "\n")

	instructions := make([]instruction, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)
		inst := instruction{
			name:  parts[0],
			value: parts[1],
		}
		if len(parts) == 3 {
			inst.third = util.Atoi(parts[2])
		}
		instructions[i] = inst
	}
	return instructions
}

func performOperations(a, b uint64, instructions []instruction) (uint64, uint64) {
	ptr := 0

	for ptr < len(instructions) {
		inst := instructions[ptr]
		a, b, ptr = doInstruction(a, b, ptr, inst)
	}

	return a, b
}

func doInstruction(a, b uint64, ptr int, inst instruction) (uint64, uint64, int) {
	var value *uint64
	if inst.value == "a" {
		value = &a
	} else {
		value = &b
	}

	switch inst.name {
	case "inc":
		*value++
		ptr++
	case "hlf":
		*value >>= 1
		ptr++
	case "tpl":
		*value *= 3
		ptr++
	case "jmp":
		ptr += util.Atoi(inst.value)
	case "jie":
		if *value%2 == 0 {
			ptr += inst.third
		} else {
			ptr++
		}
	case "jio":
		if *value == 1 {
			ptr += inst.third
		} else {
			ptr++
		}
	}

	return a, b, ptr
}
