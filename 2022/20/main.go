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
	next       *CircleNode
	initialPos int
	value      int
}

func parseInput(input string) (*CircleNode, *CircleNode, []int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	numbers := make([]int, len(lines))

	for i, line := range lines {
		numbers[i] = util.Atoi(line)
	}
	start := CircleNode{initialPos: 0, value: numbers[0]}
	end := CircleNode{initialPos: len(numbers) - 1, value: numbers[len(numbers)-1]}
	end.next = &start

	nextNode := &end
	for i := len(numbers) - 2; i >= 1; i-- {
		node := CircleNode{initialPos: i, value: numbers[i]}
		node.next = nextNode
		nextNode = &node
	}
	start.next = nextNode

	return &start, &end, numbers
}

func main() {
	util.Time(doPart1, "Part 1")
	util.Time(doPart2, "Part 2")
}

func doPart1() {
	start, end, numbers := parseInput(input)
	fmt.Println("Part 1: ", Part1(start, end, numbers))
}

func doPart2() {
	start, _, _ := parseInput(input)
	fmt.Println("Part 2: ", Part2(*start))
}

func Part1(start, end *CircleNode, numbers []int) int {

	// count up through the ids and move the nodes around!

	// find the node with id=nextId
	// make the previous nodes next value node.next
	// then move next number times and then insert again
	prevNode := end
	currentNode := start
	for i := 0; i < len(numbers); i++ {
		for currentNode.initialPos != i {
			prevNode = currentNode
			currentNode = currentNode.next
		}
		// remove current node from the cycle
		prevNode.next = currentNode.next

		// move around value times
		moves := currentNode.value
		if moves < 0 {
			moves = len(numbers) + moves - 1
		}
		for i := 0; i < moves; i++ {
			prevNode = prevNode.next
		}

		// reinsert the value to the loop
		prevNext := prevNode.next
		prevNode.next = currentNode
		currentNode.next = prevNext
	}

	// to get the score
	// find the 0
	// then modulo 1000/2000/3000 and move that far
	// then add the sums
	for currentNode.value != 0 {
		currentNode = currentNode.next
	}

	for i := 0; i < len(numbers); i++ {
		currentNode = currentNode.next
	}

	total := 0
	for i := 1; i <= 3000; i++ {
		currentNode = currentNode.next
		if i%1000 == 0 {
			total += currentNode.value
		}
	}

	return total
}

func mod(a, b int) int {
	modulus := a % b
	if modulus < 0 {
		modulus += b
	}
	return modulus
}

func Part2(numbers CircleNode) int {
	return -1
}
