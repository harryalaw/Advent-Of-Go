package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func solve(part func(string) int) int {
	value := part("data.txt")
	println(value)
	return value
}

func main() {
	start := time.Now()
	solve(part1)
	elapsed := time.Since(start)
	fmt.Printf("Part1 took %s\n", elapsed)

	start = time.Now()
	solve(part2)
	elapsed = time.Since(start)
	fmt.Printf("Part2 took %s\n", elapsed)
}

func readFile(fileName string) ([]string, int) {
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}
	words := strings.Split(string(file), "\n")

	return words, len(file)
}

func part1(fileName string) int {
	words, fileLen := readFile(fileName)
	wordsLength := 0

	for _, word := range words {
		unquoted, err := strconv.Unquote(word)
		if err != nil {
			panic(err)
		}
		wordsLength += len(unquoted)
	}

	fileLength := fileLen - len(words) + 1

	return fileLength - wordsLength
}

func part2(fileName string) int {
	words, fileLen := readFile(fileName)
	wordsLength := 0

	for _, word := range words {
		unquoted := strconv.Quote(word)
		wordsLength += len(unquoted)
	}

	fileLength := fileLen - len(words) + 1

	return wordsLength - fileLength
}
