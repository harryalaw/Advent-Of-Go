package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type instruction struct {
	label string
	x1    int
	x2    int
	y1    int
	y2    int
}

func parseInput() []instruction {
	content, err := ioutil.ReadFile("data.txt")

	if err != nil {
		panic(err)
	}

	lines := splitNewLines(string(content))

	out := make([]instruction, len(lines))

	for index, line := range lines {
		out[index] = stringToInstruction(line)
	}
	return out
}

func stringToInstruction(line string) instruction {
	var instruction = new(instruction)
	lineRegex := regexp.MustCompile(`(?s)(toggle|turn on|turn off) (\d*),(\d*) through (\d*),(\d*)`)

	matches := lineRegex.FindAllStringSubmatch(line, -1)
	values := matches[0][1:]

	instruction.label = values[0]
	instruction.x1, _ = strconv.Atoi(values[1])
	instruction.y1, _ = strconv.Atoi(values[2])
	instruction.x2, _ = strconv.Atoi(values[3])
	instruction.y2, _ = strconv.Atoi(values[4])
	return *instruction
}

func splitNewLines(content string) []string {
	return strings.Split(content, "\n")
}

func solve(part func([]instruction) int) int {
	count := part(parseInput())
	println(count)
	return count
}

func part1(instructions []instruction) int {
	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}

	for _, instruction := range instructions {
		for i := instruction.x1; i <= instruction.x2; i++ {
			for j := instruction.y1; j <= instruction.y2; j++ {
				switch instruction.label {
				case "turn on":
					grid[i][j] = 1
				case "turn off":
					grid[i][j] = 0
				case "toggle":
					if grid[i][j] == 1 {
						grid[i][j] = 0
					} else {
						grid[i][j] = 1
					}
				}
			}
		}
	}

	return countGrid(&grid)
}

func part2(instructions []instruction) int {
	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}

	for _, instruction := range instructions {
		minX := min(instruction.x1, instruction.x2)
		minY := min(instruction.y1, instruction.y2)
		maxX := max(instruction.x1, instruction.x2)
		maxY := max(instruction.y1, instruction.y2)
		for i := minX; i <= maxX; i++ {
			for j := minY; j <= maxY; j++ {
				switch instruction.label {
				case "turn on":
					grid[i][j]++
				case "turn off":
					grid[i][j] = max(grid[i][j]-1, 0)
				case "toggle":
					grid[i][j] += 2
				}
			}
		}

	}
	return countGrid(&grid)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countGrid(grid *[][]int) int {
	count := 0
	for i := 0; i < len(*grid); i++ {
		for j := 0; j < len((*grid)[i]); j++ {
			count += (*grid)[i][j]
		}
	}
	return count
}

func main() {
	start := time.Now()
	solve(part1)
	elapsed := time.Since(start)
	fmt.Printf("Part1 took %s\n", elapsed)

	start = time.Now()
	solve(part2)
	elapsed = time.Since(start)
	fmt.Printf("Part2 took %s\n", elapsed)
}

func test() {
	// fmt.Println(part1(instructions) == 500)
	// fmt.Println()
	test1()
	test2()
	test3()
	test4()
	test5()
	actual()
}

func test1() {
	inst := new(instruction)
	inst.label = "turn on"
	inst.x1 = 0
	inst.y1 = 0
	inst.x2 = 999
	inst.y2 = 999
	list := make([]instruction, 1)
	list[0] = *inst

	println(1000*1000 == part1(list))
}
func test2() {
	inst := new(instruction)
	inst.label = "toggle"
	inst.x1 = 0
	inst.y1 = 0
	inst.x2 = 999
	inst.y2 = 0
	list := make([]instruction, 1)
	list[0] = *inst

	println(1000 == part1(list))
}

func test3() {
	inst := new(instruction)
	inst.label = "turn on"
	inst.x1 = 0
	inst.y1 = 0
	inst.x2 = 999
	inst.y2 = 999

	inst2 := new(instruction)
	inst2.label = "turn off"
	inst2.x1 = 499
	inst2.y1 = 499
	inst2.x2 = 500
	inst2.y2 = 500

	list := make([]instruction, 2)
	list[0] = *inst
	list[1] = *inst2

	println(1000*1000-4 == part1(list))
}

func test4() {
	inst := new(instruction)
	inst.label = "turn on"
	inst.x1 = 0
	inst.y1 = 0
	inst.x2 = 0
	inst.y2 = 0
	list := make([]instruction, 1)
	list[0] = *inst

	println(1 == part2(list))
}
func test5() {
	list := make([]instruction, 1)
	list[0] = stringToInstruction("toggle 0,0 through 999,999")

	println(2000000 == part2(list))
}

func actual() {
	println(solve(part1) == 377891)
	println(solve(part2) == 14110788)
}
