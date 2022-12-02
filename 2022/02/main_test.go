package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

func TestParseInput(t *testing.T) {
	expected := []string{
		"A Y",
		"B X",
		"C Z",
	}

	parsedInput := parseInput(testData)

	if len(parsedInput) != len(expected) {
		t.Errorf("Did not create %d rows, got=%d", len(expected), len(parsedInput))
	}

	for i, actual := range parsedInput {
		if actual != expected[i] {
			t.Errorf("game state incorrect, expected=%s got=%s", expected[i], actual)
		}
	}
}

func TestPart1(t *testing.T) {
	input := parseInput(testData)

	value := Part1(input)

	if value != 15 {
		t.Errorf("wrong score, expected=15 got=%d", value)
	}
}

func TestPart2(t *testing.T) {
	input := parseInput(testData)

	value := Part2(input)

	if value != 12 {
		t.Errorf("wrong score, expected=12 got=%d", value)
	}
}
