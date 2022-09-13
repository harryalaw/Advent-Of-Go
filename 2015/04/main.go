package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"github.com/harryalaw/advent-of-go/util"
	"strings"
)

//go:embed data.txt
var input string

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func solveHash(prefix, input string) int {
	for i := 1; ; i++ {
		testString := fmt.Sprintf("%s%d", input, i)
		hash := getMD5Hash(testString)
		if strings.HasPrefix(hash, prefix) {
			return i
		}
	}
}

func part1(input string) int {
	return solveHash("00000", input)
}

func part2(input string) int {
	return solveHash("000000", input)
}

func solve(part func(string) int) {
	fmt.Println(part(input))
}

func test(part func(string) int, input string, expected int) bool {
	result := part(input)
	return result == expected
}

func runTests(actual bool) bool {
	test1 := test(part1, "abcdef", 609043)
	test2 := test(part1, "pqrstuv", 1048970)

	test3 := true
	test4 := true
	if actual {
		test3 = test(part1, input, 254575)
		test4 = test(part2, input, 1038736)
	}
	return test1 && test2 && test3 && test4
}

func main() {
	fmt.Println(runTests(false))
	util.Time(func() { solve(part1) }, "Part1")
	util.Time(func() { solve(part2) }, "Part2")
}
