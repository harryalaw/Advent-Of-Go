package main

import (
	_ "embed"
	"github.com/harryalaw/advent-of-go/util"
	"strconv"
	"strings"
)

func main() {
	util.Time(part1, "Part1")
	util.Time(part2, "Part2")
}

//go:embed data.txt
var input string

func parseInput() []string {
	lines := splitOnNewLines(string(input))
	return lines
}

func part1() {
	lines := parseInput()

	totalPaper := 0
	for _, line := range lines {
		w, h, l := splitLine(line)

		totalPaper += surfaceArea(w, h, l)
	}

	println(totalPaper)
}

func part2() {
	lines := parseInput()

	totalPaper := 0
	for _, line := range lines {
		w, h, l := splitLine(line)

		totalPaper += volume(w, h, l)
	}
	println(totalPaper)
}

func splitOnNewLines(line string) []string {
	return strings.Split(line, "\n")
}

func splitLine(line string) (w, h, l int) {
	splitValues := strings.Split(line, "x")

	w, _ = strconv.Atoi(splitValues[0])
	h, _ = strconv.Atoi(splitValues[1])
	l, _ = strconv.Atoi(splitValues[2])

	return
}

func surfaceArea(w, h, l int) int {
	sideA := 2 * w * h
	sideB := 2 * h * l
	sideC := 2 * l * w

	minSide := sideA
	if sideB < minSide {
		minSide = sideB
	}
	if sideC < minSide {
		minSide = sideC
	}
	minSide /= 2
	return sideA + sideB + sideC + minSide
}

func volume(w, h, l int) int {
	perimA := 2*w + 2*h
	perimB := 2*h + 2*l
	perimC := 2*l + 2*w

	minPerim := perimA
	if perimB < minPerim {
		minPerim = perimB
	}
	if perimC < minPerim {
		minPerim = perimC
	}
	return w*h*l + minPerim
}
