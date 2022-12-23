package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

type Coord struct {
	x int
	y int
}

func (c *Coord) add(o *Coord) Coord {
	return Coord{
		x: c.x + o.x,
		y: c.y + o.y,
	}
}

func (c *Coord) equals(o *Coord) bool {
	return c.x == o.x && c.y == o.y
}

func (c *Coord) hash() int {
	return c.x*100_000 + c.y
}

type ElfPosition map[int]Coord

func parseInput(input string) ElfPosition {
	elfPositions := ElfPosition{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	elfCount := 0
	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				elfPositions[elfCount] = Coord{x: col, y: row}
				elfCount++
			}
		}
	}
	return elfPositions
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

func (ep *ElfPosition) print() {
	grid := make([][]rune, 0)

	gridMin := -3
	gridMax := 11
	for j := gridMin; j <= gridMax; j++ {
		line := make([]rune, 0)
		for i := gridMin; i < gridMax; i++ {
			line = append(line, '.')
		}
		grid = append(grid, line)
	}
	for _, coord := range *ep {
		grid[coord.y+3][coord.x+3] = '#'
	}

	for _, row := range grid {
		for _, char := range row {
			fmt.Printf("%c", char)
		}
		fmt.Println()
	}
	fmt.Println("----")
}

func Part1(elves ElfPosition) int {
	for i := 0; i < 10; i++ {
		// elves.print()
		proposedPositions := proposePositions(elves, i)
		elves, _ = moveElves(elves, proposedPositions)
	}

	minX := elves[0].x
	minY := elves[0].y
	maxX := elves[0].x
	maxY := elves[0].y

	for _, coord := range elves {
		if coord.x < minX {
			minX = coord.x
		}
		if coord.y < minY {
			minY = coord.y
		}
		if coord.x > maxX {
			maxX = coord.x
		}
		if coord.y > maxY {
			maxY = coord.y
		}
	}
	rectangleArea := (maxX - minX + 1) * (maxY - minY + 1)
	return rectangleArea - len(elves)
}

var north = &Coord{0, -1}
var northEast = &Coord{1, -1}
var east = &Coord{1, 0}
var southEast = &Coord{1, 1}
var south = &Coord{0, 1}
var southWest = &Coord{-1, 1}
var west = &Coord{-1, 0}
var northWest = &Coord{-1, -1}

func getDirections(round int) [][]*Coord {
	order := round % 4
	switch order {
	case 0:
		return [][]*Coord{
			{north, northEast, northWest},
			{south, southEast, southWest},
			{west, southWest, northWest},
			{east, southEast, northEast},
		}
	case 1:
		return [][]*Coord{
			{south, southEast, southWest},
			{west, southWest, northWest},
			{east, southEast, northEast},
			{north, northEast, northWest},
		}

	case 2:
		return [][]*Coord{
			{west, southWest, northWest},
			{east, southEast, northEast},
			{north, northEast, northWest},
			{south, southEast, southWest},
		}
	default:
		return [][]*Coord{
			{east, southEast, northEast},
			{north, northEast, northWest},
			{south, southEast, southWest},
			{west, southWest, northWest},
		}
	}
}

type ProposedPosition map[Coord][]int

func proposePositions(elves ElfPosition, round int) ProposedPosition {
	occupiedSpaces := map[int]bool{}
	for _, coord := range elves {
		occupiedSpaces[coord.hash()] = true
	}
	directions := getDirections(round)

	proposedPositions := ProposedPosition{}

	for elfIdx, coord := range elves {
		hasNeighbours := getNeighbours(coord, &occupiedSpaces)

		if hasNeighbours {
			for _, direction := range directions {
				newPos1 := coord.add(direction[0])
				newPos2 := coord.add(direction[1])
				newPos3 := coord.add(direction[2])
				emptySpace := true
				if occupiedSpaces[newPos1.hash()] || occupiedSpaces[newPos2.hash()] || occupiedSpaces[newPos3.hash()] {
					emptySpace = false
				}
				if !emptySpace {
					continue
				}

				prev, exists := proposedPositions[newPos1]
				if !exists {
					prev = make([]int, 0)
				}
				prev = append(prev, elfIdx)
				proposedPositions[newPos1] = prev
				break
			}
		}
	}
	return proposedPositions
}

func getNeighbours(coord Coord, occupiedSpaces *map[int]bool) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			newCoord := Coord{coord.x + i, coord.y + j}
			if _, ok := (*occupiedSpaces)[newCoord.hash()]; ok {
				return true
			}
		}
	}
	return false
}

func moveElves(elves ElfPosition, moves ProposedPosition) (ElfPosition, int) {
	moveCounts := 0
	for coord, movingElves := range moves {
		if len(movingElves) == 1 {
			elves[movingElves[0]] = coord
			moveCounts++
		}
	}

	return elves, moveCounts
}

func Part2(elves ElfPosition) int {
	moveCounts := 1
	iter := 0
	for moveCounts != 0 {
		proposedPositions := proposePositions(elves, iter)
		elves, moveCounts = moveElves(elves, proposedPositions)
		iter++
	}

	return iter
}
