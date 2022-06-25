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

// find all subsets of an int array that sum to K
func sumToK(numbers []int, k int) [][]int {
	out := make([][]int, 0)

	for i, number := range numbers {
		attempt := []int{number}
		out = recursiveSumToK(numbers[i+1:], attempt, k, out)
	}

	return out
}

func recursiveSumToK(numbers []int, attempt []int, k int, out [][]int) [][]int {
	total := sumArray(attempt)
	if total == k {
		tempAttempt := copyInts(attempt)
		out = append(out, tempAttempt)
		return out
	}
	if len(numbers) == 0 {
		return out
	}

	for i, number := range numbers {
		if total+number <= k {
			tempAttempt := append(attempt, number)
			out = recursiveSumToK(numbers[i+1:], tempAttempt, k, out)
		}
	}
	return out
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

		tempAttempt := copyInts(attempt)
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
	numbers = sortDescending(numbers)
	target := sumArray(numbers)
	if target%3 != 0 {
		panic("Not divisible by 3")
	}
	target /= 3

	partitions := sumToKPruned(numbers, target)

	minLength := len(numbers)
	quantumEntanglement := math.MaxInt64
	for _, sum := range partitions {
		if len(sum) > minLength {
			continue
		}
		tempQE := multiplyArray(sum)
		if tempQE < quantumEntanglement || len(sum) < minLength {
			quantumEntanglement = tempQE
		}

		minLength = len(sum)
	}
	return quantumEntanglement
}

func part2(numbers []int) int {
	numbers = sortDescending(numbers)
	target := sumArray(numbers)
	if target%4 != 0 {
		panic("Not divisible by 4")
	}
	target /= 4

	partitions := sumToKPruned(numbers, target)

	minLength := len(numbers)
	quantumEntanglement := math.MaxInt64
	var finalPartition []int

	for i := len(partitions) - 1; i >= 0; i-- {
		sum := partitions[i]
		if len(sum) > minLength || sumArray(sum) != target {
			continue
		}
		tempQE := multiplyArray(sum)
		if tempQE < quantumEntanglement || len(sum) < minLength {
			quantumEntanglement = tempQE
			finalPartition = sum
		}

		if len(sum) < minLength {
			minLength = len(sum)
		}
	}
	fmt.Printf("Final Partition: %+v\n", finalPartition)
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

func copyInts(numbers []int) []int {
	out := make([]int, len(numbers))

	for i, number := range numbers {
		out[i] = number
	}

	return out
}
