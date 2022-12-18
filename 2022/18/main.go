package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

type Cube struct {
	x int
	y int
	z int
}

type ThreeSpace map[int](map[int](map[int]bool))

func (ts *ThreeSpace) contains(cube Cube) bool {

	return (*ts)[cube.x][cube.y][cube.z]
}

func (ts *ThreeSpace) set(x, y, z int, val bool) {
	if _, ok := (*ts)[x]; !ok {
		(*ts)[x] = map[int](map[int]bool){}
	}
	if _, ok := (*ts)[x][y]; !ok {
		(*ts)[x][y] = map[int]bool{}
	}
	(*ts)[x][y][z] = true
}

func parseInput(input string) ([]Cube, *ThreeSpace) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	space := ThreeSpace{}
	cubes := make([]Cube, len(lines))
	for i, line := range lines {
		dims := strings.Split(line, ",")

		x := util.Atoi(dims[0])
		y := util.Atoi(dims[1])
		z := util.Atoi(dims[2])

		space.set(x, y, z, true)

		cubes[i] = Cube{x, y, z}
	}
	return cubes, &space
}

func main() {
	util.Time(doPart1, "Part 1")
	util.Time(doPart2, "Part 2")
}

func doPart1() {
	cubes, space := parseInput(input)
	fmt.Println("Part 1: ", Part1(cubes, space))
}

func doPart2() {
	cubes, space := parseInput(input)
	fmt.Println("Part 2: ", Part2(cubes, space))
}

func (ts *ThreeSpace) isPresent(x, y, z int) int {
	val, ok := (*ts)[x][y][z]
	if ok && val {
		return 1
	}
	return 0
}

func countNeighbours(cube Cube, space *ThreeSpace) int {
	nbrCount := 0

	nbrCount += space.isPresent(cube.x, cube.y, cube.z+1)
	nbrCount += space.isPresent(cube.x, cube.y, cube.z-1)
	nbrCount += space.isPresent(cube.x, cube.y+1, cube.z)
	nbrCount += space.isPresent(cube.x, cube.y-1, cube.z)
	nbrCount += space.isPresent(cube.x+1, cube.y, cube.z)
	nbrCount += space.isPresent(cube.x-1, cube.y, cube.z)

	return nbrCount
}

func Part1(cubes []Cube, space *ThreeSpace) int {
	count := 0
	for _, cube := range cubes {
		nbrCount := countNeighbours(cube, space)
		count += 6 - nbrCount
	}

	return count
}

func Part2(cubes []Cube, space *ThreeSpace) int {
	// lets instead make a 3d visited of cubes
	// then starting at [0,0,0] we'll flood fill the outside and count external faces
	visited := ThreeSpace{}

	toVisit := make([]Cube, 0)
	toVisit = append(toVisit, Cube{0, 0, 0})

	externalFaces := 0
	for len(toVisit) != 0 {
		next := toVisit[0]
		toVisit = toVisit[1:]
		externalFaces += countNeighbours(next, space)

		toVisit = append(toVisit, validNeighbours(next, &visited, space)...)
	}

	return externalFaces
}

func validNeighbours(next Cube, visited *ThreeSpace, space *ThreeSpace) []Cube {
	out := make([]Cube, 0)

	for x := next.x - 1; x <= next.x+1; x++ {
		if x < -1 || x > 22 {
			continue
		}
		for y := next.y - 1; y <= next.y+1; y++ {
			if y < -1 || y > 22 {
				continue
			}
			for z := next.z - 1; z <= next.z+1; z++ {
				if z < -1 || z > 22 {
					continue
				}
				if x != next.x && y != next.y || x != next.x && z != next.z || y != next.y && z != next.z {
					continue
				}
				if visited := (*visited)[x][y][z]; visited {
					continue
				}
				if containedInRock := space.isPresent(x, y, z) == 1; containedInRock {
					continue
				}
				visited.set(x, y, z, true)
				out = append(out, Cube{x, y, z})
			}
		}
	}

	return out
}
