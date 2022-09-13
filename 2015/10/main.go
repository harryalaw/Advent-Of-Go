package main

import (
	"github.com/harryalaw/advent-of-go/util"
	"fmt"
	"io/ioutil"
	"strings"
)

func solve(part func() int) int {
	value := part()
	fmt.Println(value)
	return value
}

func main() {
	util.Time(func() { solve(Part1) }, "Part1")
	util.Time(func() { solve(Part2) }, "Part2")
}

func parseInput() string {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}
func LookAndSay(word string) string {
	var sb strings.Builder
	count := 1
	var prevChar = word[0]
	for i := 1; i < len(word); i++ {
		char := (word[i])
		if char != prevChar {
			sb.WriteRune(rune(count + 48))
			// sb.WriteRune(rune(count) + '0')
			sb.WriteByte(prevChar)
			prevChar = char
			count = 0
		}
		count++
	}
	sb.WriteRune(rune(count + 48))
	// sb.WriteRune(rune(count) + '0')
	sb.WriteByte(prevChar)
	return sb.String()
}

func Part1() int {
	word := parseInput()

	for i := 0; i < 40; i++ {
		word = LookAndSay(word)
	}

	return len(word)
}
func Part2() int {
	word := parseInput()

	for i := 0; i < 50; i++ {
		word = LookAndSay(word)
	}
	return len(word)
}

/*
Optimization:
	Investigate the cosmological decay of the sequence
	-> 92 possible elements that make each Part up
	-> If you can decompose the string into the elements
	-> Then you can figure out the next values
	-> and that gives you the lengths at each stage
	-> so each Part would be an array of ints[]
	and you would look up the next iteration from the ints[]
	-> then length is an easy computation

*/
