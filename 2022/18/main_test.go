package main

import (
	_ "embed"
	"testing"
)

//go:embed test_data.txt
var testData string

func TestPart1(t *testing.T) {
	cubes, space := parseInput(testData)

	value := Part1(cubes, space)
	expected := 64

	if value != expected {
		t.Errorf("Incorrect value, got=%d, expected=%d", value, expected)
	}
}

func TestPart2(t *testing.T) {
	cubes, space := parseInput(testData)

	value := Part2(cubes, space)
	expected := 58

	if value != expected {
		t.Errorf("Incorrect value, got=%d, expected=%d", value, expected)
	}
}
