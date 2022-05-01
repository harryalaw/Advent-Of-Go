package main

import (
	"testing"
)

func TestTime(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
		{"312211", "13112221"},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			output := LookAndSay(tc.input)
			if output != tc.want {
				t.Errorf("got %s; want %s", output, tc.want)
			}
		})
	}
}

func TestActual(t *testing.T) {
	testCases := []struct {
		part     string
		expected int
	}{
		{"Part1", 492982},
		{"Part2", 6989950},
	}
	for _, tc := range testCases {
		t.Run(tc.part, func(t *testing.T) {
			var output int
			if tc.part == "Part1" {
				output = Part1()
			} else {
				output = Part2()
			}
			if output != tc.expected {
				t.Errorf("got %d; want %d", output, tc.expected)
			}
		})
	}
}
