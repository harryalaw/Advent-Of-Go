package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

func TestPart1(t *testing.T) {
	stacks, instructions := parseInput(testData)

	value := Part1(stacks, instructions)

	if value != "CMZ" {
		t.Errorf("Expected=%s, got=%s", "CMZ", value)
	}
}

func TestPart2(t *testing.T) {
	stacks, instructions := parseInput(testData)

	value := Part2(stacks, instructions)

	if value != "MCD" {
		t.Errorf("Expected=%s, got=%s", "MCD", value)
	}
}
