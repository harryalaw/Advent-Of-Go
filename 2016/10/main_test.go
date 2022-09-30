package main

import "testing"

func TestCoreLogic(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{`value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2`, 30},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			output := coreLogic(tc.input, false)
			if output != tc.want {
				t.Errorf("got %d; want %d", output, tc.want)
			}
		})
	}
}
