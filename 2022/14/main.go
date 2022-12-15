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

func (c *Coord) Hash() int {
	return c.x + c.y*100000
}

func unhash(hash int) Coord {
	x := hash % 100000
	y := (hash - x) / 100000
	return Coord{x, y}
}

func (c *Coord) Direction(o Coord) Coord {
	if c.x != o.x {
		if o.x > c.x {
			return Coord{1, 0}
		} else {
			return Coord{-1, 0}
		}
	} else {
		if o.y > c.y {
			return Coord{0, 1}
		} else {
			return Coord{0, -1}
		}
	}
}

// 498,4 -> 498,6 -> 496,6
func parseInput(input string) (map[int]bool, Coord) {
	rocks := strings.Split(strings.TrimSpace(input), "\n")

	cave := map[int]bool{}

	for _, rock := range rocks {
		parts := strings.Split(rock, " -> ")
		sections := make([]Coord, len(parts))
		for i, part := range parts {
			coords := strings.Split(part, ",")
			x := util.Atoi(coords[0])
			y := util.Atoi(coords[1])
			sections[i] = Coord{x, y}
			cave[sections[i].Hash()] = true
		}

		// now fill in the lines
		previous := sections[0]

		for i := 1; i < len(sections); i++ {
			var direction int
			curr := sections[i]
			if curr.x != previous.x {
				distance := curr.x - previous.x
				if distance > 0 {
					direction = 1
				} else {
					direction = -1
				}
				distance = util.IntAbs(distance)
				for j := 0; j < distance; j++ {
					tempCoord := Coord{previous.x + j*direction, previous.y}
					cave[tempCoord.Hash()] = true
				}
			} else if curr.y != previous.y {
				distance := curr.y - previous.y
				if distance > 0 {
					direction = 1
				} else {
					direction = -1
				}
				distance = util.IntAbs(distance)
				for j := 0; j <= distance; j++ {
					tempCoord := Coord{previous.x, previous.y + j*direction}
					cave[tempCoord.Hash()] = true
				}
			}
			previous = curr
		}
	}

	// need to figure out the leftmost thing with nothing under it
	// and the rightmost thing with nothing under it
	bottom := Coord{500, 0}
	for hash, _ := range cave {
		coord := unhash(hash)
		if coord.Hash() != hash {
			panic("Oops")
		}
		if coord.y > bottom.y {
			bottom = coord
		}
	}

	return cave, bottom
}

func printCave(cave map[int]bool) {
	for y := 0; y < 10; y++ {
		for x := 494; x < 504; x++ {

			tempCoord := Coord{x, y}
			val, ok := cave[tempCoord.Hash()]
			if ok && val {
				fmt.Print("#")
			} else if ok {
				fmt.Print("o")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	util.Time(doPart1, "Part 1")
	util.Time(doPart2, "Part 2")
}

func doPart1() {
	cave, bottom := parseInput(input)
	fmt.Println("Part 1: ", Part1(cave, bottom))
}

func doPart2() {
	data, bottom := parseInput(input)
	fmt.Println("Part 2: ", Part2(data, bottom))
}

// would be good to know bounds of the cave
func Part1(cave map[int]bool, bottom Coord) int {
	count := 0
	sand := Coord{500, 0}

	for {
		belowPos := Coord{sand.x, sand.y + 1}
		if !cave[belowPos.Hash()] {
			sand = belowPos
			if sand.y > bottom.y {
				return count
			}
			continue
		}

		leftBelow := Coord{sand.x - 1, sand.y + 1}
		if !cave[leftBelow.Hash()] {
			sand = leftBelow
			continue
		}

		rightBelow := Coord{sand.x + 1, sand.y + 1}
		if !cave[rightBelow.Hash()] {
			sand = rightBelow
			continue
		}

		cave[sand.Hash()] = true
		count++
		sand = Coord{500, 0}
	}
}

func Part2(cave map[int]bool, bottom Coord) int {

	count := 0
	sand := Coord{500, 0}
	floor := 2 + bottom.y
	// Floor is lowest sand can go to
	for {
		belowPos := Coord{sand.x, sand.y + 1}
		if sand.y+1 != floor && !cave[belowPos.Hash()] {
			sand = belowPos
			continue
		}

		leftBelow := Coord{sand.x - 1, sand.y + 1}
		if sand.y+1 != floor && !cave[leftBelow.Hash()] {
			sand = leftBelow
			continue
		}

		rightBelow := Coord{sand.x + 1, sand.y + 1}
		if sand.y+1 != floor && !cave[rightBelow.Hash()] {
			sand = rightBelow
			continue
		}

		cave[sand.Hash()] = true
		count++
		if sand.Hash() == 500 {
			return count
		}
		sand = Coord{500, 0}
	}
	return -1
}
