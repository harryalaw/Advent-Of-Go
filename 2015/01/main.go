package main

import (
	_ "embed"
	"github.com/harryalaw/advent-of-go/util"
	"fmt"
)

func main() {
	util.Time(func() { fmt.Println(part1()) }, "Part1")
	util.Time(func() { fmt.Println(part2()) }, "Part2")
}

//go:embed "data.txt"
var input string

func part1() int {
	floor := 0

	for _, char := range input {
		if char == '(' {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

func part2() int {
	floor := 0

	for index, char := range input {
		if char == '(' {
			floor++
		} else {
			floor--
		}
		if floor == -1 {
			return index + 1
		}
	}
	return -1
}
