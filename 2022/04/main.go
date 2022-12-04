package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

func parseInput(input string) []ElfPair {
	elves := strings.Split(strings.TrimSpace(input), "\n")

	output := make([]ElfPair, len(elves))

	for i, pair := range elves {
		pairs := strings.Split(pair, ",")
		first := rangeFrom(pairs[0])
		second := rangeFrom(pairs[1])
		output[i] = ElfPair{first, second}
	}

	return output
}

type Range struct {
	start int
	end   int
}

func rangeFrom(input string) Range {
	numbers := strings.Split(input, "-")
	first := util.Atoi(numbers[0])
	second := util.Atoi(numbers[1])
	return Range{first, second}
}

type ElfPair struct {
	first  Range
	second Range
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

func Part1(elves []ElfPair) int {
	count := 0
	for _, pair := range elves {
		if pair.first.contains(pair.second) || pair.second.contains(pair.first) {
			count++
		}
	}
	return count
}

func Part2(elves []ElfPair) int {
	count := 0
	for _, pair := range elves {
		if pair.first.overlaps(pair.second) || pair.second.overlaps(pair.first) {
			count++
		}
	}
	return count
}

func (r *Range) contains(other Range) bool {
	return r.start <= other.start && other.end <= r.end
}

func (r *Range) overlaps(other Range) bool {
	return r.contains(other) || other.start <= r.end && r.start <= other.end
}
