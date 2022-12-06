package main

import (
	_ "embed"
	"testing"
)

func TestPart1(t *testing.T) {
	// input := parseInput(testData)
	testData := []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tt := range testData {
		index := Part1(tt.input)
		if index != tt.expected {
			t.Errorf("Wrong index found, expected=%d, got=%d", tt.expected, index)
		}
	}

}

func TestPart2(t *testing.T) {
	testData := []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, tt := range testData {
		index := Part2(tt.input)
		if index != tt.expected {
			t.Errorf("Wrong index found, expected=%d, got=%d", tt.expected, index)
		}
	}
}
