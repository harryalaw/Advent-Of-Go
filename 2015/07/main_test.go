package main

import (
	"strings"
	"testing"
)

func TestCoreLogic(t *testing.T) {
	data := "123 -> x\n456 -> y\nx AND y -> d\nx OR y -> e\nx LSHIFT 2 -> f\ny RSHIFT 2 -> g\nNOT x -> h\nNOT y -> i"

	lines := strings.Split(data, "\n")
	instructions := make([]*instruction, 0, len(lines))

	for _, line := range lines {
		instruction := MapStringToInstruction(line)
		instructions = append(instructions, instruction)
	}

	knownValues := make(map[string]int)
	CoreLogic(instructions, knownValues)

	checkValue(t, "d", knownValues, 72)
	checkValue(t, "e", knownValues, 507)
	checkValue(t, "f", knownValues, 492)
	checkValue(t, "g", knownValues, 114)
	checkValue(t, "h", knownValues, 65412)
	checkValue(t, "i", knownValues, 65079)
	checkValue(t, "x", knownValues, 123)
	checkValue(t, "y", knownValues, 456)

}

func checkValue(t *testing.T, key string, knownValues map[string]int, expected int) {
	actual, isPresent := knownValues[key]

	if !isPresent {
		t.Errorf("No value for %s found in knownValues", key)
	}
	if expected != actual {
		t.Errorf("Value was incorrect, got:%d, expected: %d.", actual, expected)
	}
}
