package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func Time(function func(), part string) {
	start := time.Now()
	function()
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", part, elapsed)
}

func main() {
	input := parseInput("data.txt")
	var nextInput string
	Time(func() {
		nextInput = Part1(input)
		fmt.Println(nextInput)
	}, "Part1")
	nextInput = IntToStringArray(incrementValues(StringToIntArray(nextInput)))
	Time(func() {
		fmt.Println(Part1(nextInput))

	}, "Part2")
}

func Validate(values []int) bool {
	increasingStraight := false

	prevValue := values[0]
	length := 0
	for _, value := range values {
		if value == 8 ||
			value == 11 ||
			value == 14 {
			return false
		}
		if !increasingStraight &&
			value == prevValue+1 {
			length++
			if length >= 3 {
				increasingStraight = true
			}
		} else {
			length = 1
		}
		prevValue = value
	}
	pairs := twoPairs(values)

	return pairs && increasingStraight
}

func twoPairs(values []int) bool {
	pair1 := -1

	for index := 0; index < len(values)-1; index++ {
		if values[index] != values[index+1] {
			continue
		} else if pair1 != -1 && pair1 != values[index] {
			return true
		} else {
			pair1 = values[index]
		}
	}
	return false
}

func StringToIntArray(value string) []int {
	output := make([]int, len(value))

	for index, char := range value {
		output[index] = int(char - 'a')
	}

	return output
}
func IntToStringArray(value []int) string {
	var sb strings.Builder
	for _, num := range value {
		sb.WriteRune(rune(num + 'a'))
	}
	return sb.String()
}

func parseInput(fileName string) string {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	return string(data)
}
func incrementValues(values []int) []int {
	/*
		start at last value, then increment last value % 26
		then perform carry overs
	*/
	carry := 0
	lastIndex := len(values) - 1
	values[lastIndex] = (values[lastIndex] + 1) % 26
	if values[lastIndex] == 0 {
		carry = 1
	}
	for carry != 0 {
		for i := len(values) - 2; i >= 0; i-- {
			values[i] = (values[i] + 1) % 26
			if values[i] == 0 {
				carry = 1
			} else {
				carry = 0
				break
			}
		}
	}
	// fmt.Printf("%+v\n", values)
	return values
}

func Part1(input string) string {
	values := StringToIntArray(input)
	for !Validate(values) {
		values = incrementValues(values)
	}
	return IntToStringArray(values)
}

/*
hxbxwxba
hxbxxaaa
hxbxxabc



*/
// abcde
// fgjij
// klmno
// pqrst
// uvwxy
// z
