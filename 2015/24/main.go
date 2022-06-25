package main

import (
	util "2015/Util"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func main() {
	util.Time(doPart1, "Part1")
	util.Time(doPart2, "Part2")
}

func doPart1() {
	numbers := parseInput("data.txt")

	fmt.Println(part1(numbers))
}

func doPart2() {
	numbers := parseInput("data.txt")

	fmt.Println(part2(numbers))
}

func parseInput(fileName string) []int {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	numbers := make([]int, len(lines))

	for i := 0; i < len(lines); i++ {
		numbers[i] = util.Atoi(lines[i])
	}

	return numbers
}

func sumToKPruned(numbers []int, k int) [][]int {
	out := make([][]int, 0)

	minlength := len(numbers)
	for i, number := range numbers {
		attempt := []int{number}
		out = recursiveSumToKPruned(numbers[i+1:], attempt, k, out, &minlength)
	}
	return out
}

func recursiveSumToKPruned(numbers, attempt []int, k int, out [][]int, minLength *int) [][]int {
	total := sumArray(attempt)
	if total == k && len(attempt) <= *minLength {
		if len(attempt) < *minLength {
			out = make([][]int, 0)
			*minLength = len(attempt)
		}

		tempAttempt := make([]int, len(attempt))
		copy(tempAttempt, attempt)
		out = append(out, tempAttempt)
		return out
	}

	for i, number := range numbers {
		if total+number <= k && len(attempt) < *minLength {
			tempAttempt := append(attempt, number)
			out = recursiveSumToKPruned(numbers[i+1:], tempAttempt, k, out, minLength)
		}
	}
	return out
}

func sumArray(array []int) int {
	total := 0
	for _, el := range array {
		total += el

	}
	return total
}

func multiplyArray(array []int) int {
	total := int(1)
	for _, el := range array {
		total *= int(el)
		if total < 0 {
			return math.MaxInt64
		}
	}
	return total
}

func part1(numbers []int) int {
	return coreLogic(numbers, 3)
}

func part2(numbers []int) int {
	return coreLogic(numbers, 4)
}

func coreLogic(numbers []int, partitionCount int) int {
	numbers = sortDescending(numbers)

	target := sumArray(numbers)
	if target%partitionCount != 0 {
		fmt.Println("Partition count:", partitionCount)
		panic("Not divisible by partitionCount")
	}
	target /= partitionCount

	partitions := sumToKPruned(numbers, target)

	quantumEntanglement := math.MaxInt64

	for _, sum := range partitions {
		tempQE := multiplyArray(sum)
		if tempQE < quantumEntanglement {
			quantumEntanglement = tempQE
		}
	}
	return quantumEntanglement
}

func sortDescending(numbers []int) []int {
	tempNumbers := numbers
	sort.Ints(tempNumbers)

	sorted := make([]int, len(tempNumbers))
	for i := 0; i < len(tempNumbers); i++ {
		sorted[len(tempNumbers)-i-1] = tempNumbers[i]
	}

	return sorted
}
