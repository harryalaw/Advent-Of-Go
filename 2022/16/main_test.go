package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

func TestPart1(t *testing.T) {
	input := parseInput(testData)

	expected := 1651
	value := Part1(input)
	if value != expected {
		t.Errorf("value not correct, expected=%d, got=%d", expected, value)
	}

}

func TestPart2(t *testing.T) {
	input := parseInput(testData)

	value := Part2(input)
	expected := 1707

	if value != expected {
		t.Errorf("value not correct, expected=%d, got=%d", expected, value)
	}
}
