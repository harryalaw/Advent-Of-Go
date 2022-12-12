package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

func parseInput(input string) ([][]int, Step, Step) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	jungleMap := make([][]int, len(lines))
	var start Step
	var end Step
	for i, line := range lines {
		jungleRow := make([]int, len(line))
		for j, char := range line {
			if char == 'S' {
				jungleRow[j] = 0
				start = Step{row: i, col: j, moves: 0}
			} else if char == 'E' {
				jungleRow[j] = 25
				end = Step{row: i, col: j, moves: 1000000}
			} else {
				jungleRow[j] = int(char - 'a')
			}
		}
		jungleMap[i] = jungleRow
	}

	return jungleMap, start, end
}

func main() {
	util.Time(doPart1, "Part 1")
	util.Time(doPart2, "Part 2")
}

func doPart1() {
	data, start, end := parseInput(input)
	fmt.Println("Part 1: ", Part1(data, start, end))
}

func doPart2() {
	data, _, end := parseInput(input)
	fmt.Println("Part 2: ", Part2(data, end))
}

type Step struct {
	row   int
	col   int
	moves int
}

func (s *Step) hash() int {
	return s.row + s.col*100000
}

func (s *Step) getNeighbours(jungleMap [][]int) []Step {
	currentHeight := jungleMap[s.row][s.col]

	out := make([]Step, 0)
	height := len(jungleMap)
	width := len(jungleMap[0])

	if s.row != 0 && jungleMap[s.row-1][s.col] <= currentHeight+1 {
		out = append(out, Step{row: s.row - 1, col: s.col, moves: s.moves + 1})
	}
	if s.row != height-1 && jungleMap[s.row+1][s.col] <= currentHeight+1 {
		out = append(out, Step{row: s.row + 1, col: s.col, moves: s.moves + 1})
	}
	if s.col != 0 && jungleMap[s.row][s.col-1] <= currentHeight+1 {
		out = append(out, Step{row: s.row, col: s.col - 1, moves: s.moves + 1})
	}
	if s.col != width-1 && jungleMap[s.row][s.col+1] <= currentHeight+1 {
		out = append(out, Step{row: s.row, col: s.col + 1, moves: s.moves + 1})
	}
	return out
}

func (s *Step) getNeighbours2(jungleMap [][]int) []Step {
	currentHeight := jungleMap[s.row][s.col]

	out := make([]Step, 0)
	height := len(jungleMap)
	width := len(jungleMap[0])

	if s.row != 0 && jungleMap[s.row-1][s.col] >= currentHeight-1 {
		out = append(out, Step{row: s.row - 1, col: s.col, moves: s.moves + 1})
	}
	if s.row != height-1 && jungleMap[s.row+1][s.col] >= currentHeight-1 {
		out = append(out, Step{row: s.row + 1, col: s.col, moves: s.moves + 1})
	}
	if s.col != 0 && jungleMap[s.row][s.col-1] >= currentHeight-1 {
		out = append(out, Step{row: s.row, col: s.col - 1, moves: s.moves + 1})
	}
	if s.col != width-1 && jungleMap[s.row][s.col+1] >= currentHeight-1 {
		out = append(out, Step{row: s.row, col: s.col + 1, moves: s.moves + 1})
	}
	return out
}

func Part1(data [][]int, start, end Step) int {
	visited := make(map[int]Step)

	toVisit := []Step{start}
	for len(toVisit) != 0 {
		minDistance := 1<<31 - 1
		var minNode Step
		var nodeIdx int
		for i, node := range toVisit {
			if node.moves < minDistance {
				minDistance = node.moves
				minNode = node
				nodeIdx = i
			}
		}
		toVisit = append(toVisit[:nodeIdx], toVisit[nodeIdx+1:]...)
		prev, ok := visited[minNode.hash()]
		if !ok {
			visited[minNode.hash()] = minNode
		} else {
			if prev.moves <= minNode.moves {
				continue
			}
		}

		for _, nbr := range minNode.getNeighbours(data) {
			prev, ok := visited[nbr.hash()]
			if ok && prev.moves <= nbr.moves {
				continue
			}
			toVisit = append(toVisit, nbr)
		}
	}

	return visited[end.hash()].moves
}

func Part2(data [][]int, end Step) int {
	visited := make(map[int]Step)

	end.moves = 0

	toVisit := []Step{end}
	for len(toVisit) != 0 {
		minDistance := 1<<31 - 1
		var minNode Step
		var nodeIdx int
		for i, node := range toVisit {
			if node.moves < minDistance {
				minDistance = node.moves
				minNode = node
				nodeIdx = i
			}
		}
		toVisit = append(toVisit[:nodeIdx], toVisit[nodeIdx+1:]...)
		prev, ok := visited[minNode.hash()]
		if !ok {
			visited[minNode.hash()] = minNode
		} else {
			if prev.moves <= minNode.moves {
				continue
			}
		}
		if data[minNode.row][minNode.col] == 0 {
			return minNode.moves
		}

		for _, nbr := range minNode.getNeighbours2(data) {
			prev, ok := visited[nbr.hash()]
			if ok && prev.moves <= nbr.moves {
				continue
			}
			toVisit = append(toVisit, nbr)
		}
	}
	return -1
}
