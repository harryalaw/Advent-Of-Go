package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

func parseInput(input string) string {
	return strings.TrimSpace(input)
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

func hasDupe(group []rune) bool {
	for i := 0; i < len(group); i++ {
		for j := i + 1; j < len(group); j++ {
			if group[i] == group[j] {
				return true
			}
		}
	}
	return false
}

func Part1(data string) int {
	lastChars := make([]rune, 3)

	for i, char := range data {
		if i < 3 {
			lastChars[i] = char
		} else {

			dupe := false
			for _, c := range lastChars {
				if c == char {
					dupe = true
				}
			}
			if !dupe && !hasDupe(lastChars) {
				return i + 1
			}
			lastChars = []rune{lastChars[1], lastChars[2], char}
		}
	}

	return -1
}

func Part2(data string) int {
	lastChars := make([]rune, 13)

	for i, char := range data {
		if i < 13 {
			lastChars[i] = char
		} else {

			dupe := false
			for _, c := range lastChars {
				if c == char {
					dupe = true
				}
			}
			if !dupe && !hasDupe(lastChars) {
				return i + 1
			}

			for i := 1; i < 13; i++ {
				lastChars[i-1] = lastChars[i]
			}
			lastChars[12] = char
		}
	}

	return -1
}
