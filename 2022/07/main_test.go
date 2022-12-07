package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

// func TestParseInput(t *testing.T) {
// 	expected := []string{}
//
// 	parsedInput := parseInput(testData)
//
// 	if len(parsedInput) != len(expected) {
// 		t.Errorf("Did not create %d rows, got=%d", len(expected), len(parsedInput))
// 	}
//
// 	t.Errorf("Not asserting on contents of parsed input")
// }

func TestPart1(t *testing.T) {
	input := parseInput(testData)

	value := Part1(input)

	if value != 95437 {
		t.Errorf("value not correct, expected=%d, got=%d", 95437, value)
	}
}

func TestPart2(t *testing.T) {
	input := parseInput(testData)

	value := Part2(input)

	if value != 24933642 {
		t.Errorf("value not correct, expected=%d, got=%d", 24933642, value)
	}
}
