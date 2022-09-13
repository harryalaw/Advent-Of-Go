package main

import (
	"github.com/harryalaw/advent-of-go/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	util.Time(doPart(CoreLogic, false), "Part1")
	util.Time(doPart(CoreLogic, true), "Part2")

}

func doPart(part func(string, bool) float64, isPart2 bool) func() {
	return func() {
		fmt.Println(part(parseInput(), isPart2))
	}
}

func parseInput() string {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func CoreLogic(input string, isPart2 bool) float64 {
	var jsonData interface{}
	err := json.Unmarshal([]byte(input), &jsonData)
	if err != nil {
		panic(err)
	}

	totalCount := 0.0

	countNumbers(jsonData, &totalCount, isPart2)

	return totalCount
}

func countNumbers(parsed interface{}, sum *float64, part2 bool) {

	switch val := parsed.(type) {
	case map[string]interface{}:
		nestedSum := 0.0
		for _, value := range val {
			if part2 && value == "red" {
				nestedSum = 0.0
				return
			}
			countNumbers(value, &nestedSum, part2)
		}
		*sum += nestedSum
	case float64:
		*sum += val
	case []interface{}:
		for _, value := range val {
			countNumbers(value, sum, part2)
		}
	}
}
