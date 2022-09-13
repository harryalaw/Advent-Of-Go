package main

import (
	"github.com/harryalaw/advent-of-go/util"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	util.Time(part1, "Part1")
	// test()
	util.Time(part2, "Part2")
}

func part1() {
	fmt.Println(dp("data.txt", 150))
}
func part2() {
	fmt.Println(dp2("data.txt", 150))
}

func test() {
	fmt.Println(dp("test.txt", 25))
	fmt.Println(dp2("test.txt", 25))
}

func parseInput(file string) *[]int {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	content := strings.Split(string(data), "\n")

	out := make([]int, len(content))

	for i, el := range content {
		temp, err := strconv.Atoi(el)
		if err != nil {
			panic(err)
		}
		out[i] = temp
	}

	return &out
}

func dp(file string, targetNum int) int {
	numToCounts := make(map[int]int)
	numbers := parseInput(file)
	for _, number := range *numbers {
		nextNumToCounts := make(map[int]int)
		for key, value := range numToCounts {
			nextNumToCounts[key] += value
			if key+number <= targetNum {
				nextNumToCounts[key+number] += value
			}
		}
		nextNumToCounts[number] += 1

		numToCounts = nextNumToCounts
	}

	return numToCounts[targetNum]
}

type pair struct {
	size int
	ways int
}

// more complicated than you're thinking this is
func dp2(file string, targetNum int) int {
	dpMap := make(map[int]pair)
	numbers := parseInput(file)

	for _, number := range *numbers {
		nextMap := make(map[int]pair)
		// copy old map
		for key, value := range dpMap {
			nextMap[key] = value
		}
		for key, value := range dpMap {
			// for new combo
			/*
				if we have already stored a pair,
				then if the stored size is equal to prevSize + 1
					-> we can add the ways together
				if the stored size is more than prevSize + 1
					-> Then we can reset to new value
				if the stored size is less than prevSize + 1
					-> Then we have a more efficient way to get here
			*/
			if key+number > targetNum {
				continue
			}

			storedPair, isPresent2 := dpMap[key+number]
			if isPresent2 && storedPair.size == value.size+1 {
				storedPair.ways += value.ways
				nextMap[key+number] = storedPair
			} else if (isPresent2 && storedPair.size > value.size+1) ||
				!isPresent2 {
				storedPair.size = value.size + 1
				storedPair.ways = value.ways
				nextMap[key+number] = storedPair
			}
		}
		prevValue, isPresent := dpMap[number]
		if isPresent && prevValue.size == 1 {
			prevValue.ways += 1
			nextMap[number] = prevValue
		} else {
			nextMap[number] = pair{size: 1, ways: 1}
		}
		dpMap = nextMap
	}

	return dpMap[targetNum].ways
}
