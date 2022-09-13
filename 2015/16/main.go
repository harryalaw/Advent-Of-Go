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
	util.Time(part2, "Part2")
}

type Sue struct {
	children    int
	cats        int
	samoyeds    int
	pomeranians int
	akitas      int
	vizslas     int
	goldfish    int
	trees       int
	cars        int
	perfumes    int
	number      int
}

func part1() {
	realSue := Sue{children: 3,
		cats:        7,
		samoyeds:    2,
		pomeranians: 3,
		akitas:      0,
		vizslas:     0,
		goldfish:    5,
		trees:       3,
		cars:        2,
		perfumes:    1,
	}

	potentialSues := parseInput()
	potentialMatches := make([]int, 0)
	for index, sue := range potentialSues {
		if matches(&realSue, &sue) {
			potentialMatches = append(potentialMatches, index)
		}
	}

	// fmt.Printf("The real Sue is: %+v\n", realSue)
	// for _, id := range potentialMatches {
	// 	fmt.Printf("%+v\n", potentialSues[id])
	// }

	fmt.Printf("%+v\n", potentialMatches[0]+1)
}
func part2() {
	realSue := Sue{children: 3,
		cats:        7,
		samoyeds:    2,
		pomeranians: 3,
		akitas:      0,
		vizslas:     0,
		goldfish:    5,
		trees:       3,
		cars:        2,
		perfumes:    1,
	}

	potentialSues := parseInput()
	potentialMatches := make([]int, 0)
	for i, sue := range potentialSues {
		if matches2(&realSue, &sue) {
			potentialMatches = append(potentialMatches, i)
		}
	}

	// fmt.Printf("The real Sue is: %+v\n", realSue)
	// for _, id := range potentialMatches {
	// 	fmt.Println("-----")
	// 	fmt.Printf("%+v\n", realSue)
	// 	fmt.Printf("%+v\n\n", potentialSues[id])
	// }

	fmt.Printf("%+v\n", potentialMatches[0]+1)
}
func matches(real *Sue, potential *Sue) bool {
	return (potential.children == -1 || potential.children == 3) &&
		(potential.cats == -1 || potential.cats == 7) &&
		(potential.samoyeds == -1 || potential.samoyeds == 2) &&
		(potential.pomeranians == -1 || potential.pomeranians == 3) &&
		(potential.akitas == -1 || potential.akitas == 0) &&
		(potential.vizslas == -1 || potential.vizslas == 0) &&
		(potential.goldfish == -1 || potential.goldfish == 5) &&
		(potential.trees == -1 || potential.trees == 3) &&
		(potential.cars == -1 || potential.cars == 2) &&
		(potential.perfumes == -1 || potential.perfumes == 1)
}
func matches2(real *Sue, potential *Sue) bool {
	return (potential.children == -1 || potential.children == 3) &&
		(potential.samoyeds == -1 || potential.samoyeds == 2) &&
		(potential.pomeranians == -1 || potential.pomeranians < 3) &&
		(potential.goldfish == -1 || potential.goldfish < 5) &&
		(potential.akitas == -1 || potential.akitas == 0) &&
		(potential.vizslas == -1 || potential.vizslas == 0) &&
		(potential.cats == -1 || potential.cats > 7) &&
		(potential.trees == -1 || potential.trees > 3) &&
		(potential.cars == -1 || potential.cars == 2) &&
		(potential.perfumes == -1 || potential.perfumes == 1)
}

// As you're about to send the thank you note, something in the MFCSAM's instructions catches your eye. Apparently, it has an outdated retroencabulator, and so the output from the machine isn't exact values - some of them indicate ranges.

// In particular, the cats and trees readings indicates that there are greater than that many (due to the unpredictable nuclear decay of cat dander and tree pollen), while the pomeranians and goldfish readings indicate that there are fewer than that many (due to the modial interaction of magnetoreluctance).

// What is the number of the real Aunt Sue?

func parseInput() []Sue {
	data, err := ioutil.ReadFile("data.txt")

	if err != nil {
		panic(err)
	}

	potentialSues := make([]Sue, 500)

	content := strings.Split(string(data), "\n")

	for i, sue := range content {
		newSue := Sue{children: -1,
			cats:        -1,
			samoyeds:    -1,
			pomeranians: -1,
			akitas:      -1,
			vizslas:     -1,
			goldfish:    -1,
			trees:       -1,
			cars:        -1,
			perfumes:    -1,
		}
		// 	cats: -1
		// }
		fields := strings.Fields(sue)
		newSue.number = i + 1
		for j, word := range fields {
			switch word {
			case "children:":
				newSue.children = atoi(fields[j+1])
			case "cats:":
				newSue.cats = atoi(fields[j+1])
			case "samoyeds:":
				newSue.samoyeds = atoi(fields[j+1])
			case "pomeranians:":
				newSue.pomeranians = atoi(fields[j+1])
			case "akitas:":
				newSue.akitas = atoi(fields[j+1])
			case "vizslas:":
				newSue.vizslas = atoi(fields[j+1])
			case "goldfish:":
				newSue.goldfish = atoi(fields[j+1])
			case "trees:":
				newSue.trees = atoi(fields[j+1])
			case "cars:":
				newSue.cars = atoi(fields[j+1])
			case "perfumes:":
				newSue.perfumes = atoi(fields[j+1])
			}
		}
		potentialSues[i] = newSue
	}
	return potentialSues
}

func atoi(word string) int {
	word = strings.Replace(word, ",", "", -1)
	num, err := strconv.Atoi(word)
	if err != nil {
		panic(err)
	}

	return num
}
