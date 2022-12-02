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

func parseInput(input string) []RPS {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	games := make([]RPS, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		rps := RPS{opponent: fields[0], player: fields[1]}
		games[i] = rps
	}

	return games
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

func Part1(games []RPS) int {
	totalScore := 0
	for _, game := range games {
		totalScore += scoreGame(game.opponent, game.player)
	}

	return totalScore
}

func Part2(games []RPS) int {
	totalScore := 0
	for _, game := range games {
		totalScore += scoreGameDecrypt(game.opponent, game.player)
	}

	return totalScore
}

func scoreGame(opponent, player string) int {
	score := 0
	switch player {
	case "X":
		score += 1
	case "Y":
		score += 2
	case "Z":
		score += 3
	}

	score += resolveGame(opponent, player)

	return score
}

func resolveGame(opponent, player string) int {
	if opponent == "A" && player == "X" ||
		opponent == "B" && player == "Y" ||
		opponent == "C" && player == "Z" {
		return 3
	} else if opponent == "A" && player == "Y" ||
		opponent == "B" && player == "Z" ||
		opponent == "C" && player == "X" {

		return 6
	}
	return 0
}

func scoreGameDecrypt(opponent, player string) int {
	score := 0

	switch player {
	case "X":
		score += 0
	case "Y":
		score += 3
	case "Z":
		score += 6
	}

	score += resolveGameDecrypt(opponent, player)

	return score
}

func resolveGameDecrypt(opponent, player string) int {
	switch player {
	case "X":
		if opponent == "A" {
			return 3
		} else if opponent == "B" {
			return 1
		} else if opponent == "C" {
			return 2
		}
	case "Y":
		if opponent == "A" {
			return 1
		} else if opponent == "B" {
			return 2
		} else if opponent == "C" {
			return 3
		}
	case "Z":
		if opponent == "A" {
			return 2
		} else if opponent == "B" {
			return 3
		} else if opponent == "C" {
			return 1
		}
	}
	return -1
}
