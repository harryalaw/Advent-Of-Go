package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

type RPS struct {
	opponent string
	player   string
}

func parseInput(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	return lines

	// games := make([]RPS, len(lines))
	//
	// for i, line := range lines {
	// 	fields := strings.Fields(line)
	// 	rps := RPS{opponent: fields[0], player: fields[1]}
	// 	games[i] = rps
	// }
	//
	// return games
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

func Part1(games []string) int {
	totalScore := 0
	for _, game := range games {
		totalScore += scoreGame(game)
	}

	return totalScore
}

func Part2(games []string) int {
	totalScore := 0
	for _, game := range games {
		totalScore += scoreGameDecrypt(game)
	}

	return totalScore
}

func scoreGame(game string) int {
	switch game {
	//loss
	case "B X":
		return 1
	case "C Y":
		return 2
	case "A Z":
		return 3
	//draws
	case "A X":
		return 4
	case "B Y":
		return 5
	case "C Z":
		return 6
	//wins
	case "C X":
		return 7
	case "A Y":
		return 8
	case "B Z":
		return 9
	default:
		return -1
	}
}

func scoreGameDecrypt(game string) int {
	switch game {
	//loss
	case "B X":
		return 1
	case "C X":
		return 2
	case "A X":
		return 3
	//draws
	case "A Y":
		return 4
	case "B Y":
		return 5
	case "C Y":
		return 6
	//wins
	case "C Z":
		return 7
	case "A Z":
		return 8
	case "B Z":
		return 9
	default:
		return -1
	}
}
