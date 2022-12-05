package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

type Stack []byte

type Instruction struct {
	count int
	start int
	end   int
}

func parseInput(input string) ([]Stack, []Instruction) {
	initialState := strings.Split(strings.TrimRight(input, "\n"), "\n\n")

	firstSection := strings.Split(initialState[0], "\n")

	columns := (len(firstSection[0]) + 1) / 4

	stacks := make([]Stack, columns)
	for index, line := range firstSection {
		if index == len(firstSection)-1 {
			break
		}
		for i := 0; i < columns; i++ {
			potentialIndex := i*4 + 1
			char := line[potentialIndex]
			if char != ' ' {
				stacks[i] = append(stacks[i], char)
			}
		}
	}

	secondSection := strings.Split(initialState[1], "\n")

	instructions := make([]Instruction, len(secondSection))

	for i, instr := range secondSection {
		fields := strings.Fields(instr)
		instructions[i] = Instruction{
			count: util.Atoi(fields[1]),
			start: util.Atoi(fields[3]) - 1,
			end:   util.Atoi(fields[5]) - 1,
		}
	}

	return stacks, instructions
}

func main() {
	util.Time(doPart1, "Part 1")
	util.Time(doPart2, "Part 2")
}

func doPart1() {
	stacks, instructions := parseInput(input)
	fmt.Println("Part 1: ", Part1(stacks, instructions))
}

func doPart2() {
	stacks, instructions := parseInput(input)
	fmt.Println("Part 2: ", Part2(stacks, instructions))
}

func Part1(stacks []Stack, instructions []Instruction) string {
	for _, instr := range instructions {
		start := instr.start
		end := instr.end
		for i := 0; i < instr.count; i++ {
			value := stacks[start][0]
			stacks[start] = stacks[start][1:]
			stacks[end] = append(Stack{value}, stacks[end]...)
		}

	}

	var s strings.Builder
	for _, stack := range stacks {
		s.WriteByte(stack[0])
	}

	return s.String()
}

func Part2(stacks []Stack, instructions []Instruction) string {
	for _, instr := range instructions {
		start := instr.start
		end := instr.end

		list := make(Stack, instr.count)
		for i := 0; i < instr.count; i++ {
			value := stacks[start][0]
			stacks[start] = stacks[start][1:]
			list[i] = value
		}
		stacks[end] = append(list, stacks[end]...)
	}

	var s strings.Builder
	for _, stack := range stacks {
		s.WriteByte(stack[0])
	}

	return s.String()
}
