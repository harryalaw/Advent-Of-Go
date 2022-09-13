package main

import (
	"github.com/harryalaw/advent-of-go/util"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	util.Time(part1, "Part1")
	util.Time(part2, "Part2")
}

func part1() {
	fmt.Println(doPart1("data.txt"))
}

func part2() {
	fmt.Println(doPart2("data.txt"))
}

func doPart1(file string) int {
	transformMap, chemical := parseInput(file)

	seenChemicals := make(map[string]struct{})

	for i, symbol := range chemical {
		transforms := transformMap[symbol]

		for _, transform := range transforms {
			tempChemical := applyTransform(chemical, i, transform)

			seenChemicals[tempChemical] = struct{}{}
		}
	}
	return len(seenChemicals)
}

func doPart2(file string) int {
	transforms, finalTransforms, endResult := parseInput2(file)

	steps := 0
	nextChemical := strings.Join(endResult, "")
	_, isPresent := finalTransforms[nextChemical]

	for !isPresent {
		lastIndex := -1
		var lastFrom, lastTo string

		for from, to := range transforms {
			tempIndex := strings.LastIndex(nextChemical, from)
			if tempIndex > lastIndex {
				lastIndex = tempIndex
				lastFrom = from
				lastTo = to
			}
		}

		nextChemical = replaceLastOccurence(nextChemical, lastFrom, lastTo, lastIndex)
		steps++
		_, isPresent = finalTransforms[nextChemical]
	}

	return steps + 1
}

func replaceLastOccurence(input, replace, with string, lastIndex int) string {
	return input[:lastIndex] + with + input[lastIndex+len(replace):]
}

func parseInput(file string) (map[string][][]string, []string) {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	sections := strings.Split(string(data), "\n\n")

	transforms := createTransforms(sections[0])
	tokens := extractTokens(sections[1])

	return transforms, tokens
}

func createTransforms(input string) map[string][][]string {
	segments := strings.Split(input, "\n")

	transforms := make(map[string][][]string)
	for _, rule := range segments {
		parts := strings.Split(rule, " => ")
		from, to := parts[0], extractTokens(parts[1])

		prevArray, isPresent := transforms[from]

		if !isPresent {
			prevArray = make([][]string, 0)
		}

		prevArray = append(prevArray, to)

		transforms[from] = prevArray
	}
	return transforms
}

func parseInput2(file string) (map[string]string, map[string]struct{}, []string) {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	sections := strings.Split(string(data), "\n\n")

	transforms, finalTransforms := createTransforms2(sections[0])
	tokens := extractTokens(sections[1])

	return transforms, finalTransforms, tokens

}

func createTransforms2(input string) (map[string]string, map[string]struct{}) {
	segments := strings.Split(input, "\n")

	finalTransforms := make(map[string]struct{})

	transforms := make(map[string]string)
	for _, rule := range segments {
		parts := strings.Split(rule, " => ")
		from, to := parts[1], parts[0]
		if to == "e" {
			finalTransforms[from] = struct{}{}
		} else {
			transforms[from] = to
		}
	}
	return transforms, finalTransforms
}

func extractTokens(input string) []string {
	re := regexp.MustCompile("[A-Z][a-z]?")

	matches := re.FindAllString(input, -1)
	return matches
}

func applyTransform(chemical []string, index int, transform []string) string {
	var sb strings.Builder
	for i, el := range chemical {
		if i != index {
			sb.WriteString(el)
		} else {
			for _, newEl := range transform {
				sb.WriteString(newEl)
			}
		}
	}
	return sb.String()
}
