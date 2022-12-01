package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

func TestParseInput(t *testing.T) {
	testCase := [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10000},
	}

	testElves := parseInput(testData)

	if len(testElves) != 5 {
		t.Errorf("Did not create 5 elves, got=%d", len(testElves))
	}

	for i, elf := range testElves {
		for j, calorie := range elf {
			if testCase[i][j] != calorie {
				t.Errorf("Expected elf %d to have %d calories. got=%d", i, testCase[i][j], calorie)
			}
		}
	}
}

func TestPart1(t *testing.T) {
	testElves := parseInput(testData)

	value := Part1(testElves)

	if value != 24000 {
		t.Errorf("Wrong value for max calories, expected=%d got=%d", 24000, value)
	}
}

func TestPart2(t *testing.T) {
	testElves := parseInput(testData)

	value := Part2(testElves)

	if value != 45000 {
		t.Errorf("Wrong value for max calories, expected=%d got=%d", 45000, value)
	}
}
