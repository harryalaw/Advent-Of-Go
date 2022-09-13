package main

import (
	"github.com/harryalaw/advent-of-go/util"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type node struct {
	name       string
	neighbours map[string]int
}

func parseInput() map[string]node {
	content, err := ioutil.ReadFile("data.txt")

	if err != nil {
		panic(err)
	}

	data := strings.Split(string(content), "\n")

	graph := make(map[string]node)

	for _, line := range data {
		fields := strings.Fields(line)

		node1 := fields[0]
		node2 := fields[2]
		distance, parseErr := strconv.Atoi(fields[4])

		if parseErr != nil {
			panic(parseErr)
		}

		addToGraph(&graph, node1, node2, distance)
		addToGraph(&graph, node2, node1, distance)
	}

	return graph
}

func addToGraph(graph *(map[string]node), source, target string, dist int) {
	ourNode, isPresent := (*graph)[source]

	if !isPresent {
		newNode := new(node)
		newNode.name = source
		newNode.neighbours = make(map[string]int)
		newNode.neighbours[target] = dist
		(*graph)[source] = *newNode
	} else {
		ourNode.neighbours[target] = dist
		(*graph)[source] = ourNode
	}
}

func part1() int {
	graph := parseInput()

	places := make([]string, len(graph))
	i := 0
	for _, value := range graph {
		places[i] = value.name
		i++
	}

	minLength := math.MaxInt

	Perm(places, func(a []string) {
		newLength := calculatePathLength(&graph, a, max)
		if newLength < minLength {
			minLength = newLength
		}
	})
	return minLength
}

func part2() int {
	graph := parseInput()

	places := make([]string, len(graph))
	i := 0
	for _, value := range graph {
		places[i] = value.name
		i++
	}

	maxLength := 0

	Perm(places, func(a []string) {
		newLength := calculatePathLength(&graph, a, min)
		if newLength > maxLength {
			maxLength = newLength
		}
	})
	return maxLength
}

func solve(part func() int) int {
	value := part()
	println(value)
	return value
}

func main() {
	util.Time(func() { solve(part1) }, "Part1")
	util.Time(func() { solve(part2) }, "Part2")
}

func Perm(a []string, f func([]string)) {
	perm(a, f, 1)
}

func perm(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}

	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func calculatePathLength(graph *map[string]node, path []string, comparator func(int, int) int) int {
	extraLength := math.MaxInt - comparator(0, math.MaxInt)
	totalLength := 0

	for i := 0; i < len(path); i++ {
		source := path[i]
		target := path[(i+1)%len(path)]
		length := (*graph)[source].neighbours[target]
		totalLength += length
		extraLength = comparator(extraLength, (*graph)[source].neighbours[target])
	}

	return totalLength - extraLength
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
