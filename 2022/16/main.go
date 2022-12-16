package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

//
type Graph [][]int

type Edge struct {
	distance int
	node     *Node
}

type Node struct {
	name  string
	cost  int
	edges []Edge
}

func parseInput(input string) map[string]*Node {

	lines := strings.Split(strings.TrimSpace(input), "\n")
	graph := map[string]*Node{}

	// find all the nodes and make them
	for _, line := range lines {
		name := line[6:8]
		var cost int
		if line[24] == ';' {
			cost = util.Atoi(line[23:24])
		} else {
			cost = util.Atoi(line[23:25])
		}

		graph[name] = &Node{name: name, cost: cost}
	}

	for _, line := range lines {
		name := line[6:8]
		current := graph[name]

		nbrIdx := 49
		if line[nbrIdx] == ' ' {
			nbrIdx++
		}
		nbrs := strings.Split(line[nbrIdx:], ", ")
		edges := make([]Edge, len(nbrs))
		for i, nbr := range nbrs {
			nbrNode := graph[nbr]
			edges[i] = Edge{distance: 1, node: nbrNode}
		}
		current.edges = edges
		graph[name] = current
	}

	// build reduced graph of only nonzero cost nodes
	nonZeroCostNodes := make([]*Node, 0)
	reducedGraph := map[string]*Node{}

	for _, node := range graph {
		if node.cost != 0 || node.name == "AA" {
			nonZeroCostNodes = append(nonZeroCostNodes, node)
			newNode := &Node{node.name, node.cost, []Edge{}}
			reducedGraph[newNode.name] = newNode
		}
	}

	createReducedGraph(&reducedGraph, &graph)

	return reducedGraph
}

func createReducedGraph(reducedGraph, originalGraph *map[string]*Node) {
	for _, node := range *reducedGraph {
		distances := map[string]int{}
		// now want to do a BFS on the original graph
		// for each nbr of my node, add an edge to edge with d = 1
		// then for each nbr of a nbr, if not in edges, add with d= 2
		// repeate until edges is full
		// then give node edges:= []Edge
		originalNbrs := (*originalGraph)[node.name].edges
		nextNbrs := make([]string, len(originalNbrs))
		for i, nbr := range originalNbrs {
			nextNbrs[i] = nbr.node.name
		}

		dist := 0
		for len(distances) != len(*originalGraph)-1 {
			dist++
			tempNbrs := []string{}
			for _, nbr := range nextNbrs {
				if _, ok := distances[nbr]; ok {
					continue
				}
				distances[nbr] = dist
				ogNeighbours := (*originalGraph)[nbr]
				for _, edge := range ogNeighbours.edges {
					if edge.node.name == node.name {
						continue
					}
					if _, ok := distances[edge.node.name]; !ok {
						tempNbrs = append(tempNbrs, edge.node.name)
					}
				}
			}
			nextNbrs = tempNbrs
		}

		edges := make([]Edge, 0)
		for name, distance := range distances {
			newNode, ok := (*reducedGraph)[name]
			if !ok {
				continue
			}
			edges = append(edges, Edge{
				distance: distance,
				node:     newNode,
			})
		}
		node.edges = edges
	}
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
	fmt.Println("Part 2: ", Part2(data))
}

func Part1(graph map[string]*Node) int {
	maxFlow := 0
	for _, node := range graph {
		maxFlow += node.cost
	}

	bestScore := 0

	return calcScore(graph["AA"],
		30,
		0,
		map[string]bool{
			"AA": true,
		},
		maxFlow,
		&bestScore,
	)
}

func calcScore(
	current *Node,
	timeRemaining, currentScore int,
	visited map[string]bool,
	maxFlow int,
	bestScore *int,
) int {
	if timeRemaining <= 0 {
		return 0
	}

	thisScore := timeRemaining * current.cost
	currentScore += thisScore

	if currentScore+(timeRemaining*maxFlow) < *bestScore {
		return 0
	}

	maxChildScore := 0
	for _, edge := range current.edges {
		if seen, ok := visited[edge.node.name]; ok && seen {
			continue
		}
		visited[edge.node.name] = true
		newScore := calcScore(edge.node, timeRemaining-edge.distance-1, currentScore, visited, maxFlow, bestScore)
		visited[edge.node.name] = false
		if newScore > maxChildScore {
			maxChildScore = newScore
		}
	}

	if currentScore+maxChildScore > *bestScore {
		*bestScore = currentScore + maxChildScore
	}

	return thisScore + maxChildScore
}

func Part2(graph map[string]*Node) int {
	return -1
}
