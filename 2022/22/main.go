package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

type Maze []Row

type Row struct {
	start int
	row   string
}

func parseInput1(input string) (Maze, []string) {
	parts := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	mazeLines := strings.Split(strings.TrimRight(parts[0], "\n"), "\n")

	maze := make(Maze, len(mazeLines))

	for i, line := range mazeLines {
		start := 0
		for i := 0; i < len(line); i++ {
			if line[i] != ' ' {
				start = i
				break
			}
		}
		maze[i] = Row{start: start, row: line}
	}

	instructions := parseInstructions(parts[1])

	return maze, instructions
}

type Face struct {
	rows [50]string
	//use right=0, down=1, left=2, up=3
	neighbours [4]*Face
	// how we need to rotate to move from this face
	// to the neighbours face
	// 0 = no rotation, 1=rotate left, 2=flip, 3=rotate right
	rotations [4]int
}

func parseInput2(input string) ([]Face, []string) {
	parts := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	mazeLines := strings.Split(strings.TrimRight(parts[0], "\n"), "\n")
	/*
		input looks like
		 12
		 3
		45
		6
	*/
	face1Rows := [50]string{}
	face2Rows := [50]string{}
	face3Rows := [50]string{}
	face4Rows := [50]string{}
	face5Rows := [50]string{}
	face6Rows := [50]string{}
	for i := 0; i < 50; i++ {
		face1Rows[i] = mazeLines[i][50:100]
		face2Rows[i] = mazeLines[i][100:]
	}
	for i := 50; i < 100; i++ {
		face3Rows[i] = mazeLines[i][50:100]
	}
	for i := 100; i < 150; i++ {
		face4Rows[i] = mazeLines[i][0:50]
		face5Rows[i] = mazeLines[i][50:100]
	}
	for i := 150; i < 200; i++ {
		face6Rows[i] = mazeLines[i][0:50]
	}
	face1 := Face{rows: face1Rows}
	face2 := Face{rows: face2Rows}
	face3 := Face{rows: face3Rows}
	face4 := Face{rows: face4Rows}
	face5 := Face{rows: face5Rows}
	face6 := Face{rows: face6Rows}

	// hard code the rest when you want to

	faces := make([]Face, 6)

	instructions := parseInstructions(parts[1])

	return faces, instructions
}

func parseInstructions(instr string) []string {
	instructions := make([]string, 0)

	// read instructions char by char and
	var buffer bytes.Buffer
	for i := 0; i < len(instr); i++ {
		buffer = *bytes.NewBuffer([]byte(""))
		buffer.WriteByte(instr[i])
		if instr[i] == 'R' || instr[i] == 'L' {
			instructions = append(instructions, buffer.String())
			continue
		}
		j := i + 1
		for j = i + 1; j < len(instr); j++ {
			if instr[j] == 'R' || instr[j] == 'L' {
				instructions = append(instructions, buffer.String())
				break
			}
			buffer.WriteByte(instr[j])
		}
		i = j - 1
	}
	instructions = append(instructions, buffer.String())
	return instructions
}

func main() {
	util.Time(doPart1, "Part 1")
	util.Time(doPart2, "Part 2")
}

func doPart1() {
	maze, instructions := parseInput1(input)
	fmt.Println("Part 1: ", Part1(maze, instructions))
}

func doPart2() {
	faces, instructions := parseInput2(input)
	fmt.Println("Part 2: ", Part2(faces, instructions))
}

type Position struct {
	row int
	col int
	// use 0 for right, 1 for down, 2 for left, 3 for up
	facing int
}

func Part1(maze Maze, instructions []string) int {
	pos := Position{row: 0, col: maze[0].start, facing: 0}

	for _, inst := range instructions {
		if inst == "R" {
			pos.facing = (pos.facing + 1) % 4
		} else if inst == "L" {
			pos.facing = (pos.facing + 3) % 4
		} else {
			amount := util.Atoi(inst)
			for i := 0; i < amount; i++ {
				var nextRow int
				var nextCol int
				switch pos.facing {
				case 0:
					nextRow = pos.row
					nextCol = pos.col + 1
					if nextCol == len(maze[pos.row].row) {
						nextCol = maze[pos.row].start
					}
					break
				case 2:
					nextRow = pos.row
					nextCol = pos.col - 1
					if nextCol == maze[pos.row].start-1 {
						nextCol = len(maze[pos.row].row) - 1
					}
					break
				case 1:
					nextCol = pos.col
					nextRow = pos.row + 1
					if nextRow == len(maze) || pos.col < maze[nextRow].start || pos.col > len(maze[nextRow].row) {
						nextRow = 0
						for i := 0; i < pos.row; i++ {
							if maze[nextRow].start < pos.col && pos.col < len(maze[nextRow].row) {
								break
							}
							nextRow++
						}
					}
					break
				case 3:
					nextCol = pos.col
					nextRow = pos.row - 1
					if nextRow == -1 || pos.col < maze[nextRow].start || pos.col > len(maze[nextRow].row) {
						nextRow = len(maze) - 1
						for i := len(maze) - 1; i > pos.row; i-- {
							if maze[nextRow].start < pos.col && pos.col < len(maze[nextRow].row) {
								break
							}
							nextRow--
						}
					}
					break
				}

				if maze[nextRow].row[nextCol] == '#' {
					break
				}
				pos.col = nextCol
				pos.row = nextRow
			}
		}
	}

	return 1000*(pos.row+1) + 4*(pos.col+1) + pos.facing
}

func Part2(faces []Face, instructions []string) int {
	// let's hardcode our cube!
	// we need to convert the input to a cube
	// but really you just need to know the four neighbours of each tile
	// then when you leave one face you can land on the next face and reorient yourself correctly?

	return -1
}
