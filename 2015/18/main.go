package main

import (
	"github.com/harryalaw/advent-of-go/util"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	doPart1("test.txt", 4, 6, 6)
	util.Time(part1, "Part1")
	doPart2("test.txt", 5, 6, 6)
	util.Time(part2, "Part2")
}

type coord struct {
	x int
	y int
}

func parseInput(filename string) map[coord]int {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	output := make(map[coord]int)
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				output[coord{x: x, y: y}] = 1
			}
		}
	}

	return output
}

func part1() {
	doPart1("data.txt", 100, 100, 100)
}

func part2() {
	doPart2("data.txt", 100, 100, 100)
}

func doPart1(fileName string, times, width, height int) {
	livingCells := parseInput(fileName)

	for i := 0; i < times; i++ {
		// printCells(livingCells, width, height)
		livingCells = playGameOfLife(livingCells, width, height)
		// fmt.Println(len(livingCells))
	}
	// printCells(livingCells, width, height)
	fmt.Println(len(livingCells))
}

func doPart2(fileName string, times, width, height int) {
	livingCells := parseInput(fileName)
	turnOnCorners(&livingCells, width, height)

	for i := 0; i < times; i++ {
		livingCells = playGameOfLife(livingCells, width, height)
		turnOnCorners(&livingCells, width, height)
		// printCells(livingCells, width, height)
	}
	fmt.Println(len(livingCells))
}

func playGameOfLife(cells map[coord]int, width, height int) map[coord]int {
	nextCells := make(map[coord]int)

	neighbours := make(map[coord]int)
	for cell := range cells {
		nbrs := getNeighbours(cell, width, height)
		aliveCounts := getLivingCount(cells, nbrs)
		if aliveCounts == 2 || aliveCounts == 3 {
			nextCells[cell] = 1
		}

		for _, nbr := range nbrs {
			neighbours[nbr] = 0
		}
	}

	for cell := range neighbours {
		_, isPresent := cells[cell]
		if !isPresent {
			nbrs := getNeighbours(cell, width, height)
			aliveCounts := getLivingCount(cells, nbrs)
			if aliveCounts == 3 {
				nextCells[cell] = 1
			}
		}
	}
	return nextCells
}

func getNeighbours(point coord, width, height int) []coord {
	outArray := make([]coord, 0)
	for i := point.x - 1; i <= point.x+1; i++ {
		for j := point.y - 1; j <= point.y+1; j++ {

			if i < 0 ||
				j < 0 ||
				i >= width ||
				j >= height ||
				(i == point.x && j == point.y) {
				continue
			}
			outArray = append(outArray, coord{x: i, y: j})
		}
	}
	return outArray
}

func getLivingCount(cellMap map[coord]int, cells []coord) int {
	count := 0

	for _, cell := range cells {
		_, isPresent := cellMap[cell]
		if isPresent {
			count++
		}
	}
	return count
}

func turnOnCorners(cellMap *map[coord]int, width, height int) {
	topLeft := coord{x: 0, y: 0}
	(*cellMap)[topLeft] = 1

	topRight := coord{x: width - 1, y: 0}
	(*cellMap)[topRight] = 1

	bottomLeft := coord{x: 0, y: height - 1}
	(*cellMap)[bottomLeft] = 1

	bottomRight := coord{x: width - 1, y: height - 1}
	(*cellMap)[bottomRight] = 1
}

func printCells(cellMap map[coord]int, width, height int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			tempCoord := coord{x: j, y: i}
			_, isPresent := cellMap[tempCoord]

			if isPresent {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}
