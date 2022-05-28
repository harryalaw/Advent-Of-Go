package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"[1,2,3]", 6},
		{"{\"a\":2,\"b\":4}", 6},
		{"[[[3]]]", 3},
		{"{\"a\":{\"b\":4},\"c\":-1}", 3},
		{"{\"a\":[-1,1]}", 0},
		{"[-1,{\"a\":1}]", 0},
		{"[]", 0},
		{"{}", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			output := CoreLogic(tc.input, false)
			if int(output) != tc.want {
				t.Errorf("got %d; want %d", int(output), tc.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"[1,2,3]", 6},
		{"[1, {\"c\": \"red\", \"b\":2},3]", 4},
		{"{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}", 0},
		{"[1, \"red\", 5]", 6},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			output := CoreLogic(tc.input, true)
			if int(output) != tc.want {
				t.Errorf("got %d; want %d", int(output), tc.want)
			}
		})
	}
}
