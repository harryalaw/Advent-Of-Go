package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

type CircleNode struct {
	next  *CircleNode
	prev  *CircleNode
	value int
}

func parseInput(input string) []*CircleNode {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	nodes := make([]*CircleNode, len(lines))

	for i, line := range lines {
		nodes[i] = &CircleNode{value: util.Atoi(line)}
	}

	nodes[0].prev = nodes[len(nodes)-1]
	nodes[0].next = nodes[1]
	nodes[len(nodes)-1].prev = nodes[len(nodes)-2]
	nodes[len(nodes)-1].next = nodes[0]

	for i := 1; i < len(nodes)-1; i++ {
		nodes[i].prev = nodes[i-1]
		nodes[i].next = nodes[i+1]
	}

	return nodes
}

func main() {
	util.Time(doPart1, "Part 1")
	util.Time(doPart2, "Part 2")
}

func doPart1() {
	nodes := parseInput(input)
	fmt.Println("Part 1: ", Part1(nodes))
}

func doPart2() {
	nodes := parseInput(input)
	fmt.Println("Part 2: ", Part2(nodes))
}

func Part1(nodes []*CircleNode) int {
	mix(&nodes)

	return findCoords(nodes)
}

func Part2(nodes []*CircleNode) int {
	for _, node := range nodes {
		node.value *= 811589153
	}
	for i := 0; i < 10; i++ {
		mix(&nodes)
	}

	return findCoords(nodes)
}

func mix(nodes *[]*CircleNode) {
	l := len(*nodes)
	for _, node := range *nodes {
		move(node, (node.value % (l - 1)))
	}
}

func move(node *CircleNode, value int) *CircleNode {
	temp := node.prev
	node.prev.next = node.next
	node.next.prev = node.prev

	distance := util.IntAbs(value)
	for i := 0; i < distance; i++ {
		if value > 0 {
			temp = temp.next
		} else {
			temp = temp.prev
		}
	}

	node.prev = temp
	node.next = temp.next
	node.next.prev = node
	node.prev.next = node

	return node
}

func findCoords(nodes []*CircleNode) int {
	zeroIdx := -1
	for i, node := range nodes {
		if node.value == 0 {
			zeroIdx = i
		}
	}

	total := 0
	currentNode := nodes[zeroIdx]
	for i := 1; i <= 3000; i++ {
		currentNode = currentNode.next
		if i%1000 == 0 {
			total += currentNode.value
		}
	}
	return total
}
