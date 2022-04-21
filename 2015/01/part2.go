package main

import (
	"io/ioutil"
)

func main() {
	content, _ := ioutil.ReadFile("data.txt")

	floor := 0

	for index, char := range string(content) {
		if char == '(' {
			floor++
		} else {
			floor--
		}
		if floor == -1 {
			println(index + 1)
			return
		}
	}
}
