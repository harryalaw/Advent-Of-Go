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
	commonParts := make([]int, len(data))

	for i, line := range data {
		length := len(line)
		firstHalf := line[0 : length/2]
		secondHalf := line[length/2:]
		firstLetters := getLetters(firstHalf)
		secondLetters := getLetters(secondHalf)
		commonParts[i] = findCommonLetter(firstLetters, secondLetters)
	}

	return util.IntSum(commonParts...)
}

func findCommonLetter(a, b []bool) int {
	if len(a) != len(b) {
		panic("lengths aren't equal")
	}
	for i, value := range a {
		if value && b[i] {
			return i
		}
	}

	panic("No match")
}

func findCommonLetters(a, b []bool) []bool {
	if len(a) != len(b) {
		panic("lengths aren't equal")
	}
	out := make([]bool, 53)
	for i, value := range a {
		if value && b[i] {
			out[i] = true
		}
	}
	return out
}

func getLetters(item string) []bool {
	out := make([]bool, 53)
	valueMap := map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
		'i': 9,
		'j': 10,
		'k': 11,
		'l': 12,
		'm': 13,
		'n': 14,
		'o': 15,
		'p': 16,
		'q': 17,
		'r': 18,
		's': 19,
		't': 20,
		'u': 21,
		'v': 22,
		'w': 23,
		'x': 24,
		'y': 25,
		'z': 26,
		'A': 27,
		'B': 28,
		'C': 29,
		'D': 30,
		'E': 31,
		'F': 32,
		'G': 33,
		'H': 34,
		'I': 35,
		'J': 36,
		'K': 37,
		'L': 38,
		'M': 39,
		'N': 40,
		'O': 41,
		'P': 42,
		'Q': 43,
		'R': 44,
		'S': 45,
		'T': 46,
		'U': 47,
		'V': 48,
		'W': 49,
		'X': 50,
		'Y': 51,
		'Z': 52,
	}
	for _, char := range item {

		out[valueMap[char]] = true
	}

	return out
}

func Part2(data []string) int {
	commonParts := make([]int, len(data))
	for i := 0; i < len(data)-2; i += 3 {
		firstBag := getLetters(data[i])
		secondBag := getLetters(data[i+1])
		thirdBag := getLetters(data[i+2])

		commonParts[i] = findCommonLetter(findCommonLetters(firstBag, secondBag), thirdBag)
	}

	return util.IntSum(commonParts...)
}
