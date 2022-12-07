package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

/*
	parse input into a tree
	each node has children
	probably include a parent as well?
*/
type Node struct {
	children []*Node
	parent   *Node
	name     string
	size     int
}

func NewNode(parent *Node, name string) Node {
	children := make([]*Node, 0)
	newNode := Node{children: children, parent: parent, size: 0, name: name}

	parent.children = append(parent.children, &newNode)
	return newNode
}

func NewLeaf(parent *Node, name string, size int) Node {
	var children []*Node = nil
	newNode := Node{children: children, parent: parent, size: size, name: name}

	parent.children = append(parent.children, &newNode)
	return newNode
}

func parseInput(input string) *Node {
	data := strings.Split(strings.TrimSpace(input), "\n$ ")

	superParent := Node{children: make([]*Node, 0), parent: nil, size: 0, name: "root"}
	root := NewNode(&superParent, "/")

	parent := &superParent
	current := &root

	for i, line := range data {
		if i == 0 {
			continue
		}
		result := strings.Split(line, "\n")
		if len(result) == 1 && strings.HasPrefix(result[0], "cd") {
			dir := strings.Fields(result[0])[1]
			switch dir {
			case "..":
				current = parent
				parent = parent.parent
			default:
				for _, child := range current.children {
					if (child.name) == dir {
						parent = current
						current = child
					}
				}
			}
		} else {
			for j, res := range result {
				if j == 0 {
					continue
				}
				fields := strings.Fields(res)
				if fields[0] == "dir" {
					// dir e
					NewNode(current, fields[1])
				} else {
					// 62596 h.lst
					NewLeaf(current, fields[1], util.Atoi(fields[0]))
				}
			}
		}
	}

	return &root
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

func Part1(root *Node) int {
	sizeTree(root)
	// printTree(root, "")
	return findSmallDirs(root)
}

func Part2(root *Node) int {
	sizeTree(root)
	totalSpace := 70000000
	totalUnused := totalSpace - root.size
	totalNeeded := 30000000 - totalUnused
	return biggestDeletableNode(root, totalNeeded)
}

func printTree(root *Node, indent string) {
	fmt.Printf("%s%s: %d\n", indent, root.name, root.size)
	for _, child := range root.children {
		printTree(child, indent+"  ")
	}
}

func sizeTree(node *Node) int {
	size := node.size
	for _, child := range node.children {
		size += sizeTree(child)
	}
	node.size = size

	return node.size
}

func findSmallDirs(node *Node) int {
	total := 0

	if len(node.children) == 0 {
		return total
	}

	if node.size <= 100000 {
		total += node.size
	}
	for _, child := range node.children {
		total += findSmallDirs(child)
	}

	return total
}

func biggestDeletableNode(node *Node, diff int) int {
	if len(node.children) == 0 {
		return 70000000
	}
	minValue := 70000000

	if node.size >= diff {
		minValue = util.IntMin(minValue, node.size)
	}
	for _, child := range node.children {
		minValue = util.IntMin(minValue, biggestDeletableNode(child, diff))
	}

	return minValue
}
