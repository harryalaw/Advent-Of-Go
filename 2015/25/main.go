package main

import (
	util "2015/Util"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type coord struct {
	row int
	col int
}

//To continue, please consult the code grid in the manual.  Enter the code at row 3010, column 3019.
func main() {
	util.Time(doPart1, "Part1")
}

func doPart1() {
	target := parseInput("data.txt")
	fmt.Println(part1(target))
}

func part1(target coord) int {
	pos := coord{row:1, col:1}
	currentValue := 20151125;

	MULTIPLIER := 252533
	DIVIDER := 33554393

	for (pos.row != target.row) || (pos.col != target.col) {
		currentValue = (currentValue * MULTIPLIER) % DIVIDER
		pos = nextCoord(pos);
	}


	return currentValue
}

func nextCoord(current coord) coord {
	nextCoord := coord{}
	if current.row == 1 {
		nextCoord.row = current.col + 1
		nextCoord.col = 1
	} else {
		nextCoord.row = current.row - 1
		nextCoord.col = current.col + 1
	}

	return nextCoord
}

func parseInput(file string) coord {
	content, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	data := string(content)

	re := regexp.MustCompile("(row|column) [0-9]+")
	matches := re.FindAllString(data, -1)

	if len(matches) != 2 {
		panic("Too many matches")
	}

	output := coord{}
	for _, match := range matches {
		fields := strings.Fields(match)
		switch fields[0] {
		case "row":
			output.row = util.Atoi(fields[1])
		case "column":
			output.col = util.Atoi(fields[1])
		}
	}

	return output
}
