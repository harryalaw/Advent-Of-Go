package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

func TestPart1(t *testing.T) {
	cave, bottom := parseInput(testData)

	value := Part1(cave, bottom)
	expected := 24

	if value != expected {
		t.Errorf("Value wrong. expected=%d, got=%d", expected, value)
	}
}

func TestPart2(t *testing.T) {
	input, bottom := parseInput(testData)

	value := Part2(input, bottom)
	expected := 93

	if value != expected {
		t.Errorf("Value wrong. expected=%d, got=%d", expected, value)
	}
}
