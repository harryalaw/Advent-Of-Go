package main

import (
	"github.com/harryalaw/advent-of-go/util"
	"io/ioutil"
)

func main() {
	println(tests(false))
	util.Time(func() { solve(part1) }, "Part1")
	util.Time(func() { solve(part2) }, "Part2")
}

func parseInput() string {
	content, err := ioutil.ReadFile("data.txt")

	if err != nil {
		panic(err)
	}

	return string(content)
}

type coord struct {
	x int
	y int
}

func solve(part func(string) int) int {
	value := part(parseInput())
	println(value)
	return value
}

func part1(input string) int {
	santaPos := coord{0, 0}
	santas := [2]*coord{&santaPos, &santaPos}

	return coreLogic(input, santas[:])
}

func part2(input string) int {
	santaPos := coord{0, 0}
	roboPos := coord{0, 0}

	santas := [2]*coord{&santaPos, &roboPos}

	return coreLogic(input, santas[:])
}

func coreLogic(input string, santas []*coord) int {
	houses := make(map[coord]int)

	houses[coord{0, 0}] = 1

	for ind, char := range input {
		activeSanta := santas[ind%2]
		switch char {
		case '^':
			activeSanta.y++
		case '>':
			activeSanta.x++
		case '<':
			activeSanta.x--
		case 'v':
			activeSanta.y--
		}
		houses[*activeSanta] = houses[*activeSanta] + 1
	}
	return len(houses)
}

func test1() bool {
	result := part1(">")
	return result == 2
}

func test2() bool {
	result := part1("^>v<")
	return result == 4
}

func test3() bool {
	result := part1("^v^v^v^v^v")
	return result == 2
}

func test4() bool {
	result := part2("^>v<")
	return result == 3
}
func test5() bool {
	result := part2("^v")
	return result == 3
}
func test6() bool {
	result := part2("^v^v^v^v^v")
	return result == 11
}

func actual1() bool {
	result := part1(parseInput())
	return result == 2572
}

func actual2() bool {
	result := part2(parseInput())
	return result == 2631
}

func tests(actual bool) bool {
	pass1 := test1()
	pass2 := test2()
	pass3 := test3()
	pass4 := test4()
	pass5 := test5()
	pass6 := test6()

	passActual1 := true
	passActual2 := true

	if actual {
		passActual1 = actual1()
		passActual2 = actual2()
	}

	return pass1 && pass2 && pass3 && pass4 && pass5 && pass6 && passActual1 && passActual2
}
