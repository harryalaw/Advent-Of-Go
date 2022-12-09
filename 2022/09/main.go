package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

type Move struct {
	direction *Coord
	amount    int
}

type Coord struct {
	x int
	y int
}

func (c *Coord) Hash() int {
	return c.x + 100000*c.y
}

func (c *Coord) Add(o Coord) Coord {
	newX := c.x + o.x
	newY := c.y + o.y
	return Coord{x: newX, y: newY}
}

func (c *Coord) Sub(o Coord) Coord {
	newX := c.x - o.x
	newY := c.y - o.y
	return Coord{x: newX, y: newY}
}

func (c *Coord) Euclid() int {
	return util.IntAbs(c.x*c.x) + util.IntAbs(c.y*c.y)
}

func (c *Coord) Normalize() {
	xSign := 1
	if c.x < 0 {
		xSign = -1
	}
	ySign := 1
	if c.y < 0 {
		ySign = -1
	}

	if c.x != 0 {
		c.x = 1 * xSign
	}
	if c.y != 0 {
		c.y = 1 * ySign
	}
}

var up = Coord{x: 0, y: 1}
var down = Coord{x: 0, y: -1}
var left = Coord{x: -1, y: 0}
var right = Coord{x: 1, y: 0}

func directionToCoord(dir string) *Coord {
	switch dir {
	case "U":
		return &up
	case "D":
		return &down
	case "L":
		return &left
	case "R":
		return &right
	}
	panic("No such direction")
}

func parseInput(input string) []Move {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	moves := make([]Move, len(lines))

	for i, move := range lines {
		parts := strings.Fields(move)
		moves[i] = Move{
			direction: directionToCoord(parts[0]),
			amount:    util.Atoi(parts[1]),
		}
	}

	return moves
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

func Part1(moves []Move) int {
	visited := map[int]struct{}{}
	head := Coord{0, 0}
	tail := Coord{0, 0}
	visited[tail.Hash()] = struct{}{}

	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			head = head.Add(*move.direction)
			tail = advanceTail(head, tail)
			visited[tail.Hash()] = struct{}{}
		}
	}

	return len(visited)
}

func Part2(moves []Move) int {
	visited := map[int]struct{}{}
	snake := [10]Coord{}
	for i := range snake {
		snake[i] = Coord{0, 0}
	}
	tail := snake[9]
	visited[tail.Hash()] = struct{}{}

	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			snake[0] = snake[0].Add(*move.direction)
			for i := 1; i < len(snake); i++ {
				snake[i] = advanceTail(snake[i-1], snake[i])
			}
			visited[snake[9].Hash()] = struct{}{}
		}
	}

	return len(visited)
}

func advanceTail(head, tail Coord) Coord {
	difference := head.Sub(tail)
	mag := difference.Euclid()
	if mag < 3 {
		return tail
	}
	difference.Normalize()
	// H . T
	if mag == 4 {
		return tail.Add(difference)
	}
	// H . .
	// . . T
	// actually want T to move diagonally -1,1
	// H .
	// . .
	// . T
	return tail.Add(difference)
}
