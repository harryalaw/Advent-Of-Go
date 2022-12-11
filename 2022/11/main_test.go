package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

func TestPart1(t *testing.T) {
	input, _ := parseInput(testData)

	value := Part1(input)
	expected := 10605

	if value != expected {
		t.Errorf("Value wrong, expected=%d, got=%d", expected, value)
	}
}

func TestPart2(t *testing.T) {
	input, lcm := parseInput(testData)

	value := Part2(input, lcm)
	expected := 2713310158

	if value != expected {
		t.Errorf("Value wrong, expected=%d, got=%d", expected, value)
	}
}
