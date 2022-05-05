package main

import (
	"strconv"
	"testing"
)

func TestPart1(t *testing.T) {
	t.Run("Blah", func(t *testing.T) {
		Part1()
	})
}

func TestFactorise(t *testing.T) {
	testCases := []struct {
		input int
		want  []int
	}{
		{5, []int{5}},
		{12, []int{4, 3}},
		{36, []int{4, 9}},
		{210, []int{2, 3, 5, 7}},
		{1549, []int{1549}},
	}
	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.input), func(t *testing.T) {
			output := Factorise(tc.input)
			if !arrayEquals(output, tc.want) {
				t.Errorf("got %+v; want %+v", output, tc.want)
			}
		})
	}
}

func arrayEquals(actual, want []int) bool {
	if len(actual) != len(want) {
		return false
	}
	for i := 0; i < len(actual); i++ {
		if actual[i] != want[i] {
			return false
		}
	}
	return true
}
