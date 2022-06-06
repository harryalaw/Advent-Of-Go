package main

import (
	util "2015/Util"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	util.Time(func() { part1("test.txt") }, "Part1 Test")
	util.Time(func() { part1("data.txt") }, "Part1")

	util.Time(func() { part2("test.txt") }, "Part2 Test")
	util.Time(func() { part2("data.txt") }, "Part2")

}

type node struct {
	name       string
	neighbours map[string]int
}

func part1(fileName string) int {
	graph := parseInput(fileName)

	places := make([]string, len(graph))
	i := 0
	for _, value := range graph {
		places[i] = value.name
		i++
	}

	return coreLogic(&graph, places)
}

func part2(fileName string) int {
	graph := parseInput(fileName)
	for person := range graph {
		addToGraph(&graph, "Me", person, 0)
		addToGraph(&graph, "Me", person, 0)
	}

	places := make([]string, len(graph))
	i := 0
	for _, value := range graph {
		places[i] = value.name
		i++
	}
	return coreLogic(&graph, places)
}

func coreLogic(graph *map[string]node, places []string) int {
	minLength := 0

	Perm(places, func(a []string) {
		newLength := calculateCycleLength(graph, a)
		if newLength > minLength {
			minLength = newLength
		}
	})
	fmt.Println(minLength)
	return minLength
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

func calculateCycleLength(graph *map[string]node, path []string) int {
	totalLength := 0
	for i := 0; i < len(path); i++ {
		source := path[i]
		target := path[(i+1)%len(path)]
		length := (*graph)[source].neighbours[target]
		totalLength += length
	}

	for i := 0; i < len(path); i++ {
		source := path[i]
		j := (i - 1) % len(path)
		if j < 0 {
			j += len(path)
		}
		target := path[j]
		length := (*graph)[source].neighbours[target]
		totalLength += length
	}
	return totalLength
}

func parseInput(fileName string) map[string]node {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	data := strings.Split(string(content), "\n")

	graph := make(map[string]node)

	for _, line := range data {
		withoutPeriod := strings.Replace(line, ".", "", -1)
		fields := strings.Fields(withoutPeriod)

		node1 := fields[0]
		node2 := fields[10]
		distance, parseErr := strconv.Atoi(fields[3])
		positive := fields[2]

		if positive == "lose" {
			distance *= -1
		}
		if parseErr != nil {
			panic(parseErr)
		}

		addToGraph(&graph, node1, node2, distance)
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
