package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

func TestParseInput(t *testing.T) {
	expected := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	parsedInput := parseInput(testData)

	if len(parsedInput) != len(expected) {
		t.Errorf("Did not create %d rows, got=%d", len(expected), len(parsedInput))
	}

	for i, line := range parsedInput {
		if line != expected[i] {
			t.Errorf("line parsing failed, expected=%s, got=%s", expected[i], line)
		}
	}
}

func TestPart1(t *testing.T) {
	input := parseInput(testData)

	value := Part1(input)

	if value != 157 {
		t.Errorf("expected=157, got=%d", value)
	}
}

func TestPart2(t *testing.T) {
	input := parseInput(testData)

	value := Part2(input)

	if value != 70 {
		t.Errorf("expected=70, got=%d", value)
	}
}
