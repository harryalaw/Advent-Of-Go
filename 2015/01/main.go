package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func solve(part func() int) int {
	value := part()
	println(value)
	return value
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

func part1() int {
	content, _ := ioutil.ReadFile("data.txt")

	floor := 0

	for _, char := range string(content) {
		if char == '(' {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

func part2() int {
	content, _ := ioutil.ReadFile("data.txt")

	floor := 0

	for index, char := range string(content) {
		if char == '(' {
			floor++
		} else {
			floor--
		}
		if floor == -1 {
			return index + 1
		}
	}
	return -1
}
