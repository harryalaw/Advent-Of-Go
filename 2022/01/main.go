package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

func main() {
	util.Time(doPart1, "Part 1")
	util.Time(doPart2, "Part 2")
}

func parseInput(input string) [][]int {
	data := strings.Split(strings.TrimSpace(input), "\n\n")

	out := make([][]int, len(data))

	for i, line := range data {
		calories := strings.Split(line, "\n")
		outputLine := make([]int, len(calories))
		for j, val := range calories {
			outputLine[j] = util.Atoi(val)
		}
		out[i] = outputLine
	}

	return out
}

func doPart1() {
	elves := parseInput(input)

	mostCalories := Part1(elves)
	fmt.Println("Part 1: ", mostCalories)
}

func doPart2() {
	elves := parseInput(input)

	mostCalories := Part2(elves)
	fmt.Println("Part 2: ", mostCalories)
}

func Part1(elves [][]int) int {
	mostCalories := 0

	for _, elf := range elves {
		totalCalories := util.IntSum(elf...)
		mostCalories = util.IntMax(mostCalories, totalCalories)
	}

	return mostCalories
}

func Part2(elves [][]int) int {
	mostCalories := []int{0, 0, 0}

	for _, elf := range elves {
		totalCalories := util.IntSum(elf...)
		if totalCalories >= mostCalories[0] {
			mostCalories[2] = mostCalories[1]
			mostCalories[1] = mostCalories[0]
			mostCalories[0] = totalCalories
		} else if totalCalories >= mostCalories[1] {
			mostCalories[2] = mostCalories[1]
			mostCalories[1] = totalCalories
		} else if totalCalories >= mostCalories[2] {
			mostCalories[2] = totalCalories
		}
	}

	return util.IntSum(mostCalories...)
}
