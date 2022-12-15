package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

func TestPart1(t *testing.T) {
	input := parseInput(testData)

	value := Part1(input, 10)
	expected := 26

	if expected != value {
		t.Errorf("Value wrong, expected=%d, got=%d", expected, value)
	}
}

func TestPart2(t *testing.T) {
	input := parseInput(testData)

	value := Part2(input, 20)
	expected := 56000011

	if expected != value {
		t.Errorf("Value wrong, expected=%d, got=%d", expected, value)
	}
}
