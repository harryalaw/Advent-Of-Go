package main

import (
	_ "embed"
	"fmt"
	"github.com/harryalaw/advent-of-go/util"
	"strings"
)

//go:embed data.txt
var input string

func main() {
	util.Time(doPart1, "Part1")
	util.Time(doPart2, "Part2")
}

func doPart1() {
	fmt.Println(Part1(strings.TrimSpace(input)))
}
func doPart2() {
	fmt.Println(Part2(strings.TrimSpace(input)))
}

type Marker struct {
	length  int
	repeats int
}


func Part1(input string) int {
	return parse(input, false)
}

func Part2(input string) int {
	return parse(input, true)
}

func parse(input string, isPart2 bool) int {
	length := 0
	head := 0

	for head < len(input) {
		c := input[head]
		if c != '(' {
			length++
			head++
		} else {
			marker := parseMarker(input, &head)

			if !isPart2 {
				length += marker.repeats * marker.length
			} else {
				length += marker.repeats * parse(input[head:head+marker.length], true)
			}
			head += marker.length
		}
	}

	return length
}

func parseMarker(input string, head *int) Marker {
	*head++
	lengthText := strings.Builder{}
	repeatText := strings.Builder{}

	for input[*head] != 'x' {
		lengthText.WriteByte(input[*head])
		*head++
	}
	*head++

	for input[*head] != ')' {
		repeatText.WriteByte(input[*head])
		*head++
	}
	*head++

	marker := Marker{
		length:  util.Atoi(lengthText.String()),
		repeats: util.Atoi(repeatText.String()),
	}

	return marker
}
