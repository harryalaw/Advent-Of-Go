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
	expected := 33

	if value != expected {
		t.Errorf("Value not correct, got=%d, expected=%d", value, expected)
	}
}

type BlueprintTestCase struct {
	blueprint Blueprint
	expected  int
}

func TestProbabilisticBlueprintPart1(t *testing.T) {
	testCases := []BlueprintTestCase{
		{blueprint: Blueprint{oreCost: Cost{4, 0, 0}, clayCost: Cost{2, 0, 0}, obsidianCost: Cost{3, 14, 0}, geodeCost: Cost{2, 0, 7}}, expected: 9},
		{blueprint: Blueprint{oreCost: Cost{2, 0, 0}, clayCost: Cost{3, 0, 0}, obsidianCost: Cost{3, 8, 0}, geodeCost: Cost{3, 0, 12}}, expected: 12},
	}

	for _, tt := range testCases {
		value := ProbabilisticBlueprint(tt.blueprint, 1_000, 24)

		if value != tt.expected {
			t.Errorf("Value incorrect, expected=%d, got=%d", tt.expected, value)
		}
	}
}

func TestProbabilisticBlueprintPart2(t *testing.T) {
	testCases := []BlueprintTestCase{
		{blueprint: Blueprint{oreCost: Cost{4, 0, 0}, clayCost: Cost{2, 0, 0}, obsidianCost: Cost{3, 14, 0}, geodeCost: Cost{2, 0, 7}}, expected: 56},
		{blueprint: Blueprint{oreCost: Cost{2, 0, 0}, clayCost: Cost{3, 0, 0}, obsidianCost: Cost{3, 8, 0}, geodeCost: Cost{3, 0, 12}}, expected: 62},
	}

	for _, tt := range testCases {
		value := ProbabilisticBlueprint(tt.blueprint, 100_000, 32)

		if value != tt.expected {
			t.Errorf("Value incorrect, expected=%d, got=%d", tt.expected, value)
		}
	}
}
