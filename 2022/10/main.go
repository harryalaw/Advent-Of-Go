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

type OpCode struct {
	command     string
	cycleLength int
	value       int
}

var noop OpCode = OpCode{command: "noop", cycleLength: 1, value: 0}

func parseInput(input string) []*OpCode {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	out := make([]*OpCode, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "noop" {
			out[i] = &noop
		} else {
			out[i] = &OpCode{command: fields[0], cycleLength: 2, value: util.Atoi(fields[1])}
		}
	}

	return out
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
	fmt.Printf("Part 2:\n%s\n", Part2(data))
}

func Part1(data []*OpCode) int {
	x := 1
	signalStrength := 0
	ptr := 0
	delayed := 0
	for i := 1; i <= 220; i++ {
		opcode := data[ptr]
		if (i-20)%40 == 0 {
			signalStrength += x * i
		}
		if delayed != 0 {
			x += delayed
			delayed = 0
			ptr++
			continue
		}
		switch opcode.command {
		case "noop":
			ptr++
			continue
		case "addx":
			delayed = opcode.value
		}
	}
	return signalStrength
}

func Part2(data []*OpCode) string {
	var out bytes.Buffer

	x := 1
	ptr := 0
	delayed := 0
	cursorIdx := 0
	for i := 0; i <= 239; i++ {
		opcode := data[ptr]
		cursorIdx = i % 40
		if cursorIdx == 0 && i != 0 {
			out.WriteRune('\n')
		}

		if x == cursorIdx-1 || x == cursorIdx || x == cursorIdx+1 {
			out.WriteRune('#')
		} else {
			out.WriteRune('.')
		}
		if delayed != 0 {
			x += delayed
			delayed = 0
			ptr++
			continue
		}
		if opcode.command == "noop" {
			ptr++
			continue
		} else {
			delayed = opcode.value
		}
	}

	return out.String()
}
