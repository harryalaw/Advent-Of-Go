package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

var testData2 string = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`

func TestPart1(t *testing.T) {
	input := parseInput(testData)

	value := Part1(input)
	expected := 13

	if value != expected {
		t.Errorf("Value not correct, expected=%d got=%d", expected, value)
	}
}

func TestPart2(t *testing.T) {
	manyTests := []struct {
		input    string
		expected int
	}{
		{testData, 1},
		{testData2, 36},
	}

	for _, tt := range manyTests {
		input := parseInput(tt.input)

		value := Part2(input)
		expected := tt.expected

		if value != expected {
			t.Errorf("Value not correct, expected=%d got=%d", expected, value)
		}

	}
}
