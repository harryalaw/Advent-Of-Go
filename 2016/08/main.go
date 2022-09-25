package main

import (
	_ "embed"
	"fmt"
	"github.com/harryalaw/advent-of-go/util"
	"strings"
)

func main() {
	util.Time(doPart1, "Part1")
}

//go:embed "data.txt"
var input string

const HEIGHT = 6
const WIDTH = 50

type lightGrid [HEIGHT][WIDTH]bool

func parseLine(line string, grid lightGrid) lightGrid {
	words := strings.Fields(line)
	switch words[0] {
	case "rect":
		numbers := strings.Split(words[1], "x")
		w := util.Atoi(string(numbers[0]))
		h := util.Atoi(string(numbers[1]))
		grid = rect(w, h, grid)
	case "rotate":
		index := util.Atoi(strings.Split(words[2], "=")[1])
		mag := util.Atoi(words[4])
		if "row" == words[1] {
			grid = rotateRow(index, mag, grid)
		} else {
			grid = rotateCol(index, mag, grid)
		}
	}
	return grid
}

func doPart1() {
	grid := initGrid()

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		grid = parseLine(line, grid)
	}

	lights := 0;
	for _, line := range grid {
		for _, char := range line {
			if char {
				lights++;
				fmt.Print("⬜️")
			} else {
				fmt.Print("⬛️")
			}
		}
		fmt.Println()
	}
	fmt.Println("Number of lights: ", lights)
}

func rect(w, h int, grid lightGrid) lightGrid {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			grid[i][j] = true
		}
	}
	return grid
}

func rotateCol(index, mag int, grid lightGrid) lightGrid {
	prevGrid := grid
	for i := 0; i < HEIGHT; i++ {
		grid[(i+mag)%HEIGHT][index] = prevGrid[i][index]
	}
	return grid
}

func rotateRow(index, mag int, grid lightGrid) lightGrid {
	prevGrid := grid
	for i := 0; i < WIDTH; i++ {
		grid[index][(i+mag)%WIDTH] = prevGrid[index][i]
	}
	return grid
}

func initGrid() lightGrid {
	grid := lightGrid{}
	return grid
}
