package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

func parseInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
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

func Part1(data []string) int {
	total := 0

	for _, line := range data {
		length := len(line)
		firstHalf := line[0 : length/2]
		secondHalf := line[length/2:]
		firstLetters := getLetters(firstHalf)
		secondLetters := getLetters(secondHalf)

		total += findCommonLetter(firstLetters, secondLetters) + 1
	}

	return total
}

func findCommonLetter(a, b [52]bool) int {
	for i, value := range a {
		if value && b[i] {
			return i
		}
	}

	panic("No match")
}

func findCommonLetters(a, b [52]bool) [52]bool {
	out := [52]bool{}
	for i, value := range a {
		if value && b[i] {
			out[i] = true
		}
	}
	return out
}

func getLetters(item string) [52]bool {
	out := [52]bool{}
	for _, char := range item {
		if char >= 97 {
			out[char-97] = true
		} else {
			out[char-65+26] = true
		}
	}

	return out
}

func Part2(data []string) int {
	total := 0

	for i := 0; i < len(data)-2; i += 3 {
		firstBag := getLetters(data[i])
		secondBag := getLetters(data[i+1])
		thirdBag := getLetters(data[i+2])

		total += findCommonLetter(findCommonLetters(firstBag, secondBag), thirdBag) + 1
	}

	return total
}
