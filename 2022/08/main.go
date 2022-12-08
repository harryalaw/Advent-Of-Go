package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

type Grid [][]int

func parseInput(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]int, len(lines))
	for i, line := range lines {
		gridLine := make([]int, len(line))
		for j, char := range line {
			gridLine[j] = runeToNum(char)
		}
		grid[i] = gridLine
	}
	return grid
}

func runeToNum(r rune) int {
	return int(r - 48)
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

func Part1(grid [][]int) int {
	visible := make([][]bool, len(grid))
	for i := range grid {
		line := make([]bool, len(grid[i]))
		visible[i] = line
	}

	visibleFromOutside(grid, &visible)

	total := 0
	for _, line := range visible {
		for _, val := range line {
			if val {
				total++
			}
		}
	}

	// printVisible(visible)

	return total
}

func visibleFromOutside(grid [][]int, visible *[][]bool) {
	height := len(grid)
	width := len(grid[0])

	for i := 0; i < height; i++ {
		highest := -1
		for j := 0; j < width; j++ {
			tree := grid[i][j]
			if tree > highest {
				(*visible)[i][j] = true
				highest = tree
			}
		}
	}
	for j := 0; j < width; j++ {
		highest := -1
		for i := 0; i < height; i++ {
			tree := grid[i][j]
			if tree > highest {
				(*visible)[i][j] = true
				highest = tree
			}
		}
	}

	for i := 0; i < height; i++ {
		highest := -1
		for j := width - 1; j >= 0; j-- {
			tree := grid[i][j]
			if tree > highest {
				(*visible)[i][j] = true
				highest = tree
			}
		}
	}

	for j := 0; j < width; j++ {
		highest := -1
		for i := height - 1; i >= 0; i-- {
			tree := grid[i][j]
			if tree > highest {
				(*visible)[i][j] = true
				highest = tree
			}
		}
	}
}

func Part2(grid [][]int) int {

	scenicScore := 0

	for i, line := range grid {
		for j := range line {
			tempScore := score(i, j, grid)
			if tempScore > scenicScore {
				scenicScore = tempScore
			}
		}
	}
	return scenicScore
}

func score(i, j int, grid [][]int) int {
	height := len(grid)
	width := len(grid[0])

	leftCount := 0
	rightCount := 0
	topCount := 0
	bottomCount := 0

	startHeight := grid[i][j]

	for idx := j + 1; idx < width; idx++ {
		tree := grid[i][idx]
		rightCount++
		if tree >= startHeight {
			break
		}
	}

	for idx := j - 1; idx >= 0; idx-- {
		tree := grid[i][idx]
		leftCount++
		if tree >= startHeight {
			break
		}
	}

	for idx := i + 1; idx < height; idx++ {
		tree := grid[idx][j]
		bottomCount++
		if tree >= startHeight {
			break
		}
	}

	for idx := i - 1; idx >= 0; idx-- {
		tree := grid[idx][j]
		topCount++
		if tree >= startHeight {
			break
		}
	}

	return leftCount * rightCount * topCount * bottomCount
}

func printVisible(visible [][]bool) {
	for _, line := range visible {
		fmt.Printf("%+v\n", line)
	}
}
