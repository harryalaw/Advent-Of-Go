package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

func TestPart1(t *testing.T) {
	start, end, numbers := parseInput(testData)

	value := Part1(start, end, numbers)
	expected := 3

	if value != expected {
		t.Errorf("Value not correct, expected=%d, got=%d", expected, value)
	}
}

func TestPart2(t *testing.T) {
	// input, _, _ := parseInput(testData)
	//
	// value := Part2(*input)
	// expected := 3
	//
	// if value != expected {
	// 	t.Errorf("Value not correct, expected=%d, got=%d", expected, value)
	// }
}
