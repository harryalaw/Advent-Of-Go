package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("data.txt")

	// floor := 0

	string := string(content)

	up := strings.Count(string, "(")

	down := strings.Count(string, ")")

	floor := up - down

	println(floor)
	// for _, char := range string(content) {
	// 	if char == '(' {
	// 		floor++
	// 	} else {
	// 		floor--
	// 	}
	// }
	// fmt.Println(floor)
}
