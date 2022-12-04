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

	if value != 2 {
		t.Errorf("expected=2, got=%d", value)
	}
}

func TestPart2(t *testing.T) {
	input := parseInput(testData)

	value := Part2(input)

	if value != 4 {
		t.Errorf("expected=4, got=%d", value)
	}
}
