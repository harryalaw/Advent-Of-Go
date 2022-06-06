package main

import (
	util "2015/Util"
	"fmt"
	"io/ioutil"
)

func main() {
	util.Time(func() { fmt.Println(part1()) }, "Part1")
	util.Time(func() { fmt.Println(part2()) }, "Part2")
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
