package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// println(part1("test.txt"))
	println(part1("data.txt"))

	println(part2("test.txt"))
	println(part2("data.txt"))
}

func part1(fileName string) int {
	return coreLogic(fileName, strconv.Unquote)
}
func part2(fileName string) int {
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}
	words := strings.Split(string(file), "\n")

	wordsLength := 0

	for _, word := range words {
		unquoted := strconv.Quote(word)
		wordsLength += len(unquoted)
	}

	fileLength := len(file) - len(words) + 1
	println()

	return wordsLength - fileLength
}

func coreLogic(fileName string, operate func(string) (string, error)) int {
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}
	words := strings.Split(string(file), "\n")

	wordsLength := 0

	for _, word := range words {
		unquoted, err := operate(word)
		if err != nil {
			panic(err)
		}
		wordsLength += len(unquoted)
	}

	fileLength := len(file) - len(words) + 1
	println()

	return fileLength - wordsLength
}
