package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

func TestPart1(t *testing.T) {
	maze, instructions := parseInput1(testData)

	value := Part1(maze, instructions)
	expected := 6032
	if value != expected {
		t.Errorf("Value not correct, expected=%d, got=%d", expected, value)
	}
}

// func TestPart2(t *testing.T) {
// 	faces, instructions := parseInput2(testData)
//
// 	value := Part2(faces, instructions)
// 	expected := -1
// 	if value != expected {
// 		t.Errorf("Value not correct, expected=%d, got=%d", expected, value)
// 	}
// }
