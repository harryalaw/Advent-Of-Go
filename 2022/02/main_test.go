package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

func TestParseInput(t *testing.T) {
	expected := []RPS{
		{opponent: "A", player: "Y"},
		{opponent: "B", player: "X"},
		{opponent: "C", player: "Z"},
	}

	parsedInput := parseInput(testData)

	if len(parsedInput) != len(expected) {
		t.Errorf("Did not create %d rows, got=%d", len(expected), len(parsedInput))
	}

	for i, actual := range parsedInput {
		if actual.opponent != expected[i].opponent {
			t.Errorf("opponent incorrect, expected=%s got=%s", expected[i], actual)
		}
		if actual.player != expected[i].player {
			t.Errorf("player incorrect, expected=%s got=%s", expected[i], actual)
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
