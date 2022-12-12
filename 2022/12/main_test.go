package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

func TestPart1(t *testing.T) {
	input, start, end := parseInput(testData)

	value := Part1(input, start, end)

	expected := 31

	if value != expected {
		t.Errorf("Wrong value, expected=%d, got=%d", expected, value)
	}
}

func TestPart2(t *testing.T) {
	input, _, end := parseInput(testData)

	value := Part2(input, end)
	expected := 29

	if value != expected {
		t.Errorf("Wrong value, expected=%d, got=%d", expected, value)
	}
}
