package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

func TestPart1(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: testData, expected: 13140},
	}

	for _, tt := range testCases {
		input := parseInput(tt.input)

		value := Part1(input)
		if value != tt.expected {
			t.Errorf("Wrong value, expected=%d, got=%d", tt.expected, value)
		}

	}
}

func TestPart2(t *testing.T) {
	input := parseInput(testData)
	expected := `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`
	value := Part2(input)

	if value != expected {
		t.Errorf("Incorrect value\n expected=\n%s\n\n got=\n%s\n", expected, value)
	}
}
