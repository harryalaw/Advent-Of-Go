package main

import "testing"

func TestValidate(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		// {"hijklmmn", false},
		// {"abbceffg", false},
		// {"abbcegjk", false},
		// {"abcdffaa", true},
		{"ghjaabcc", true},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			output := Validate(StringToIntArray(tc.input))
			if output != tc.want {
				t.Errorf("got %t; want %t", output, tc.want)
			}
		})
	}
}

func TestStringToIntArray(t *testing.T) {
	testCases := []struct {
		input string
		want  []int
	}{
		{"abc", []int{0, 1, 2}},
		{"hijklmmn", []int{7, 8, 9, 10, 11, 12, 12, 13}},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			output := StringToIntArray(tc.input)
			if !Equal(output, tc.want) {
				t.Errorf("got %+v; want %+v", output, tc.want)
			}
		})
	}
}

func TestIntArrayToString(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		want  string
	}{
		{"0,1,2", []int{0, 1, 2}, "abc"},
		{"7,8,9,10,11,12,12,13", []int{7, 8, 9, 10, 11, 12, 12, 13}, "hijklmmn"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := IntToStringArray(tc.input)
			if output != tc.want {
				t.Errorf("got %s; want %s", output, tc.want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			output := Part1(tc.input)
			if output != tc.want {
				t.Errorf("got %s; want %s", output, tc.want)
			}
		})
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
