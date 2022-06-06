package main

import (
	util "2015/Util"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Reindeer struct {
	name     string
	speed    int
	duration int
	rest     int
	points   int
	position int
}

func parseInput(fileName string) []Reindeer {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	data := strings.Split(string(content), "\n")

	output := make([]Reindeer, len(data))
	for i, line := range data {
		words := strings.Fields(line)
		speed := atoi(words[3])
		duration := atoi(words[6])
		rest := atoi(words[13])
		output[i] = Reindeer{
			name:     words[0],
			speed:    speed,
			duration: duration,
			rest:     rest,
		}
	}
	return output
}

func atoi(word string) int {
	num, err := strconv.Atoi(word)
	if err != nil {
		panic(err)
	}
	return num
}

func Time(part string, function func()) {
	start := time.Now()
	function()
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", part, elapsed)
	println()
}

func main() {
	util.Time(func() {
		part1("test.txt", 1000)
	}, "Part1 Test")
	util.Time(func() {
		part1("data.txt", 2503)
	}, "Part1 actual")
	util.Time(func() {
		part2("test.txt", 1000)
	}, "Part2 Test")
	util.Time(func() {
		part2("data.txt", 2503)
	}, "Part2 actual")
}

func part1(fileName string, time int) int {
	reindeer := parseInput(fileName)
	maxDistance := 0
	for _, deer := range reindeer {
		deerDistance := deerTravel(deer, time)
		if deerDistance > maxDistance {
			maxDistance = deerDistance
		}
	}

	println(maxDistance)
	return maxDistance
}

func part2(fileName string, time int) int {
	reindeer := parseInput(fileName)
	furthestDistance := -1

	for i := 1; i < time; i++ {
		for j, deer := range reindeer {
			deer.position = deerTravel(deer, i)
			if deer.position > furthestDistance {
				furthestDistance = deer.position
			}
			reindeer[j] = deer
		}
		for j, deer := range reindeer {
			if deer.position == furthestDistance {
				tempDeer := reindeer[j]
				tempDeer.points += 1
				reindeer[j] = tempDeer

			}
		}
	}

	maxPoints := 0
	for _, deer := range reindeer {
		if deer.points > maxPoints {
			maxPoints = deer.points
		}
	}
	println(maxPoints)
	return maxPoints
}

func deerTravel(deer Reindeer, time int) int {
	remainingTime := time % (deer.duration + deer.rest)
	numberOfFullReps := time / (deer.duration + deer.rest)

	distance := 0
	distance += numberOfFullReps * (deer.speed * deer.duration)

	if remainingTime > deer.duration {
		distance += deer.speed * deer.duration
	} else {
		distance += deer.speed * remainingTime
	}

	return distance
}
