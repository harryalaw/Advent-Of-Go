package main

import (
	"fmt"
	"time"
)

func Time(function func(), part string) {
	start := time.Now()
	function()
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", part, elapsed)
}

func main() {
	Time(Part1, "Part1")
	Time(Part2, "Part2")
}

func Part1() {
	println(bruteForce(33100000))
}
func Part2() {
	println(bruteForce2(33100000))
}

func bruteForce(target int) int {
	length := target / 40
	houses := make([]int, length+1)

	for i := 1; i <= length; i++ {
		for j := i; j <= length; j += i {
			houses[j] += 10 * i
		}
	}

	for i := 0; i <= length; i++ {
		if houses[i] > target {
			return i
		}
	}
	return -1
}

func bruteForce2(target int) int {
	length := target / 40
	houses := make([]int, length+1)

	for i := 1; i <= length; i++ {
		numIters := 0
		for j := i; j <= length && numIters < 50; j += i {
			houses[j] += 11 * i
			numIters++
		}
	}

	for i := 0; i <= length; i++ {
		if houses[i] > target {
			return i
		}
	}
	return -1
}
