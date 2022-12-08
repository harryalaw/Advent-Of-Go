package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

func TestPart1(t *testing.T) {
	input := parseInput(testData)

	value := Part1(input)

	if value != 21 {
		t.Errorf("Value incorrect, expected=%d, got=%d", 21, value)
	}
}

func TestPart2(t *testing.T) {
	input := parseInput(testData)

	value := Part2(input)

	if value != 8 {
		t.Errorf("Value incorrect, expected=%d, got=%d", 8, value)
	}
}
